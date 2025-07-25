package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPercentage(t *testing.T) {
	require.Equal(t, 0.5, GetPercentage(0))
	require.Equal(t, 0.75, GetPercentage(90))
	require.Equal(t, 1.0, GetPercentage(180))
	require.Equal(t, 0.25, GetPercentage(270))
}
