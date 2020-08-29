package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/owarai/zgh/conf"
	"github.com/owarai/zgh/request"
	"github.com/owarai/zgh/utils/qq_captcha"
)

func TestRequest(t *testing.T) {
	resp := new(qq_captcha.QqCaptchaResponse)
	res, _, err := request.New().Get(conf.QCapUrl).
		Param("aid", "3333").
		Param("AppSecretKey", "232342").
		Param("Ticket", "23423").
		Param("Randstr", "234324").
		Param("UserIP", "127.0.0.1").
		Timeout(time.Minute * time.Duration(1)).Type(request.TypeUrlencoded).EndStruct(resp)
	fmt.Println(res, err)
}
