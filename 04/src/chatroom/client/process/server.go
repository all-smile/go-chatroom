// 保持跟服务端的通信
package process

import (
	"fmt"
	"go_code/chatroom/client/utils"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("***************** 欢迎XXX登录多人聊天系统 *****************")
	fmt.Println("***************** 1. 显示在线用户列表 *****************")
	fmt.Println("***************** 2. 发送消息 *****************")
	fmt.Println("***************** 3. 信息列表 *****************")
	fmt.Println("***************** 4. 退出系统 *****************")
	var key int
	fmt.Println("\t\t\t 请选择(1-4):")
	fmt.Scanf("%d \n", &key)
	switch key {
	case 1:
		fmt.Println("在线用户列表--")
	case 2:
		fmt.Println("发送消息--")
	case 3:
		fmt.Println("信息列表--")
	case 4:
		fmt.Println("退出系统--")
		os.Exit(0)
	default:
		fmt.Println("输入有误，重新输入")
	}
}

// 监听 服务端消息 并显示
func ServerProcessMes(conn net.Conn) {
	// 创建 Transfer 实例， 等待信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端： %v ,正在监听服务端消息 \n", conn.RemoteAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端读取信息失败---", err)
			return
		}
		fmt.Println("服务端推送的消息 = ", mes)
	}
}
