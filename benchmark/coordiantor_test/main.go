package main

import (
	"bufio"
	"context"
	"coordiantor_test/informer"
	pb "coordiantor_test/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"os"
	"strconv"

	"sync"
	"time"
)

func main() {
	// 1. 告诉dispatcher scheduler的地址
	schelist := make([]*pb.SchedulerInfo, 0)
	for k, v := range informer.Scheduler_map {
		schelist = append(schelist, &pb.SchedulerInfo{
			NodeName: k,
			Address:  v,
		})
	}
	for k, v := range informer.Dispatcher_map {
		connDispatcher, _ := grpc.Dial(v+":16444", grpc.WithInsecure())
		defer connDispatcher.Close()
		clientDispatcher := pb.NewDispatcherClient(connDispatcher)
		resp, _ := clientDispatcher.UpdateSchedulerView(context.Background(), &pb.SchedulerViewUpdate{List: schelist, Action: "ADD"})
		log.Printf("Update Scheduler View %s resp: %d", k, resp.State)
	}

	// 2. 为scheduler分配管理的node 相关信息
	schedulerKeys := make([]string, 0, len(informer.Scheduler_map))
	for key := range informer.Scheduler_map {
		schedulerKeys = append(schedulerKeys, key)
	}
	i := 0
	for k, v := range informer.Virtualnode_map {
		//fmt.Println(informer.Scheduler_map[schedulerKeys[i%len(schedulerKeys)]] + ":16445")
		conn, _ := grpc.Dial(informer.Scheduler_map[schedulerKeys[i%len(schedulerKeys)]]+":16445", grpc.WithInsecure())
		defer conn.Close()
		client := pb.NewSchedulerClient(conn)
		list := make([]*pb.NodeResource, 0)
		list = append(list, &pb.NodeResource{
			NodeName: k,
			HaveCpu:  1024 * 4,
			HaveMem:  1024 * 4,
			Address:  v,
		})
		i++
		_, _ = client.UpadateNodeResource(context.Background(), &pb.NodeResourceUpdate{List: list, Action: "ADD"})
		fmt.Printf("Assign %d nodes to scheduler %s \n", len(list), k)
	}
	// 3. 告知所有的scheduler peerSchedulers
	for k, v := range informer.Scheduler_map {
		schelist := make([]*pb.PeerSchedulerInfo, 0)
		for pk, pv := range informer.Scheduler_map {
			if k != pk {
				schelist = append(schelist, &pb.PeerSchedulerInfo{
					NodeName: pk,
					Address:  pv,
				})
			}
		}
		conn, _ := grpc.Dial(v+":16445", grpc.WithInsecure())
		defer conn.Close()
		client := pb.NewSchedulerClient(conn)

		resp, _ := client.PeerSchedulerUpdate(context.Background(), &pb.PeerSchedulersUpdate{
			List:   schelist,
			Action: "ADD",
		})
		fmt.Println(resp)
	}

	// 4. 向不同的dispatcher转发请求
	workload := func_load("../load_generator/data/origin")
	total := 1500
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
	//dispatcher_load(per_worker, workers, sl, new_workerload)

	//select {}
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
