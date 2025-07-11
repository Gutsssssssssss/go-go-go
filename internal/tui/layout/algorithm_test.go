package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEvenlyDistributedBooleans(t *testing.T) {
	// Test 1: Total < trueCount
	_, err := getEvenlyDistributedBooleans(4, 5)
	require.EqualError(t, err, ErrInvalidTotalBoolean.Error())

	// Test 2: trueCount < 0
	_, err = getEvenlyDistributedBooleans(5, -1)
	require.EqualError(t, err, ErrInvalidTrueCount.Error())

	// Test 3: 0 trueCount
	slice, err := getEvenlyDistributedBooleans(5, 0)
	require.NoError(t, err)
	require.Equal(t, []bool{false, false, false, false, false}, slice)

	// Test: total is Odd
	slice, err = getEvenlyDistributedBooleans(5, 1)
	require.NoError(t, err)
	require.Equal(t, []bool{false, false, true, false, false}, slice)

	slice, err = getEvenlyDistributedBooleans(5, 2)
	require.NoError(t, err)
	require.Equal(t, []bool{true, false, false, false, true}, slice)
	// Test 4: get evenly distributed booleans
	slice, err = getEvenlyDistributedBooleans(5, 3)
	require.NoError(t, err)
	require.Equal(t, []bool{true, false, true, false, true}, slice)

	slice, err = getEvenlyDistributedBooleans(5, 4)
	require.NoError(t, err)
	require.Equal(t, []bool{true, true, false, true, true}, slice)

	slice, err = getEvenlyDistributedBooleans(5, 5)
	require.NoError(t, err)
	require.Equal(t, []bool{true, true, true, true, true}, slice)

	// Test: total is Even
	slice, err = getEvenlyDistributedBooleans(6, 1)
	require.NoError(t, err)
	require.Equal(t, []bool{false, false, true, false, false, false}, slice)

	slice, err = getEvenlyDistributedBooleans(6, 2)
	require.NoError(t, err)
	require.Equal(t, []bool{true, false, false, false, false, true}, slice)

	// Test 5: get evenly distributed booleans
	slice, err = getEvenlyDistributedBooleans(6, 3)
	require.NoError(t, err)
	require.Equal(t, []bool{true, false, true, false, false, true}, slice)

	slice, err = getEvenlyDistributedBooleans(6, 4)
	require.NoError(t, err)
	require.Equal(t, []bool{true, true, false, false, true, true}, slice)

	// Test 6: get evenly distributed booleans
	slice, err = getEvenlyDistributedBooleans(6, 5)
	require.NoError(t, err)
	require.Equal(t, []bool{true, true, true, false, true, true}, slice)

	slice, err = getEvenlyDistributedBooleans(6, 6)
	require.NoError(t, err)
	require.Equal(t, []bool{true, true, true, true, true, true}, slice)

	// Test 7: corner cases
	slices, err := getEvenlyDistributedBooleans(0, 0)
	require.NoError(t, err)
	require.Equal(t, []bool{}, slices)
}

func TestGetEvenlySplitedSlice(t *testing.T) {
	// Test 1: 0 count
	slice, err := getEvenlySplitedSlice(10, 0)
	require.Equal(t, []int{}, slice)
	require.NoError(t, err)

	// Test 2: Invalid sum and count
	_, err = getEvenlySplitedSlice(10, -1)
	require.EqualError(t, err, ErrInvalidCount.Error())

	// Test 2: Valid inputs (sum: 10)
	slice, err = getEvenlySplitedSlice(10, 1)
	require.NoError(t, err)
	require.Equal(t, []int{10}, slice)

	slice, err = getEvenlySplitedSlice(10, 2)
	require.NoError(t, err)
	require.Equal(t, []int{5, 5}, slice)

	slice, err = getEvenlySplitedSlice(10, 3)
	require.NoError(t, err)
	require.Equal(t, []int{3, 4, 3}, slice)

	slice, err = getEvenlySplitedSlice(10, 4)
	require.NoError(t, err)
	require.Equal(t, []int{3, 2, 2, 3}, slice)

	slice, err = getEvenlySplitedSlice(10, 5)
	require.NoError(t, err)
	require.Equal(t, []int{2, 2, 2, 2, 2}, slice)

	slice, err = getEvenlySplitedSlice(10, 6)
	require.NoError(t, err)
	require.Equal(t, []int{2, 2, 1, 1, 2, 2}, slice)

	slice, err = getEvenlySplitedSlice(10, 7)
	require.NoError(t, err)
	require.Equal(t, []int{2, 1, 1, 2, 1, 1, 2}, slice)

	slice, err = getEvenlySplitedSlice(10, 8)
	require.NoError(t, err)
	require.Equal(t, []int{2, 1, 1, 1, 1, 1, 1, 2}, slice)

	slice, err = getEvenlySplitedSlice(10, 9)
	require.NoError(t, err)
	require.Equal(t, []int{1, 1, 1, 1, 2, 1, 1, 1, 1}, slice)

	slice, err = getEvenlySplitedSlice(10, 10)
	require.NoError(t, err)
	require.Equal(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, slice)
}
