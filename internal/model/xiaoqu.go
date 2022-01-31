package model

import "github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"

type Xiaoqu struct {
	*Model
	XiaoquId       string `json:"xiaoqu_id"`
	Url            string `json:"url"`
	BuildDate      string `json:"build_date"`
	ChengJiaoText  string `json:"cheng_jiao_text"`
	ChengJiaoUrl   string `json:"cheng_jiao_url"`
	Img            string `json:"img"`
	HouseSellUrl   string `json:"house_sell_url"`
	TagList        string `json:"tag_list"`
	Title          string `json:"title"`
	TotalPrice     string `json:"total_price"`
	TotalSellCount string `json:"total_sell_count"`
	Status         string `json:"status"`
}

func (xiaoqu Xiaoqu) TableName() string {
	return "lj_xiaoqu"
}

type XiaoquSwagger struct {
	List  []*Xiaoqu
	Pager *app.Pager
}
