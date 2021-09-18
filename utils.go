package commons

import (
	"bytes"
	"reflect"
)

func Equal(obj1 interface{}, obj2 interface{}) bool {
	if obj1 == nil || obj2 == nil {
		return obj1 == obj2
	}

	if obj1Byte, ok := obj1.([]byte); ok {
		obj2Byte, ok := obj2.([]byte)
		if !ok {
			return false
		}

		if obj1Byte == nil || obj2Byte == nil {
			return true
		}

		return bytes.Equal(obj1Byte, obj2Byte)
	}

	return reflect.DeepEqual(obj1, obj2)
}

func NotEqual(obj1 interface{}, obj2 interface{}) bool {
	return !Equal(obj1, obj2)
}


