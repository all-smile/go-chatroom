package message

// 定义消息类型常量
const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	// 消息类型
	Type string `json:"type"`
	// 消息内容
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	// 状态码： 500-未注册 200-成功 403-密码不正确
	Code int `json:"code"`
	// 具体错误信息
	Error string `json:"error"`
}

// 注册类型
type RegisterMes struct {
	User User `json:"user"`
	// UserName string
	// UserPwd  string
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
