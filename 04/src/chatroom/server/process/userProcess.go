package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

// 登录逻辑处理
func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	// 定义返回消息类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	// 校验信息
	// redis数据库校验
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		// 组织错误信息
		if err == model.ERROR_USER_NOEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Printf("用户id: %v, 登录成功 \n", user.UserId)
	}

	/* if loginMes.UserId == 100 && loginMes.UserPwd == "abc" {
		loginResMes.Code = 200
		loginResMes.Error = "登录成功"
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
	} */

	// 序列化信息
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	// 发送
	/*
		MVC 分层
		创建 Transfer 实例， 进行读取
	*/
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}
