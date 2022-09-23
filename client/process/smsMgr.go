package process

import (
	"fmt"
	"chatroom/common/message"
	"encoding/json"
)

func outputGroupMes(mes *message.Message) {
	//顯示出來即可
	

	//1.反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}

	//2.顯示訊息
	info := fmt.Sprintf("用戶id:\t%d 對大家說:\t%s",
			smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
