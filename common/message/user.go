package message

//定義一個用戶的結構體

type User struct {
	//確定字串
	//為了序列化和反序列化成功 
	//必須保證Json的Key和結構體的字段對應的tag保持一致
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`
}