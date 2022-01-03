// data access object
package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

// 定义 UserDao 结构体， 完成对 User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建 UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1. 根据用户 id 返回一个 User 实例
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过 id 去redis 查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEXISTS
			fmt.Printf("该用户: %v 不存在\n", id)
			return
		}
	}
	fmt.Println("res = ", res)
	// 反序列化 res
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	return
}

// 2. 完成登录校验
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
