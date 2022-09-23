package process

import (
	"fmt"
	"chatroom/common/message"
	"chatroom/client/utils"
	"encoding/json"
)

type SmsProcess struct {

}

//發送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1. 創建一個Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.創建一個SmsMes 實例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	//4.再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}

	//5.將mes發送給SERVER
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	//6.發送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err", err)
		return
	}
	
	return
}