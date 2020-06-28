package test

import (
	"fmt"
	"gintest/util/response"
	"gintest/util/sign"
	"github.com/gin-gonic/gin"
	"time"
)

func Md5Test(c *gin.Context) {
	startTime  := time.Now()
	appSecret  := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		sign.MD5(appSecret + encryptStr + appSecret)

		// 验证签名
		sign.MD5(appSecret + encryptStr + appSecret)
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func AesTest(c *gin.Context) {
	startTime  := time.Now()
	appSecret  := "IgkibX71IEf382PT"
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 1000000
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := sign.Encrypt(encryptStr, []byte(appSecret), appSecret)

		// 验证签名
		sign.Decrypt(sn, []byte(appSecret), appSecret)
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}

func RsaTest(c *gin.Context) {
	startTime  := time.Now()
	encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
	count      := 500
	for i := 0; i < count; i++ {
		// 生成签名
		sn, _ := sign.PublicEncrypt(encryptStr, "rsa/public.pem")

		// 验证签名
		sign.PrivateDecrypt(sn, "rsa/private.pem")
	}
	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, fmt.Sprintf("%v次 - %v", count, time.Since(startTime)), nil)
}