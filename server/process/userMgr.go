package process2

import (
	"fmt"
)

//UserMgr 在SERVER只有一個 且很多地方都會使用到
//將其定義為全局變量
var (
	userMgr *UserMgr
)


type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成對UserMgr初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess, 1024),
	}
}

//完成對onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	
	this.onlineUsers[up.UserId] = up 
}

//完成對onlineUsers刪除
func (this *UserMgr) DelOnlineUser(userId int) {
	//刪除
	delete(this.onlineUsers, userId)
}

//返回當前線上所有用戶
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根據Id返回一個對應的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何從map取出一個值 帶檢測的方式
	up, ok := this.onlineUsers[userId]
	if !ok { //說明用戶不在線上
		err = fmt.Errorf("用戶%d不存在", userId)
		return
	}
	return
}