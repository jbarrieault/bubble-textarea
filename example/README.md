# Dynamic Textarea Example

This example demonstrates the dynamic height functionality of the bubble-textarea component.

## Running the Example

```bash
go run main.go
```

## What to Try

1. **Type text**: Watch the textarea grow as you type
2. **Add newlines**: Press Enter to see hard line breaks expand the height
3. **Long lines**: Type very long lines to see soft wrapping in action
4. **Reach the limit**: Keep typing until you hit the MaxHeight (5 lines) and see scrolling take over
5. **Delete content**: Remove text and watch the textarea shrink back down

## Key Features Demonstrated

- Dynamic height starting from 1 line
- Maximum height of 5 lines with viewport scrolling
- Debug information showing:
  - Current visual height
  - Maximum height setting
  - Hard line count (actual newlines)
  - Content length
  - Window width
- Custom prompt styling
- Border styling

## Controls

- **Ctrl+C** or **Esc**: Quit the application
- **Enter**: Insert newline
- **Backspace/Delete**: Remove characters
- Standard text navigation keys work as expected