package leetcodego

import (
	"reflect"
	"testing"
)

func Test002Solution(t *testing.T) {
	res := Solution002([]int{1, 199, 30, 4}, 5)
	if !reflect.DeepEqual(res, []int{0, 3}) {
		t.Errorf("%v does not equal %v", res, []int{0, 3})
	}
}
