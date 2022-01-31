package model

import "github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"

type District struct {
	*Model
	Name              string `json:"name"`
	Quanpin           string `json:"quanpin"`
	LianjiaDistrictId int    `json:"lianjia_district_id"`
	Status            string `json:"status"`
}

func (district District) TableName() string {
	return "lj_district"
}

type DistrictSwagger struct {
	List  []*District
	Pager *app.Pager
}
