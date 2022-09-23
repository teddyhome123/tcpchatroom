package model

import (
	"net"
	"chatroom/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}