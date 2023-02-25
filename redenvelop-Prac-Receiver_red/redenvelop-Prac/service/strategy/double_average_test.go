package strategy

import (
	"fmt"
	"testing"
)

func TestDoubleAverage(t *testing.T) {
	count := int64(5)
	amount := int64(1000)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		remain -= x
		sum += x
		fmt.Println(i+1, "=", float64(x)/float64(100), ",")

	}
	fmt.Print("#{sum}\n")
}
