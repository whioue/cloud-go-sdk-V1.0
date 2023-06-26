package dto

import "time"

type LogDeleteArgs struct{}

type LogDownloadArgs struct{}

type LogDownloadReply struct{}

type LogListArgs struct {
	Page     int   `form:"page" json:"page" binding:"required"`
	PageSize int   `form:"page_size" json:"page_size" binding:"required"`
	TaskID   int64 `form:"task_id" json:"task_id" binding:"required"`
}

type LogListReply struct {
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	List     []*LogDetailReply `json:"list"`
	TotalCnt int64             `json:"total_cnt"`
}

type LogDetailArgs struct{}

type LogDetailReply struct {
	FileName   string    `json:"file_name"`
	CreateTime time.Time `json:"create_time"`
	Size       int64     `json:"size"`
}
