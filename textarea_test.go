package textarea

import (
	"strings"
	"testing"
)

func TestDynamicHeight(t *testing.T) {
	ta := New()
	ta.SetWidth(20)           // Small width to force wrapping
	ta.MaxHeight = 100        // Allow many content lines
	ta.SetMaxVisualHeight(10) // But limit visual height to 10

	// Empty textarea should have height 1
	if ta.Height() != 1 {
		t.Errorf("Empty textarea height should be 1, got %d", ta.Height())
	}

	// Single line should have height 1
	ta.SetValue("Hello")
	expectedHeight := 1
	if ta.Height() != expectedHeight {
		t.Errorf("Single line height should be %d, got %d", expectedHeight, ta.Height())
	}

	// Two lines should have height 2
	ta.SetValue("Hello\nWorld")
	expectedHeight = 2
	if ta.Height() != expectedHeight {
		t.Errorf("Two line height should be %d, got %d", expectedHeight, ta.Height())
	}

	// Long line that wraps should increase height
	longLine := "This is a very long line that should wrap around multiple times when displayed in the textarea with a narrow width"
	ta.SetValue(longLine)
	// With width 20, this should wrap to multiple lines
	if ta.Height() <= 1 {
		t.Errorf("Long wrapping line should have height > 1, got %d", ta.Height())
	}

	// Height should not exceed MaxVisualHeight
	manyLines := ""
	for i := 0; i < 15; i++ {
		manyLines += "Line " + string(rune('A'+i)) + "\n"
	}
	ta.SetValue(manyLines)
	if ta.Height() > ta.MaxVisualHeight {
		t.Errorf("Height %d should not exceed MaxVisualHeight %d", ta.Height(), ta.MaxVisualHeight)
	}

	// Disabling dynamic height should keep current height
	ta.SetMaxVisualHeight(0)
	initialHeight := ta.Height()
	ta.SetValue("Short")
	if ta.Height() != initialHeight {
		t.Errorf("With dynamic height disabled, height should remain %d, got %d", initialHeight, ta.Height())
	}
}

func TestMaxHeightGreaterThanMaxVisualHeight(t *testing.T) {
	// Test scenario: MaxHeight > MaxVisualHeight
	// Should allow more content lines than visual height
	ta := New()
	ta.SetWidth(20)
	ta.MaxHeight = 100       // Allow up to 100 content lines
	ta.SetMaxVisualHeight(3) // But only show 3 lines visually

	// Should be able to add newlines beyond the visual height limit
	lines := []string{}
	for i := 0; i < 10; i++ {
		lines = append(lines, "Line "+string(rune('A'+i)))
	}
	content := strings.Join(lines, "\n")
	ta.SetValue(content)

	// Visual height should be clamped to MaxVisualHeight
	if ta.Height() != 3 {
		t.Errorf("Visual height should be 3 (MaxVisualHeight), got %d", ta.Height())
	}

	// Content should have all 10 lines
	if ta.LineCount() != 10 {
		t.Errorf("Content should have 10 lines, got %d", ta.LineCount())
	}

	// Should be able to add more lines up to MaxHeight
	moreLines := []string{}
	for i := 0; i < 20; i++ {
		moreLines = append(moreLines, "Line "+string(rune('A'+i)))
	}
	moreContent := strings.Join(moreLines, "\n")
	ta.SetValue(moreContent)

	// Visual height should still be clamped
	if ta.Height() != 3 {
		t.Errorf("Visual height should still be 3, got %d", ta.Height())
	}

	// Content should have all 20 lines
	if ta.LineCount() != 20 {
		t.Errorf("Content should have 20 lines, got %d", ta.LineCount())
	}
}

func TestMaxHeightLessThanMaxVisualHeight(t *testing.T) {
	// Test scenario: MaxHeight < MaxVisualHeight
	// Content constraint should kick in before visual constraint
	ta := New()
	ta.SetWidth(20)
	ta.MaxHeight = 3          // Only allow 3 content lines
	ta.SetMaxVisualHeight(10) // But visual could grow to 10

	// Should not be able to add more than 3 lines
	lines := []string{}
	for i := 0; i < 10; i++ {
		lines = append(lines, "Line "+string(rune('A'+i)))
	}
	content := strings.Join(lines, "\n")
	ta.SetValue(content)

	// Content should be limited to MaxHeight
	if ta.LineCount() > 3 {
		t.Errorf("Content should be limited to 3 lines (MaxHeight), got %d", ta.LineCount())
	}

	// Visual height should match content (since content is limited)
	if ta.Height() > 3 {
		t.Errorf("Visual height should not exceed 3, got %d", ta.Height())
	}
}

func TestMaxVisualHeightZeroDisabled(t *testing.T) {
	// Test scenario: MaxVisualHeight = 0 (disabled)
	// Should use fixed height, no dynamic resizing
	ta := New()
	ta.SetWidth(20)
	ta.SetHeight(5)          // Set fixed height
	ta.MaxHeight = 100       // Allow many content lines
	ta.SetMaxVisualHeight(0) // Disable dynamic height

	initialHeight := ta.Height()

	// Adding content should not change height
	ta.SetValue("Line 1\nLine 2\nLine 3\nLine 4\nLine 5\nLine 6\nLine 7")

	if ta.Height() != initialHeight {
		t.Errorf("Height should remain fixed at %d when MaxVisualHeight=0, got %d", initialHeight, ta.Height())
	}

	// Should still be able to add content lines
	if ta.LineCount() != 7 {
		t.Errorf("Should have 7 content lines, got %d", ta.LineCount())
	}
}

func TestContentHeightCalculation(t *testing.T) {
	ta := New()
	ta.SetWidth(10) // Very narrow for testing wrapping

	height := ta.CalculateContentHeight()
	if height != 1 {
		t.Errorf("Empty content height should be 1, got %d", height)
	}

	ta.SetValue("Hi")
	height = ta.CalculateContentHeight()
	if height != 1 {
		t.Errorf("Single short line height should be 1, got %d", height)
	}

	ta.SetValue("This is a long line that will wrap")
	height = ta.CalculateContentHeight()
	if height <= 1 {
		t.Errorf("Wrapping line height should be > 1, got %d", height)
	}

	ta.SetValue("Line 1\nLine 2\nLine 3")
	height = ta.CalculateContentHeight()
	if height < 3 {
		t.Errorf("Three lines should have height >= 3, got %d", height)
	}
}
