package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render404 -
func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "很抱歉！您浏览的页面不存在。")
}