package exception

import (
	"fmt"
	"gintest/config"
	"gintest/util/error"
	jsonUtil "gintest/util/json"
	"gintest/util/response"
	"gintest/util/time"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.AppConfig.Name)
				// 日志格式
				errorLogMap := make(map[string]interface{})

				errorLogMap["error_msg"] = fmt.Sprintf("%s", err)
				errorLogMap["request_time"] = time.GetCurrentDate()
				errorLogMap["request_method"] = c.Request.Method
				errorLogMap["request_uri"] = c.Request.RequestURI
				errorLogMap["request_proto"] = c.Request.Proto
				errorLogMap["request_ua"] = c.Request.UserAgent()
				errorLogMap["request_client_ip"] = c.ClientIP()
				//errorLogMap["debug_stack"] = DebugStack


				errorLogJson, _ := jsonUtil.Encode(errorLogMap)

				_ = error.NewError(errorLogJson, "ERROR")

				utilGin := response.Gin{Ctx: c}
				utilGin.Response(500, subject, err)
			}
		}()
		c.Next()
	}
}
