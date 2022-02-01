package plugin

import (
  "github.com/gin-gonic/gin"	
)

type Plugin interface {
	Register(config map[string]interface{}, r *gin.RouterGroup) error
	Name() string
}
