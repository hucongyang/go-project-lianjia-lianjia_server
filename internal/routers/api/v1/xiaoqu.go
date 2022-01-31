package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/app"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/errcode"
)

type Xiaoqu struct{}

func NewXiaoqu() Xiaoqu {
	return Xiaoqu{}
}

func (x Xiaoqu) Get(c *gin.Context) {}
func (x Xiaoqu) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (x Xiaoqu) Create(c *gin.Context) {}
func (x Xiaoqu) Update(c *gin.Context) {}
func (x Xiaoqu) Delete(c *gin.Context) {}
