package process2

import (
	"fmt"
	"net"
	"chatroom/common/message"
	"chatroom/server/utils"
	"chatroom/server/model"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
	//增加一個字段 表示該conn是哪個用戶
	UserId int
}

//編寫通知所有線上用戶的方法
//userId要通知其他用戶上線
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍歷onlineUsers 然後一個一個發送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//過濾自己
		if id == userId {
			continue
		}
		//開始通知
		up.NotifyMeOnlineUser(userId)
	}

}

func (this *UserProcess) NotifyMeOnlineUser(userId int) {
	//組裝NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//將notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	//將序列化後的mes賦給mes.Data
	mes.Data = string(data)

	//對mes再進行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}
	
	//發送, 創建一個Transfer實例
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnlineUser", err)
	}
}

//處理註冊需求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	
	//1. 先從mes中 取出mes.Data 並反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}

	//1.先聲明一個resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知錯誤"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//4.將data賦值給resMes
	resMes.Data = string(data)

	//5.對resMes進行序列化 準備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//6.發送data 將其封裝到writePkg
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//編寫一個函數serverProcessLogin 專門處理登錄請求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1. 先從mes中 取出mes.Data 並反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	
	//1.先聲明一個resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.在聲明一個 LoginResMes 並完成賦值
	var loginResMes message.LoginResMes

	//到redis去完成驗證
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "未知錯誤"
		}
		
	} else {
		loginResMes.Code = 200
		//用戶登錄成功 於是把該登錄成功的用戶放入userMgr中
		//將登錄成功的userId賦給this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他在線用戶 用戶上線了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//將當前線上用戶的ID放入loginResMes.UsersId
		//遍歷 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println(user, "登入成功")
	}

	// //如果用戶id = 100 密碼 = 123456 就正確 否則不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	//不存在
	// 	loginResMes.Code = 500 //500狀態碼表示該用戶不存在
	// 	loginResMes.Error = "用戶不存在"
	// }

	//3.將loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//4.將data賦值給resMes
	resMes.Data = string(data)

	//5.對resMes進行序列化 準備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//6.發送data 將其封裝到writePkg
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}