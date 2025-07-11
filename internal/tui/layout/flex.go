package layout

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type flexDirection int

var (
	ErrNotEnoughSpace = fmt.Errorf("not enough space")
)

const (
	Vertical flexDirection = iota
	Horizontal
)

type element struct {
	content  string
	expanded bool
}

func Expanded(content string) element {
	return element{
		content:  content,
		expanded: true,
	}
}

func Fixed(content string) element {
	return element{
		content:  content,
		expanded: false,
	}
}

func FlexVertical(height int, elements ...element) string {
	contents := getFlexContents(Vertical, height, elements)
	return lipgloss.JoinVertical(lipgloss.Center, contents...)
}

func FlexHorizontal(width int, elements ...element) string {
	contents := getFlexContents(Horizontal, width, elements)
	return lipgloss.JoinVertical(lipgloss.Center, contents...)
}

// getFlexContents return the contents of the flex layout
// it uses evenly splited expanded elements to fill the space
func getFlexContents(direction flexDirection, maxSize int, elements []element) []string {
	var contents []string
	exTotal, exCount, err := getExpandedTotal(direction, maxSize, elements)
	if err != nil {
		return getPlainContents(elements)
	}
	// fmt.Println(exTotal, exCount)
	exSizes, err := getEvenlySplitedSlice(exTotal, exCount)
	if err != nil {
		return getPlainContents(elements)
	}
	exIdx := 0
	switch direction {
	case Vertical:
		exStyle := lipgloss.NewStyle().AlignVertical(lipgloss.Center)
		for _, e := range elements {
			if !e.expanded {
				contents = append(contents, e.content)
			} else {
				contents = append(contents,
					exStyle.Height(exSizes[exIdx]).Render(e.content))
				exIdx++
			}
		}
	case Horizontal:
		exStyle := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
		for _, e := range elements {
			if !e.expanded {
				contents = append(contents, e.content)
			} else {
				contents = append(contents,
					exStyle.Width(exSizes[exIdx]).Render(e.content))
				exIdx++
			}
		}
	}
	return contents
}

// getExpandedTotal return the total size of the expanded elements and the count of expanded elements
func getExpandedTotal(direction flexDirection, maxSize int, elements []element) (int, int, error) {
	totalFixedSize, expandedCount := 0, 0
	for _, e := range elements {
		if !e.expanded {
			if direction == Vertical {
				totalFixedSize += lipgloss.Height(e.content)
			} else {
				totalFixedSize += lipgloss.Width(e.content)
			}
		} else {
			expandedCount++
		}
	}
	exTotal := maxSize - totalFixedSize
	if exTotal < 0 {
		return 0, 0, ErrNotEnoughSpace
	}
	return exTotal, expandedCount, nil
}

func getPlainContents(elements []element) []string {
	var contents []string
	for _, e := range elements {
		contents = append(contents, e.content)
	}
	return contents
}
