package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/etherealblaade/rattus_rex/internal/chain"
	"strings"
)

type model struct {
	chain    *chain.ModelChain
	input    string
	messages []string
	err      error
}

func NewModel(c *chain.ModelChain) model {
	return model{
		chain:    c,
		messages: make([]string, 0),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) handleCommand(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	switch parts[0] {
	case "/clear":
		m.messages = make([]string, 0)
		return "History cleared", nil
	case "/model":
		if len(parts) < 2 {
			return "", fmt.Errorf("model name required")
		}
		m.chain.OpenRouterModel = parts[1]
		return fmt.Sprintf("Model changed to %s", parts[1]), nil
	case "/reasoning":
		m.chain.ShowReasoning = !m.chain.ShowReasoning
		status := "enabled"
		if !m.chain.ShowReasoning {
			status = "disabled"
		}
		return fmt.Sprintf("Reasoning %s", status), nil
	default:
		return "", fmt.Errorf("unknown command")
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.input == "" {
				return m, nil
			}
			if strings.HasPrefix(m.input, "/") {
				resp, err := m.handleCommand(m.input)
				if err != nil {
					m.err = err
				} else {
					m.messages = append(m.messages, resp)
				}
				m.input = ""
				return m, nil
			}
			resp, err := m.chain.Process(m.input)
			if err != nil {
				m.err = err
				return m, nil
			}
			m.messages = append(m.messages, m.input, resp)
			m.input = ""
			return m, nil
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil
		default:
			if len(msg.String()) == 1 {
				m.input += msg.String()
			}
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "🐀 Rattus Rex\n\n"

	for i := 0; i < len(m.messages); i += 2 {
		s += fmt.Sprintf("You: %s\n", m.messages[i])
		if i+1 < len(m.messages) {
			s += fmt.Sprintf("AI: %s\n", m.messages[i+1])
		}
		s += "\n"
	}

	if m.err != nil {
		s += fmt.Sprintf("\nError: %v\n", m.err)
	}

	s += fmt.Sprintf("\n> %s", m.input)
	return s
}
