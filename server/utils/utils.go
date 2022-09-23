package utils

import (
	"fmt"
	"net"
	"chatroom/common/message"
	"encoding/json"
	"encoding/binary"
)

//
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}


func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	fmt.Println("等待讀取客戶端發送的數據")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err=", err)
		//err = errors.New("read pkg header error")
		return
	}
	//根據buf[:4]轉成一個uint32類型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根據pkgLen讀取消息內容
	//是指conn讀取多少字節 > buf
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn.Read err", err)
		//err = errors.New("read pkg body error")
		return
	}
	//把pkgLen反序列化成 -> message.Message
	//&mes !
	json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return

}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先發送一個長度給對方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//發送長度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err", err)
		return
	}

	//發送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) err", err)
		return
	}
	return
}