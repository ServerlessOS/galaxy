package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServerlessOS/galaxy/constant"
	"github.com/ServerlessOS/galaxy/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"math/rand/v2"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	localRpcAddr    string
	gatewayAddr     string
	funcManagerName string
	funcYaml        = make(map[string]*function) //通过函数名索引函数
	ExitCh          = make(chan int)
	functionDir     = "." + string(os.PathSeparator) + "function"
	localIp         = getLocalIPv4().String()
)

type function struct {
	Name       string
	Label      string
	Annotation string
	Document   string
}

var Cmd = &cobra.Command{
	Use:   "gateway",
	Short: `初始化func-manager程序`,
	//本函数用于执行命令并返回错误
	RunE: func(cmd *cobra.Command, args []string) error {
		funcManagerName = strconv.Itoa(int(rand.Uint32()))
		var errChanRpc chan error
		if !cmd.Flags().Changed("gatewayAddr") {
			return errors.New("gatewayAddr is required")
		}
		register()
		rpcServer(errChanRpc)
		err := <-errChanRpc
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			return err
		}
		return nil
	},
}

// Run 提供给顶层用于启动cobra根命令
func Run(cmd *cobra.Command) (code int) {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return 0
}
func init() {
	Cmd.Flags().StringVarP(&localRpcAddr, "localRpcAddr", "r", "0.0.0.0:"+constant.FuncManagerPort, "The addr used for binding to the RPC server. ")
	Cmd.Flags().StringVarP(&gatewayAddr, "gatewayAddr", "g", "", "The address information of the gateway needs to be registered with the gateway to work properly. ")
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)

	//创建保存function的文件
	if _, err := os.Stat(functionDir); os.IsNotExist(err) {
		// 不存在则创建
		err := os.Mkdir(functionDir, 0755)
		if err != nil {
			log.Fatalln("create function folder err:", err)
		}
		fmt.Println("Folder created successfully:", functionDir)
	} else {
		//存在则载入
		files, err := ioutil.ReadDir(functionDir)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// 循环处理每个文件
		for _, file := range files {
			fileName := file.Name()
			if !file.IsDir() && strings.HasSuffix(fileName, ".yaml") {
				// 解析文件名
				parts := strings.Split(strings.TrimSuffix(fileName, ".yaml"), "-")

				// 初始化标签和注释
				name := ""
				label := ""
				annotation := ""

				// 如果文件名中包含足够的部分，则解析标签和注释
				name = parts[0]
				if len(parts) >= 2 {
					label = parts[1]
				}
				if len(parts) >= 3 {
					annotation = parts[2]
				}
				// 读取文件内容
				fileContent, err := ioutil.ReadFile(functionDir + string(os.PathSeparator) + fileName)
				if err != nil {
					log.Errorln("Error reading yaml file:", fileName, err)
					continue
				}

				// 文件内容为yaml编码成的字符串
				fileContentStr := string(fileContent)
				// 构建函数结构体
				funcStruct := &function{
					Name:       name,
					Label:      label,
					Annotation: annotation,
					Document:   fileContentStr,
				}

				// 添加到函数映射中
				funcYaml[name] = funcStruct
			}
		}
	}
	log.Println("Init func-manager success.")
}
func register() {
	//通过gateway向顶层控制器注册
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	connGateway, err := grpc.Dial(gatewayAddr+":"+constant.GatewayRpcPort, grpc.WithInsecure(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		log.Fatalln("dial gateway error:", err)
	}
	client := proto.NewGatewayClient(connGateway)
	tcpAddr, err := net.ResolveTCPAddr("tcp", localRpcAddr)
	_, err = client.Register(ctx, &proto.RegisterReq{
		Type:    1, //    coordinator = 0; funcManager = 1;
		Name:    funcManagerName,
		Address: tcpAddr.IP.String(),
	})
}

func rpcServer(errChannel chan<- error) {
	lis, err := net.Listen("tcp", localRpcAddr)
	if err != nil {
		errChannel <- err
	}
	// 实例化grpc服务端
	s := grpc.NewServer()

	// 在gRPC服务器注册服务
	proto.RegisterFuncManagerServer(s, &rpcServerProcess{})

	// 启动grpc服务
	err = s.Serve(lis)
	errChannel <- err
}
func GracefulExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	log.Printf("catch signal, %+v", sig)

	file, err := yaml.Marshal(funcYaml)
	if err != nil {
		log.Fatal("yaml conversion error")
	}
	if err := os.WriteFile("data.txt", file, 0644); err != nil {
		log.Fatal("file save err:", err)
	}
	log.Printf("data persistence successful")
	close(ExitCh)
}

type rpcServerProcess struct{}

func (s *rpcServerProcess) Create(ctx context.Context, req *proto.CreateReq) (*proto.CreateResp, error) {
	//todo：成熟时应该添加某种程度的校验和鉴权
	newFunc := &function{
		Name:       req.Request.Name,
		Label:      req.Request.Labels,
		Annotation: req.Annotations,
		Document:   req.Document,
	}

	functionPath := functionDir + string(os.PathSeparator) + buildName(newFunc) //此处的文件名使用name-label-annotation的格式构造
	funcYaml[req.Request.Name] = newFunc
	err := os.WriteFile(functionPath, []byte(req.Document), 0666)
	if err != nil {
		log.Errorln("save yaml err:", err)
		return nil, fmt.Errorf("save yaml err:%v", err)
	}
	return &proto.CreateResp{
		RequestId:        req.Request.RequestId,
		StatusCode:       0,
		Description:      "OK",
		ErrorInformation: "",
	}, nil
}

func (s *rpcServerProcess) Get(ctx context.Context, req *proto.GetReq) (*proto.GetResp, error) {
	targetFunc, ok := funcYaml[req.Request.Name]
	if !ok {
		return nil, fmt.Errorf("functione not found")
	}
	return &proto.GetResp{
		RequestId:  req.Request.RequestId,
		StatusCode: 0,
		Document:   targetFunc.Document,
	}, nil

}
func (s *rpcServerProcess) Delete(ctx context.Context, req *proto.DeleteReq) (*proto.DeleteResp, error) {
	f, ok := funcYaml[req.Request.Name]
	if !ok {
		return nil, fmt.Errorf("functione not found")
	}
	err := os.Remove(functionDir + string(os.PathSeparator) + buildName(f))
	if err != nil {
		log.Errorln("function delete err:", err)
		return nil, fmt.Errorf("function delete err:%v", err)
	}
	delete(funcYaml, req.Request.Name)
	return &proto.DeleteResp{
		RequestId:        req.Request.RequestId,
		StatusCode:       0,
		Description:      "OK",
		ErrorInformation: "",
	}, nil
}

func (s *rpcServerProcess) List(ctx context.Context, req *proto.ListReq) (*proto.ListResp, error) {
	var funcName []string
	funcName = maps.Keys(funcYaml)
	yamlData, err := yaml.Marshal(funcName)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, fmt.Errorf("yaml conversion error")
	}
	return &proto.ListResp{
		RequestId:  req.RequestId,
		StatusCode: 0,
		List:       string(yamlData),
	}, nil
}
func buildName(f *function) string {
	res := f.Name
	if f.Label != "" {
		res += "-" + f.Label
	}
	if f.Annotation != "" {
		if f.Label == "" {
			res += "-"
		}
		res += "-" + f.Annotation
	}
	return res
}
func isIPAddress(addr string) bool {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return false
	}
	return net.ParseIP(ip) != nil
}
func getLocalIPv4() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println(ipNet.IP.String())
				return ipNet.IP
			}
		}
	}
	return nil
}
