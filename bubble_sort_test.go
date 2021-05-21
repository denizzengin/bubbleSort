package bubble_sort

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_BubbleSort(t *testing.T) {

	arr := []int{1, 2, 3, 9, 6, 7, 11, 5}
	Sort(arr)
	expectedArr := []int{1, 2, 3, 5, 6, 7, 9, 11}
	fmt.Println(arr)
	fmt.Println(expectedArr)
	assert.Equal(t, arr, expectedArr)
}
