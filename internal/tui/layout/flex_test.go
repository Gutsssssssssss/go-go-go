package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetExpandedTotal(t *testing.T) {
	elements := []element{
		Expanded("a"),
		Fixed("ce"),
		Expanded("d"),
	}
	// Test Vertical
	total, count, err := getExpandedTotal(Vertical, 9, elements)
	require.NoError(t, err)
	require.Equal(t, 8, total) // should be 9 - 1
	require.Equal(t, 2, count)

	// Test Horizontal
	total, count, err = getExpandedTotal(Horizontal, 9, elements)
	require.NoError(t, err)
	require.Equal(t, 7, total) // should be 9 - 2
	require.Equal(t, 2, count)
}

func TestGetFlexContents(t *testing.T) {
	elements := []element{
		Fixed("a"),
		Expanded(""),
		Fixed("b"),
		Expanded(""),
		Fixed("c"),
	}

	// Test Vertical
	t.Run("Vertical", func(t *testing.T) {
		contents := getFlexContents(Vertical, 10, elements)
		require.Equal(t, []string{
			"a", "\n\n\n", "b", "\n\n", "c",
		}, contents)
	})

	// Test Horizontal
	t.Run("Horizontal", func(t *testing.T) {
		contents := getFlexContents(Horizontal, 10, elements)
		require.Equal(t, []string{
			"a", "    ", "b", "   ", "c",
		}, contents)

		contents = getFlexContents(Horizontal, 4, elements)
		require.Equal(t, []string{
			"a", " ", "b", "", "c",
		}, contents)
	})
}

func TestFlex(t *testing.T) {
	elements := []element{
		Fixed("word1"),
		Expanded(""),
		Fixed("word2"),
	}
	layout := FlexVertical(5, elements...)
	require.Equal(t,
		"word1\n     \n     \n     \nword2",
		layout)
}
