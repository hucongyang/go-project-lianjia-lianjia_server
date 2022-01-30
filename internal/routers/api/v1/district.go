package v1

import "github.com/gin-gonic/gin"

type District struct{}

func NewDistrict() District {
	return District{}
}

func (d District) Get(c *gin.Context) {}
func (d District) List(c *gin.Context) {}
func (d District) Create(c *gin.Context) {}
func (d District) Update(c *gin.Context) {}
func (d District) Delete(c *gin.Context) {}