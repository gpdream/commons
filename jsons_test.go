package commons

import (
	"testing"
)
import "github.com/stretchr/testify/assert"
type Test struct {
	Name string
	Age int32
}

func TestParseJson(t *testing.T) {
	jsonStr := "{\"Name\":\"wangba\",\"Age\":17}"
	o := &Test{}

	err := ParseJson(jsonStr, o)
	a := assert.New(t)
	a.Nil(err)
	a.True(o.Age == 17 && o.Name=="wangba")
}

func TestToJson(t *testing.T) {
	o := Test{
		Name: "wangba",
		Age:  17,
	}

	a := assert.New(t)
	jsonStr,err := ToJson(o)
	t.Logf("object:%v, jsonStr:%v", o, jsonStr)
	a.Nil(err)
}
