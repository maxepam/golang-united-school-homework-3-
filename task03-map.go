package homework

import "sort"

type SliceInt []int

func (v SliceInt) Len() int {
	return len(v)
}

func (v SliceInt) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v SliceInt) Less(i, j int) bool {
	return v[i] < v[j]
}

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Sort(SliceInt(keys))

	output := make([]string, 0, len(input))
	for _, k := range keys {
		output = append(output, input[k])
	}

	return output
}