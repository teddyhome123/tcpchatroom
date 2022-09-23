package main

import (
	"fmt"
	"net"
	"time"
	"chatroom/server/model"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	buf := make([]byte, 8096)
// 	fmt.Println("等待讀取客戶端發送的數據")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//fmt.Println("conn.Read err=", err)
// 		//err = errors.New("read pkg header error")
// 		return
// 	}
// 	//根據buf[:4]轉成一個uint32類型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])
// 	//根據pkgLen讀取消息內容
// 	//是指conn讀取多少字節 > buf
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//fmt.Println("conn.Read err", err)
// 		//err = errors.New("read pkg body error")
// 		return
// 	}
// 	//把pkgLen反序列化成 -> message.Message
// 	//&mes !
// 	json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err=", err)
// 		return
// 	}

// 	return

// }

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//先發送一個長度給對方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	//發送長度
// 	n, err := conn.Write(buf[0:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes) err", err)
// 		return
// 	}

// 	//發送data本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(bytes) err", err)
// 		return
// 	}
// 	return
// }


// //編寫一個函數serverProcessLogin 專門處理登陸請求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//1. 先從mes中 取出mes.Data 並反序列化成LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err", err)
// 		return
// 	}
	
// 	//1.先聲明一個resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	//2.在聲明一個 LoginResMes 並完成賦值
// 	var loginResMes message.LoginResMes

// 	//如果用戶id = 100 密碼 = 123456 就正確 否則不合法
	
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		loginResMes.Code = 200
// 	} else {
// 		//不存在
// 		loginResMes.Code = 500 //500狀態碼表示該用戶不存在
// 		loginResMes.Error = "用戶不存在"
// 	}

// 	//3.將loginResMes序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err", err)
// 		return
// 	}

// 	//4.將data賦值給resMes
// 	resMes.Data = string(data)

// 	//5.對resMes進行序列化 準備發送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err", err)
// 		return
// 	}

// 	//6.發送data 將其封裝到writePkg
// 	err = writePkg(conn, data)
// 	return
// }

//編寫一個ServerProcessMes 
//功能:根據客戶端發送的消息種類 決定調用哪個函數處理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 		case message.LoginMesType :
// 			//處理登入邏輯
// 			err = serverProcessLogin(conn, mes)
// 		case message.RegisterMesType : 
// 			//處理註冊
// 		default :
// 			fmt.Println("消息類型不存在")
// 	}
// 	return
// }

//處理和客戶端的通訊
func process(conn net.Conn) {
	//延遲關閉
	defer conn.Close()
	
	//這裡調用總控 創建一個
	processor := &Prcoessor {
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客戶端和服務器端通訊的錯誤 err", err)
		return
	}
}

//完成對UserDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main(){
	//當SERVER啟動時就去初始化redis連接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
	//提示訊息
	fmt.Println("Server在8889端口監聽")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
	}
	//一旦監聽成功就等待客戶端來連接
	for {
		fmt.Println("等待客戶端連接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept err", err)
		}
		//一旦連接成功 則啟動一個協程和客戶端保持數據的通訊
		go process(conn)
	}
}