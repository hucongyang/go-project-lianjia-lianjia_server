package service

// 每个接口的参数各自定义一个校验参数
// form：表单的映射字段名
// binding：入参校验的规则内容

type CountDistrictRequest struct {
	Status uint8 `form:"status,default=1" binding:"oneof=0 1"`
}

type DistrictListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type CreateDistrictRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type UpdateDistrictRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,min=3,max=100"`
	Status     uint8  `form:"status,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteDistrictRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
