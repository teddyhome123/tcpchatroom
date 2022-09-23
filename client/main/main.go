package main

import (
	"fmt"
	"os"
	"chatroom/client/process"
)

//定義兩個變量
//一個表示用戶ID 一個表示密碼
var userId int
var userPwd string
var userName string

func main() {
	//接收用戶的選擇
	var key int
	//判斷是否還繼續顯示菜單
	//var loop = true
	for true {
		fmt.Println("-------------歡迎登入多人聊天系統-------------")
		fmt.Println("\t\t\t 1 登陸聊天室")
		fmt.Println("\t\t\t 2 註冊用戶")
		fmt.Println("\t\t\t 3 退出系統")
		fmt.Println("\t\t\t 請選擇(1-3):")
		
		fmt.Scanf("%d\n", &key)
		switch key {
			case 1 :
				fmt.Println("登陸聊天室")
				fmt.Println("請輸入帳號(ID):")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("請輸入密碼:")
				fmt.Scanf("%s\n", &userPwd)
				//完成登錄
				//1.創建一個UserProcess的實例
				up := &process.UserProcess{}
				up.Login(userId, userPwd)
				//loop = false
			case 2 :
				fmt.Println("註冊用戶")
				fmt.Println("請輸入用戶ID:")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("請輸入用戶密碼:")
				fmt.Scanf("%s\n", &userPwd)
				fmt.Println("請輸入用戶名子:")
				fmt.Scanf("%s\n", &userName)
				up := &process.UserProcess{}
				up.Register(userId, userPwd, userName)
				//loop = false
			case 3 :
				fmt.Println("退出系統")
				//loop = false
				os.Exit(0)
			default :
				fmt.Println("輸入錯誤")
		}
	}

	//根據用戶的輸入顯示新的菜單
	// if key == 1 {
	// 	//說明用戶要登陸
	// 	fmt.Println("請輸入帳號(ID):")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("請輸入密碼:")
	// 	fmt.Scanf("%s\n", &userPwd)

	// 	//分層結構
	// 	//先把登入的函數寫到另一個文件 login.go
	// 	//login(userId, userPwd)
	// 	//if err != nil {
	// 	//	fmt.Println("登陸失敗")
	// 	//} else {
	// 	//	fmt.Println("登陸成功")
	// 	//}
	// } else if key == 2 {
	// 	fmt.Println("進行用戶註冊")
	// }
}