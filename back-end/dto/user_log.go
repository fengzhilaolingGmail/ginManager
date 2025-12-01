// dto/oper_log.go
package dto

import "time"

type UserLogListReq struct {
	PageReq
	Module    string    `form:"module"`     // 模块模糊
	Username  string    `form:"username"`   // 用户模糊
	Status    uint8     `form:"status"`     // 0/1/<2 全部
	StartTime time.Time `form:"start_time"` // 开始时间
	EndTime   time.Time `form:"end_time"`   // 结束时间
}

type UserLogExportReq struct {
	Module    string    `form:"module"`
	Username  string    `form:"username"`
	Status    uint8     `form:"status"`
	StartTime time.Time `form:"start_time"`
	EndTime   time.Time `form:"end_time"`
}
