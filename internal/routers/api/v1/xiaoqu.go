package v1

import "github.com/gin-gonic/gin"

type Xiaoqu struct {}

func NewXiaoqu() Xiaoqu {
	return Xiaoqu{}
}

func (x Xiaoqu) Get(c *gin.Context) {}
func (x Xiaoqu) List(c *gin.Context) {}
func (x Xiaoqu) Create(c *gin.Context) {}
func (x Xiaoqu) Update(c *gin.Context) {}
func (x Xiaoqu) Delete(c *gin.Context) {}