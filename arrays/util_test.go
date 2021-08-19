package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Item struct {
	Name string
	Age int32
}

func TestContains(t *testing.T) {
	arr := genItems(10)
	item1 := arr[0]
	contain,err := Contains(arr, item1)

	a := assert.New(t)
	a.True(contain)
	a.Nil(err)

	item2 := Item{
		Name: "liuqi",
		Age:  10,
	}

	contain, err = Contains(arr, item2)
	a.False(contain)
	a.Nil(err)

	contain, err = Contains(item1, item2)
	a.Nil(contain)
	a.NotNil(err)
}

func TestSplit(t *testing.T) {
	source := genItems(101)
	res,err := Split(source, 10)

	a := assert.New(t)
	a.True(len(res)==11)
	a.True(len(res[0])==10)
	a.Nil(err)

	res,err = Split(source, 102)
	a.True(len(res)==1)
	a.True(len(res[0])==1)
}

func genItems(size int) []Item {
	items := make([]Item, 0, size)
	for i:=0;i<size;i++ {
		item := Item{
			Name: "wangba",
			Age:  int32(i),
		}
		items = append(items, item)
	}

	return items
}
