package process

import (
	"fmt"
	"os"
	"net"
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
)

//顯示登入成功後的介面
func ShowMenu() {
	fmt.Println("----------xxx登錄成功----------")
	fmt.Println("----------1. 顯示線上用戶列表----------")
	fmt.Println("----------2. 發送訊息----------")
	fmt.Println("----------3. 查看訊息列表----------")
	fmt.Println("----------4. 退出系統----------")
	fmt.Println("請選擇(1-4):")
	var key int
	var content string
	//SmsProcess實例
	SmsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
		case 1 :
			//fmt.Println("顯示線上用戶列表")
			outputOnlineUser()
		case 2 :
			//fmt.Println("發送訊息")
			fmt.Println("請輸入聊天內容:")
			fmt.Scanf("%s\n", &content)
			SmsProcess.SendGroupMes(content)
		case 3 :
			fmt.Println("查看訊息列表")
		case 4 :
			fmt.Println("退出系統")
			os.Exit(0)
		default :
			fmt.Println("輸入錯誤")
	}
}

//和SERVER保持通訊
func serverProcessMes(conn net.Conn) {
	//創建一個transfer實例 不停讀取SERVER發送的訊息
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Println("正在等待讀取SERVER發送的訊息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err", err)
			return
		}
		//如果讀取到訊息
		//fmt.Println("mes=", mes)
		//判斷收到的message類型
		switch mes.Type {
			case message.NotifyUserStatusMesType : //有人上線了
				//1.取出NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				//2.把這個用戶的狀態保存到客戶端的map[int]User中
				updateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType : //有人發送消息了
				outputGroupMes(&mes)
			default :
				fmt.Println("服務器端返回未知類型")
		}
	}
}