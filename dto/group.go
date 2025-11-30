// dto/group.go
package dto

type GroupListReq struct {
	PageReq
	GroupName string `form:"group_name"` // 模糊
}

type GroupAddReq struct {
	GroupCode string `json:"group_code" binding:"required,max=50"`
	GroupName string `json:"group_name" binding:"required,max=50"`
	Sort      int    `json:"sort"`
	Status    uint8  `json:"status" binding:"oneof=0 1"`
}
