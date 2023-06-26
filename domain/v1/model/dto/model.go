package dto

import "time"

type ModelDeleteArgs struct{}

type ModelDownloadArgs struct{}

type ModelDownloadReply struct{}

type ModelDetailReply struct {
	FileName   string    `json:"file_name"`
	CreateTime time.Time `json:"create_time"`
	Size       int64     `json:"size"`
}

type ModelListArgs struct {
	Page     int    `form:"page" json:"page" binding:"required"`
	PageSize int    `form:"page_size" json:"page_size" binding:"required"`
	TaskID   int64  `form:"task_id" json:"task_id" binding:"required"`
	Catalog  string `form:"catalog" json:"catalog" binding:"required"`
}

type ModelListReply struct {
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	List     []*ModelDetailReply `json:"list"`
	TotalCnt int64               `json:"total_cnt"`
}

type ModelCatalogsDetailReply struct {
	DirName string `json:"dir_name"`
}

type ModelCatalogsArgs struct {
	Page     int   `form:"page" json:"page"`
	PageSize int   `form:"page_size" json:"page_size"`
	TaskID   int64 `form:"task_id" json:"task_id" binding:"required"`
}

type ModelCatalogsReply struct {
	Page     int                         `json:"page"`
	PageSize int                         `json:"page_size"`
	List     []*ModelCatalogsDetailReply `json:"list"`
	TotalCnt int64                       `json:"total_cnt"`
}
