package layout

import (
	"fmt"
)

var (
	ErrInvalidTrueCount    = fmt.Errorf("invalid true count")
	ErrInvalidTotalBoolean = fmt.Errorf("invalid total booleans")

	ErrInvalidCount = fmt.Errorf("invalid count")
)

// getEvenlyDistributedBooleans return a slice of evenly distributed booleans.
//
// It always try to distribute symmetrically, and true values are always trying to
// be placed to the each sides (left and right).
// If total is even and trueCount is odd,
// the last value will be on the right side of the middle.
//
// Example:
// total = 5, trueCount = 3
// return: [true, false, true, false, true]
func getEvenlyDistributedBooleans(total, trueCount int) ([]bool, error) {
	if total < trueCount {
		return nil, ErrInvalidTotalBoolean
	}
	if trueCount < 0 {
		return nil, ErrInvalidTrueCount
	}
	result := make([]bool, total)
	left := 0
	right := total - 1
	for trueCount > 0 && left <= right {
		if left == right {
			result[left] = true
			trueCount--
		} else {
			if trueCount >= 2 {
				result[left] = true
				result[right] = true
				trueCount -= 2
			}
		}
		left++
		right--
	}
	// total is even and trueCount is odd
	if trueCount > 0 {
		result[right] = true
	}
	return result, nil
}

// getEvenlySplitedSlice return a slice of evenly splited elements
// with sum of elements equals to total, and len(slice) == count.
//
// It always try to split symmetrically.
//
// Example:
// total = 10, count = 3
// return [3, 4, 3]
func getEvenlySplitedSlice(sum, count int) ([]int, error) {
	if count < 0 {
		return []int{}, ErrInvalidCount
	}
	if count == 0 {
		return []int{}, nil
	}
	base := sum / count
	higher := base + 1
	higherCount := sum % count // we need to add higher value as this amoumt
	bools, err := getEvenlyDistributedBooleans(count, higherCount)
	if err != nil {
		return nil, err
	}
	slice := make([]int, count)
	for i, high := range bools {
		if high {
			slice[i] = higher
		} else {
			slice[i] = base
		}
	}
	return slice, nil
}
