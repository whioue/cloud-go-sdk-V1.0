package dto

type TrainStartArgs struct {
	TaskID int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TrainStopArgs struct {
	TaskID int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TrainLogArgs struct {
	PodName string `form:"pod_name" json:"pod_name" binding:"required"`
}

type TrainLogReply struct {
	Log string `json:"log"`
}

type TrainDetailReply struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TrainListArgs struct {
	Page     int   `form:"page" json:"page"`
	PageSize int   `form:"page_size" json:"page_size"`
	TaskID   int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TrainListReply struct {
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	List     []*TrainDetailReply `json:"list"`
	TotalCnt int64               `json:"total_cnt"`
}
