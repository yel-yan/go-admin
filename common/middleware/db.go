package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yel-yan/go-admin-core/sdk"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", sdk.Runtime.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
