package processwhx

import (
	"char_room/client/model"
	"char_room/common/message"
	"fmt"
)

//客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //我们在用户登录成功后，完成对CurUser初始化

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}
