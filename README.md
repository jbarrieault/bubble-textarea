# Bubble Textarea

A dynamic height textarea component for [Bubble Tea](https://github.com/charmbracelet/bubbletea) applications.

This component extends the functionality of [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles) textarea with support for dynamic height. The textarea automatically grows and shrinks in height based on its content, respecting a configurable maximum height.

## Usage

### Key Methods for Dynamic Height

```go
ta := textarea.New()

// Enable or disable dynamic height functionality
ta.SetDynamicHeight(true)

// Set initial height (will grow from this size)
ta.SetHeight(1)

// Set maximum height (component will scroll when exceeded)
ta.MaxHeight = 10
```

### Basic Example

```go
package main

import (
    "github.com/jbarrieault/bubble-textarea"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    textarea textarea.Model
}

func initialModel() model {
    ta := textarea.New()
    ta.Placeholder = "Start typing..."
    ta.Focus()

    // Enable dynamic height
    ta.SetDynamicHeight(true)
    ta.SetHeight(1)      // Start with 1 line
    ta.MaxHeight = 10    // Grow up to 10 lines before viewport scrolling kicks in

    return model{textarea: ta}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.textarea, cmd = m.textarea.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return m.textarea.View()
}
```



## Example

See the [example](./example/) directory for a complete working application that demonstrates the dynamic height functionality.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Credits

This project is just a small extension of the wonderful [charmbracelet/bubbles/textarea](https://github.com/charmbracelet/bubbles/tree/main/textarea).
