package main

import (
	/*
		#include <stdlib.h>
		void clear() {
			system("pause");
			system("cls");
		}
	*/
	"C"
	"fmt"
)

type View struct {
	choice  uint8
	curUser *client
}

func NewView(curUser *client) *View {
	view := &View{255, curUser}
	return view
}

func (this *View) showMenu() bool {
	choices := 255
	fmt.Println("------------------------------")
	fmt.Println("         1. 私聊模式           ")
	fmt.Println("         2. 群聊模式           ")
	fmt.Println("         3. 修改名称           ")
	fmt.Println("         0. 退出系统           ")
	fmt.Println("------------------------------")
	fmt.Scanln(&choices)
	if choices < 0 || choices > 3 {
		fmt.Println("请输入合法范围内的数字")
		return false
	} else {
		this.choice = uint8(choices)
		return true
	}
}

func (this *View) run() {
	for this.choice != 0 {
		for this.showMenu() == false {
			C.clear()
		}
		switch this.choice {
		case 1:
			this.curUser.privateChat()
		case 2:
			this.curUser.publicChar()
		case 3:
			this.curUser.ModifyUserName()
		}
		C.clear()
	}
}
