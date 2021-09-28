package commons

import (
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/pkg/errors"
	"reflect"
)

// Contains returns true if an element is present in a iteratee.
func Contains(arr interface{}, obj interface{}) (bool, error) {
	if isNotArray(arr) {
		return false, errors.Errorf("arr is not array or slice")
	}
	return linq.From(arr).Contains(obj), nil
}

func Split(arr []interface{}, size int) ([][]interface{}, error) {
	if size <= 0 {
		return nil, errors.Errorf("size should greater than 0, actually:%v", size)
	}

	if arr == nil {
		return nil, errors.Errorf("arr can not be nil")
	}

	length := len(arr)
	res := make([][]interface{}, 0, length/size+1)
	if length == 0 {
		return res, nil
	}

	for idx, item := range arr {
		childArrIdx := idx / size
		childArr := res[childArrIdx]
		if childArr == nil {
			childArr = make([]interface{}, 0, size)
			res = append(res, childArr)
		}
		childArr = append(childArr, item)
	}

	return res, nil
}

func isNotArray(arr interface{}) bool {
	t := reflect.ValueOf(arr)
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		return false
	default:
		return true
	}
}

func Intersect(x interface{}, y interface{}, res interface{}) error {
	if isNotArray(x) || isNotArray(y) || isNotArray(res) {
		return errors.New("x,y,res must be array")
	}

	xType := reflect.ValueOf(x).Type()
	yType := reflect.ValueOf(y).Type()
	zType := reflect.ValueOf(res).Type()
	if NotEqual(xType, yType) || NotEqual(xType, zType) {
		return errors.Errorf("x,y,res type must equal, actually:%v,%v,%v", xType, yType, zType)
	}
	linq.From(xType).
		IntersectBy(linq.From(yType),
			func(p interface{}) interface{} { return p },
		).
		ToSlice(res)
	return nil
}
