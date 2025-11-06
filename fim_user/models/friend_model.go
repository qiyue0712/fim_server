package models

import "fim_server/common/models"

type FriendModel struct {
	models.Model
	SendUserID uint   `json:"sendUserID"`
	RevUserID  uint   `json:"revUserID"`
	Notice     string `json:"notice"` // 备注
}
