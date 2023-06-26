package dto

type TensorboardStartArgs struct {
	TaskID int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TensorboardStartReply struct {
	TensorboardID  int64  `json:"tensorboard_id"`
	TensorboardUrl string `json:"tensorboard_url"`
}

type TensorboardRestartArgs struct {
	TaskID int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TensorboardStopArgs struct {
	TaskID int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TensorboardDetailArgs struct {
	TaskID int64 `form:"task_id" json:"task_id" binding:"required"`
}

type TensorboardDetailReply struct {
	TensorboardID   int64  `json:"tensorboard_id"`
	TensorboardName string `json:"tensorboard_name"`
	TensorboardUrl  string `json:"tensorboard_url"`
}

type TensorboardListArgs struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}

type TensorboardListReply struct{}
