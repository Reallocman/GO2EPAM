package fibonacci
import "fmt"

func Loop() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}
func PrintLoop(n int) {
	f:=Loop()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ",f())
	}
}
