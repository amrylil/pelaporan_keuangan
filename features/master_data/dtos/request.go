package dtos

type InputMaster_data struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}