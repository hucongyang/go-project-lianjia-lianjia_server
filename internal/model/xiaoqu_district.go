package model

type XiaoquDistrict struct {
	*Model
	XiaoquId string `json:"xiaoqu_id"`
	DistrictId string `json:"district_id"`
}

func (xiaoquDistrict *XiaoquDistrict) TableName() string {
	return "lj_xiaoqu_district"
}