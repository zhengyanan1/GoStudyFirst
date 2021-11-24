package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
	flag int // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client{
//	创建客户端对象
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		flag: -1,
	}
//	连接server
	 conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	 if err != nil{
	 	fmt.Println("net.Dial error", err)
	 	return nil
	 }
	 client.conn = conn
//	返回对象
	return client
}

func (client *Client) menu() bool{
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3{
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>请输入合法范围内的数字<<<")
		return false
	}
}

func (client *Client) Run(){
	for client.flag != 0{
		for client.menu() != true {
		}
	//	根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			fmt.Println("choose 公聊")
			break
		case 2:
			fmt.Println("choose 私聊")
			break
		case 3:
			fmt.Println("choose 更新用户名")
			break
		}
	}
}



var serverIp string
var serverPort int

//./client -ip 127.0.0.1
func init(){
	flag.StringVar(&serverIp, "ip","127.0.0.1", "设置服务器ip地址")
	flag.IntVar(&serverPort, "port",8888, "设置服务器端口")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil{
		fmt.Println("连接服务器失败...")
		return
	}

	fmt.Println("连接服务器成功...")

//	启动客户端的业务
	client.Run()
}