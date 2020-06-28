package error

import (
	"gintest/config"
	"gintest/util/json"
	"log"
	"os"
	"time"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func NewError (text string, level string) error {
	alarm(level, text)
	return &errorString{text}
}


// 告警方法
func alarm(level string, str string) {
	// 执行记日志
	if f, err := os.OpenFile(config.AppConfig.LogDir + "-error.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		log.Println(err)
	} else {
		errorLogMap := make(map[string]interface{})
		errorLogMap["time"] = time.Now().Format("2006/01/02 - 15:04:05")
		errorLogMap[level] = str

		errorLogJson, _ := json.Encode(errorLogMap)
		_, _ = f.WriteString(errorLogJson + "\n")
	}

}
