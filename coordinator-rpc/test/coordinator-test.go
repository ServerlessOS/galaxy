package main

import (
	"bufio"
	"context"
	"coordinator_rpc/informer"
	"fmt"
	pb "github.com/ServerlessOS/galaxy/proto"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type nodeInfo struct {
	name string
	addr string
}

func main() {

	conn, _ := grpc.Dial(":16000", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewCoordinatorClient(conn)
	time.Sleep(3 * time.Second)
	sadd := addScheduler(client)
	// add scheduler
	for schedulerName, _ := range informer.Scheduler_map {
		fmt.Printf("Add the new Scheduler: %s\n", schedulerName)
		sadd()
	}
	//add dispatcher
	for dispatcherName, dispatcherAddr := range informer.Dispatcher_map {
		fmt.Printf("Add new dispatcher %s\n", dispatcherName)
		ndispatcher := &pb.DispatcherInfoUpdate{
			SchedulerName: dispatcherName,
			Address:       dispatcherAddr,
		}
		resp, err := client.AddDispatcherInfo(context.Background(), ndispatcher)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp.Message)
	}
	for nodeName, nodeAddr := range informer.Virtualnode_map {
		nu := &pb.NodeInfoUpdate{
			NodeName: nodeName,
			Address:  nodeAddr,
			HaveCpu:  1024 * 4,
			HaveMem:  1024 * 4,
		}
		fmt.Printf("Add Node: %s:%s\n", nu.NodeName, nu.Address)
		resp, err := client.AddNodeInfo(context.Background(), nu)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp.Message)
	}
	time.Sleep(2 * time.Second)
	// load generator
	workload := func_load("/home/tank/bsz/ryze/rpc/benchmark/load_generator/data/origin")
	total := 1000
	workers := 50
	per_worker := total / workers

	new_workerload, err := random_select(workload, total)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("len: %d ,Sum: %d\n", len(new_workerload), Sum(new_workerload))
	//os.Exit(0)
	dispatcher_urls := make([]string, 0)
	for _, v := range informer.Dispatcher_map {
		dispatcher_urls = append(dispatcher_urls, v)
		//println("arr:" + v)
	}
	dispatcher_load(per_worker, workers, dispatcher_urls, new_workerload)
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func addNode(client pb.CoordinatorClient) {
	nu := &pb.NodeInfoUpdate{
		NodeName: fmt.Sprintf("Node-%s", randomString(6)),
		Address:  randomIP(),
		HaveCpu:  1024 * 4,
		HaveMem:  1024 * 4,
	}
	fmt.Printf("Add Node: %s:%s\n", nu.NodeName, nu.Address)
	resp, err := client.AddNodeInfo(context.Background(), nu)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Message)
}
func addScheduler(client pb.CoordinatorClient) func() {
	index := 0
	scheduler_list := make([]string, 0, len(informer.Scheduler_map))
	for key := range informer.Scheduler_map {
		scheduler_list = append(scheduler_list, key)
	}
	return func() {
		su := &pb.SchedulerInfoUpdate{
			SchedulerName: scheduler_list[index],
			Address:       informer.Scheduler_map[scheduler_list[index]],
		}
		fmt.Printf("Add Scheduler: %s:%s\n", su.SchedulerName, su.Address)
		resp, err := client.AddSchedulerInfo(context.Background(), su)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp.Message)
		index++
	}
}

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
func randomIP() string {
	ip := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return ip.String()
}

func dispatcher_load(per_worker, n int, addr []string, workload []int) {
	var wg sync.WaitGroup
	dispatcher_num := len(addr)
	st := time.Now()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go dispatcher_thread(i*per_worker, per_worker, addr[i%dispatcher_num], workload[per_worker*i:per_worker*(i+1)], &wg)
		//go dispatcher_thread(i*per_worker, per_worker, "127.0.0.1", workload[per_worker*i:per_worker*(i+1)], &wg)
	}
	wg.Wait()
	endTime := time.Now()

	// 计算执行时间，并将其转换为毫秒
	elapsedTime := endTime.Sub(st).Milliseconds()
	fmt.Printf("Cost time : %d ms", elapsedTime)

}
func dispatcher_thread(start_index, total int, dispatcher_address string, workload []int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	fmt.Printf("Thread %s start_index %d has %d wkload\n", dispatcher_address, start_index, len(workload))
	connDispatcher, _ := grpc.Dial(dispatcher_address+":16444", grpc.WithInsecure())
	defer connDispatcher.Close()
	client := pb.NewDispatcherClient(connDispatcher)
	l := make([]*pb.UserRequest, 0)
	for j := start_index; j < start_index+total; j++ {
		ur := &pb.UserRequest{
			RequestId:  int64(j),
			FuncName:   fmt.Sprintf("Func-%d", j),
			RequireCpu: 1,
			RequireMem: int64(workload[i]),
		}
		i++
		l = append(l, ur)
	}
	_, _ = client.Dispatch(context.Background(), &pb.UserRequestList{List: l})
}

func func_load(filepath string) []int {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return nil
	}
	defer file.Close()

	memorys := make([]int, 0)
	// 使用 bufio.Scanner 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m, _ := strconv.Atoi(line)
		memorys = append(memorys, m)
	}
	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
	}
	return memorys
}
func random_select(workload []int, n int) ([]int, error) {
	if n > len(workload) {
		return nil, fmt.Errorf("选择的数量大于数组长度")
	}

	rand.Seed(time.Now().UnixNano())

	// 复制数组，避免修改原始数组
	copyArr := make([]int, len(workload))
	copy(copyArr, workload)

	selected := make([]int, 0)
	times := n / len(workload)
	for times > 0 {
		selected = append(selected, workload...)
		times--
	}

	// 随机选取 n%len 个不重复的元素
	for i := 0; i < n%len(copyArr); i++ {
		// 随机生成索引
		randomIndex := rand.Intn(len(copyArr))
		// 从数组中取出元素并从数组中移除
		selected = append(selected, copyArr[randomIndex])
		copyArr = append(copyArr[:randomIndex], copyArr[randomIndex+1:]...)
	}

	return selected, nil
}
func Sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
