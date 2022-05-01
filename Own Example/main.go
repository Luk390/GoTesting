package main

import "fmt"

type avgSlice struct {
	values []float64
}

func avg(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	avg := sum / float64(len(s))
	return avg
}

func main() {
	nums := avgSlice{
		values: []float64{1.2, 3.2, 1, 45.8},
	}
	nums2 := avgSlice{
		values: []float64{1.2, 1.2, 1.2},
	}
	nums3 := avgSlice{
		values: []float64{1,1,1},
	}
	fmt.Println(avg(nums.values))
	fmt.Println(avg(nums2.values))
	fmt.Println(avg(nums3.values))

}
