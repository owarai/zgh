package hashid

import (
	"github.com/speps/go-hashids"

	"github.com/owarai/zgh"
	"github.com/owarai/zgh/conf"
)

type HashIdParams struct {
	Salt      string
	MinLength int
}

var hashIdParams *HashIdParams

func (hd *HashIdParams) SetHashIdSalt(salt string) func(*HashIdParams) interface{} {
	return func(hd *HashIdParams) interface{} {
		hs := hd.Salt
		hd.Salt = salt
		return hs
	}
}

func (hd *HashIdParams) SetHashIdLength(minLength int) func(*HashIdParams) interface{} {
	return func(hd *HashIdParams) interface{} {
		ml := hd.MinLength
		hd.MinLength = minLength
		return ml
	}
}

func (hd *HashIdParams) HashIdInit(options ...func(*HashIdParams) interface{}) (*hashids.HashID, error) {
	q := &HashIdParams{
		Salt:      conf.HASHIDSALT,
		MinLength: conf.HASHIDMINLENGTH,
	}
	for _, option := range options {
		option(q)
	}
	hashIdParams = q
	hds := hashids.NewData()
	hds.Salt = hashIdParams.Salt
	hds.MinLength = hashIdParams.MinLength
	h, err := hashids.NewWithData(hds)
	if err != nil {
		zgh.ZLog().Error("content", "hash new with data is error", "error", err.Error())
		return nil, err
	}
	return h, nil
}
