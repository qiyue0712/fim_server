package models

import "fim_server/common/models"

type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                  `json:"sendUserID"`
	RevUserID            uint                  `json:"revUserID"`
	Status               int8                  `json:"status"`             // 状态 0 未操作 1 同意 2 拒绝 3 忽略
	AdditionalMessages   string                `json:"additionalMessages"` // 附加消息
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion"`
}
