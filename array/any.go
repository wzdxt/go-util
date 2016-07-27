package array

import (
	"github.com/wzdxt/go-util/set"
	"math"
)

func Union(args ...[]interface{}) []interface{} {
	s := make(set.Set)
	for _, arr := range args {
		s.AddArr(arr)
	}
	return s.Elements()
}

func Minus(arr []interface{}, args ...[]interface{}) []interface{} {
	s := make(set.Set, len(arr)).AddArr(arr)
	for _, arg := range args {
		s.RemoveArr(arg)
	}
	return s.Elements()
}

func Intersect(arr []interface{}, args ...[]interface{}) []interface{} {
	remain := make([]interface{}, 0, len(arr))
	copy(remain, arr)
	for _, arg := range args {
		r := make([]interface{}, 0, math.Min(len(remain), len(arg)))
		s := make(set.Set, len(arg)).AddArr(arg)
		for _, v := range remain {
			if s.Has(v) {
				r = append(r, v)
			}
		}
		remain = r
	}
	return remain
}
