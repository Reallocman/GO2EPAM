package maps

import (
	"fmt"
	"sort"
)

func PrintSorted(m map[int]string) {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _,v:=range keys{
		fmt.Println(m[v])
	}
}
