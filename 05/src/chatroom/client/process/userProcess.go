package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

// 定义结构体, 并绑定方法
type UserProcess struct {
}

func (up *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接失败： ", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 组织消息
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 2.1. 序列化 registerMes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.Data = string(data)
	// 2.1. 序列化 mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	// 调用封装的 writePkg() 函数， 向服务端发送数据
	// 创建 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("客户端发送消息失败", err)
		return
	}

	// 处理服务端返回
	mes, err = tf.ReadPkg()
	fmt.Printf("mes 类型1： %T, 值： %v \n", mes, mes)
	if err != nil {
		fmt.Println("服务器处理失败", err)
		return
	}

	// 反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功， 请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	// 定协议
	/* fmt.Printf("userId = %d, userPwd = %s \n", userId, userPwd)
	return nil */
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("客户端连接失败： ", err)
		return
	}
	// 延时关闭
	defer conn.Close()
	// 2. 组织消息
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 2.1. 序列化 loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.Data = string(data)
	// 2.1. 序列化 mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	// 调用封装的 writePkg() 函数， 向服务端发送数据
	// 创建 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("客户端发送消息失败", err)
		return
	}

	// 处理服务端返回
	mes, err = tf.ReadPkg()
	fmt.Printf("mes 类型： %T, 值： %v \n", mes, mes)
	if err != nil {
		fmt.Println("服务器处理失败", err)
		return
	}
	// 反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")

		/* 再启动一个 goroutine, 该 goroutine 保持和服务端的通信，
		如果服务端有信息推送，则显示出来 */
		go ServerProcessMes(conn)

		// 1. 显示登陆后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
