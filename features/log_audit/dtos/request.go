package dtos

type InputLog_audit struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}