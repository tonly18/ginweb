package command

import (
	"cmp"
	"sort"
)

// MapKeys 获取map所有key
//
// @params:
//
//	m		map[K]V	map
//	order	int8	0不排序|1升序|2降序
//
// @return:
func MapKeys[K cmp.Ordered, V any](m map[K]V, order int8) []K {
	//keys
	keys := make([]K, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}

	//order by
	if order == 1 {
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
	} else if order == 2 {
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i] > keys[j]
		})
	}

	return keys
}
