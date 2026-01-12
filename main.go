package main

import (
	"fim_server/core"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_models"
	"flag"
	"fmt"
)

type Options struct {
	DB bool
}

func main() {
	var opts Options
	flag.BoolVar(&opts.DB, "db", false, "db")
	flag.Parse()

	if opts.DB {
		db := core.InitGorm("root:gK3T9n%q2M@j7Z4@tcp(127.0.0.1:3307)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local")
		err := db.AutoMigrate(
			&user_models.UserModel{},
			&user_models.FriendModel{},
			&user_models.FriendVerifyModel{},
			&user_models.UserConfModel{},
			&chat_models.ChatModel{},
			&group_models.GroupModel{},
			&group_models.GroupMemberModel{},
			&group_models.GroupMsgModel{},
			&group_models.GroupVerifyModel{},
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功!")
	}
}
