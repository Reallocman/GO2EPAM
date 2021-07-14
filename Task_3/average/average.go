package average

func Mean(array [6]float64 )  float64 {
	var sum float64
	for i:=0;i<len(array);i++{
		sum+=array[i]
	}
	return sum/float64(len(array))
}
