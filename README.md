# Bubble Textarea

A dynamic height textarea component for [Bubble Tea](https://github.com/charmbracelet/bubbletea) applications.

This component extends the wonderful [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles) `textarea` with support for dynamic height. The textarea can be configured to automatically grow and shrink in height based on its content, respecting a configurable maximum visual height.

<img src="./example/example.gif" width="600" alt="Bubble Textarea Demo">

## Usage

Dynamic height is enabled by calling `SetMaxVisualHeight` with a positive integer:

```go
ta := textarea.New()
// Allow the rendered height to grow up to 10 lines of content, including both hard & soft line breaks.
ta.SetMaxVisualHeight(10)
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
    ta.SetHeight(1)

    // Enables dynamic height; Grow up to 10 lines before viewport scrolling kicks in
    ta.SetMaxVisualHeight(10)

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

MIT License - see [LICENSE](LICENSE) for details.
