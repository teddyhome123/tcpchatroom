package process2

import (
	"fmt"
	"encoding/json"
	"chatroom/common/message"
	"chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

//寫方法轉發消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍歷SERVER的UserMgr map[int]*UserProcess
	//將消息轉發出去
	
	//取出mes的內容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}

	for id, up := range userMgr.onlineUsers {
		//這裡還需要過濾自己 不要重新發給自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}

}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	//創建一個Transfer 發送data
	tf := &utils.Transfer {
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("轉發消息失敗 err=", err)
	}
}