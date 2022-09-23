package main

import (
	"fmt"
	"net"
	"chatroom/common/message"
	"chatroom/server/utils"
	"chatroom/server/process"
	"io"

)

type Prcoessor struct {
	Conn net.Conn
}

//編寫一個ServerProcessMes 
//功能:根據客戶端發送的消息種類 決定調用哪個函數處理
func (this *Prcoessor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType :
			//處理登入邏輯
			//創建一個UserProcess實例
			up := &process2.UserProcess {
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType : 
			//處理註冊
			up := &process2.UserProcess {
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.SmsMesType : 
			//創建一個SmsProcess實例完成轉發群聊消息
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default :
			fmt.Println("消息類型不存在")
	}
	return
}

func (this *Prcoessor) process2() (err error){
	//循環讀客戶端發送的訊息
	for {

		//將讀取數據 直接封裝成一個函數readpkg() 返回一個message, err
		//創建一個Transfer實例
		tf := &utils.Transfer {
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出 SERVER也退出..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		//fmt.Println("mes=", mes)

		err =  this.serverProcessMes(&mes)
		if err != nil {
			return err
		}

	}
}