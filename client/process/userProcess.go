package process

import (
	"fmt"
	"encoding/json"
	"encoding/binary"
	"net"
	"chatroom/common/message"
	"chatroom/client/utils"
	"os"
)

type UserProcess struct {
	//字段

}

//用戶註冊
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	
	//1. 連接到SERVER
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net Dial err=", err)
		return
	}
	//延遲關閉
	defer conn.Close()

	//2. 準備通過conn發送消息給SERVER
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3. 創建一個LoginMes結構體
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 將RegisterMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}

	//5. 把Data賦給 mes.Data字段
	mes.Data = string(data)

	//6. mes進行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}

	tf := &utils.Transfer {
		Conn : conn,
	}
	//發送Data給SERVER
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("註冊發送package err", err)
	}
	
	mes, err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) err", err)
		return
	}

	//將mes的Data反序列化成RegisterResMes
	var registerResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("註冊成功 請重新登入")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}


//關聯一個用戶登陸的方法
//寫一個函數 完成登錄
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//下一個就要開始訂協議
	//fmt.Printf("userId = %d userPwd = %s", userId, userPwd)
	//return nil

	//1. 連接到SERVER
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net Dial err=", err)
		return
	}
	//延遲關閉
	defer conn.Close()

	//2. 準備通過conn發送消息給SERVER
	var mes message.Message
	mes.Type = message.LoginMesType

	//3. 創建一個LoginMes結構體
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 將loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}

	//5. 把Data賦給 mes.Data字段
	mes.Data = string(data)

	//6. mes進行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}

	//7. data就是要發送的訊息
	//7.1 先把data長度發送給SERVER
	//先獲取data的長度 -> 轉成一個表示長度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//發送長度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err", err)
		return
	}
	
	fmt.Printf("客戶端 發送消息的長度=%d 內容=%s", len(data), string(data))
	
	//發送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err", err)
		return
	}
	//服務器端返回的訊息

	//測試休眠20秒
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠20")

	//處理SERVER返回的消息
	tf := &utils.Transfer {
		Conn : conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err", err)
		return
	}
	//將mes的Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//fmt.Println("登入成功")
		//可以顯示線上用戶列表
		fmt.Println("當前線上用戶如下:")
		for _, v := range loginResMes.UsersId {

			//如果要求自己不要顯示在線上
			if v == userId {
				continue
			}
			//完成客戶端的onlineUsers完成初始化
			user := &message.User{
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
			fmt.Println("用戶id:\t", v)
		}

		//這裡還需要啟動一個goroutine
		//保持和SERVER的通訊 如果SERVER有數據推送給Client
		//則可以接收並且顯示在Client上

		go serverProcessMes(conn)

		//1. 顯示登錄成功後的菜單(循環顯示)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}