package command

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

//MapKeys 获取map所有key
//
//@params:
//	m		map[K]V	map
//	order	int8	0不排序|1升序|2降序
//@return:
func MapKeys[K constraints.Ordered, V any](m map[K]V, order int8) []K {
	//keys
	result := maps.Keys(m)

	//order by
	if order == 1 {
		slices.SortStableFunc(result, func(a, b K) bool {
			return a < b
		})
	} else if order == 2 {
		slices.SortStableFunc(result, func(a, b K) bool {
			return a > b
		})
	}

	return result
}
