package slices

func Max(slice []string) string {
	max:=slice[0]
	for i:=0;i<len(slice);i++{
		if len([]rune (max))<len([]rune (slice[i])){
		  	max=slice[i]}
	}
	return max
}

func Reverse(slice []int64) []int64{
	copy:=make([]int64,0)
	for i:=len(slice)-1;i>=0;i--{
		copy=append(copy,slice[i])
	}
	return copy
}
