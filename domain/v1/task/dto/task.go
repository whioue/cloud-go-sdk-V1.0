package dto

type TaskCreateArgs struct {
	Name      string            `form:"name" json:"name" binding:"required"`
	Actor     *TaskBaseResource `form:"actor" json:"actor"`
	Predictor *TaskBaseResource `form:"predictor" json:"predictor"`
	Learner   *TaskBaseResource `form:"learner" json:"learner"`
}

type TaskCreateReply struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type TaskDeleteArgs struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type TaskDetailArgs struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type TaskDetailReply struct {
	ID           int64  `json:"id"`
	Uuid         string `json:"uuid"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	TrainingTime int64  `json:"training_time"`
	Cpu          int64  `json:"cpu"`
	Gpu          int64  `json:"gpu"`
	Memory       int64  `json:"memory"`
	CommonImage  string `json:"commonImage,omitempty"`
	ActorImage   string `json:"actorImage,omitempty"`
}

type TaskListArgs struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}

type TaskListReply struct {
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	List     []*TaskDetailReply `json:"list"`
	TotalCnt int64              `json:"total_cnt"`
}

type TaskUpdateArgs struct {
	ID          int64  `form:"id" json:"id" binding:"required"`
	Name        string `form:"name" json:"name"`
	CommonImage string `form:"common_image" json:"common_image"`
	ActorImage  string `form:"actor_image" json:"actor_image"`
}

type TaskBaseResource struct {
	Number int   `form:"number" json:"number"`
	Cpu    int64 `form:"cpu" json:"cpu"`
	Gpu    int64 `form:"gpu" json:"gpu"`
	Memory int64 `form:"memory" json:"memory"`
}
