package process

import (
	"fmt"
	"chatroom/common/message"
	"chatroom/client/model"

)

//客戶端要維護的Map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在用戶成功後 完成對CurUser的初始化

//在客戶端顯示當前在線的用戶
func outputOnlineUser() {

	//遍歷onlineUsers
	fmt.Println("線上用戶列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用戶id:\t", id)

	}
}

//編寫一個方法 處理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原來map沒有這個用戶
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
			UserStatus : notifyUserStatusMes.Status,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}