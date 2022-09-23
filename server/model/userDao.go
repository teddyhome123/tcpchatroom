package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"chatroom/common/message"
)

//在server啟動後就初始化一個userDao實例
//做成全局變量 在需要操作redis時 可直接使用
var (
	MyUserDao *UserDao
)

//定義一個UserDao結構體
//完成對User結構體的各種操作

type UserDao struct {
	pool  *redis.Pool
}

//使用工廠模式 創建一個UserDao實例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao {
		pool : pool,
	}
	return
}



//1. 根據用戶ID返回一個User實例+error
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	
	//通過給定的ID去redis查詢這個用戶
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//錯誤
		if err == redis.ErrNil { //表示在users hash中沒有此用戶
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//需要把res反序列化成User實例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
	}

	return
}

//註冊
func (this *UserDao) Register(user *message.User) (err error) {
	//先存UserDao的連接池取出一個連接
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err != nil { //如果為空 代表用戶已存在
		err = ERROR_USER_EXISTS //用戶已存在
		return
	}


	//說明該ID還沒註冊過 
	//序列化 結構體>字串
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//寫入redis
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存註冊用戶錯誤", err)
		return
	}
	return
}

//完成登錄的驗證
//1.完成用戶的驗證
//2.如果用戶id和pwd都正確 返回一個user實例
//3.如果用或id或pwd錯誤 返回一個錯誤訊息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先存UserDao的連接池取出一個連接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//到此證明用戶獲取成功 
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}




