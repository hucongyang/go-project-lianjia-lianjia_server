package model

type District struct {
	*Model
	Name string `json:"name"`
	Quanpin string `json:"quanpin"`
	LianjiaDistrictId int `json:"lianjia_district_id"`
	Status string `json:"status"`
}

func (district District) TableName() string {
	return "lj_district"
}
