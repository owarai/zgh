package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/owarai/zgh/log"
)

type validate interface {
	Message() map[string]int
}

func (g *Gin) Validate(obj validate) bool {
	valid := validation.Validation{}
	b, err := valid.Valid(obj)
	if err != nil {
		log.L().Error("message", "valid error", "err", err.Error())
		g.Response(http.StatusOK, 400000000, nil)
		return false
	}

	if !b {
		errorMaps := obj.Message()
		field := valid.Errors[0].Key
		if v, ok := errorMaps[field]; ok {
			g.Response(http.StatusOK, v, nil)
			return b
		}
		g.Response(http.StatusOK, 100000001, nil)
		return b
	}
	return true
}
