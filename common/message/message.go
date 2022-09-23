package message

const (
	LoginMesType 			= "LoginMes"
	LoginResMesType 		= "LoginResMes"
	RegisterMesType 		= "RegisterMes"
	RegisterResMesType	 	= "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType 				= "SmsMes"
)

//定義幾個用戶狀態的變數
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`  //消息類型
	Data string `json:"data"`  //消息內容
}

//定義兩個消息 後面需要再增加
type LoginMes struct {
	UserId int `json:"userId"`  //用戶的ID
	UserPwd string `json:"userPwd"`  //用戶的密碼
	UserName string `json:"userName"`  //用戶名
}

type LoginResMes struct {
	Code int `json:"code"`  //返回一個狀態碼 500表示用戶未註冊 200表示登入成功
	UsersId []int 						//增加一個保存用戶ID的切片
	Error string `json:"error"`  //返回的錯誤訊息
}

type RegisterMes struct {
	User User `json:"user"` 
}

type RegisterResMes struct {
	Code int `json:"code"`  //返回一個狀態碼400表示已有該用戶 200表示註冊成功
	Error string `json:"error"` //返回的錯誤訊息
}

//為了配合SERVER發送用戶狀態變化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用戶ID
	Status int `json:"status"` //用戶狀態
}

//增加一個SmsMes 發送消息
type SmsMes struct {
	Content string 	`json:"content"` //內容
	User //匿名結構體 繼承
}