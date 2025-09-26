package textarea

import (
	"testing"
)

func TestDynamicHeight(t *testing.T) {
	ta := New()
	ta.SetWidth(20) // Small width to force wrapping
	ta.MaxHeight = 10
	ta.SetDynamicHeight(true)

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

	// Height should not exceed MaxHeight
	manyLines := ""
	for i := 0; i < 15; i++ {
		manyLines += "Line " + string(rune('A'+i)) + "\n"
	}
	ta.SetValue(manyLines)
	if ta.Height() > ta.MaxHeight {
		t.Errorf("Height %d should not exceed MaxHeight %d", ta.Height(), ta.MaxHeight)
	}

	// Disabling dynamic height should keep current height
	ta.SetDynamicHeight(false)
	initialHeight := ta.Height()
	ta.SetValue("Short")
	if ta.Height() != initialHeight {
		t.Errorf("With dynamic height disabled, height should remain %d, got %d", initialHeight, ta.Height())
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
