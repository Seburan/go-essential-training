// Challenge : Find the maximal value
package main

import (
	"fmt"
)

func main() {
	nums := []int{16,8,42,4,23,15};

	// initialize max to first element
	var max = nums[0];
	// skip the first value
	for _ , v := range nums[1:] {
		if max < v {
			max = v;
		}
	}

	fmt.Println(max);
}
