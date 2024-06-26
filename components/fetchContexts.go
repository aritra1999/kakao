package components

import (
	"fmt"
	"kakau/k8s"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"k8s.io/client-go/util/homedir"
)

type model struct {
	contexts []item
	list, item    int
}

type item struct {
	text    string
	checked bool
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch typed := msg.(type) {
	case tea.KeyMsg:
		return m, m.handleKeyMsg(typed)
	}
	return m, nil
}

func (m *model) handleKeyMsg(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	case "esc", "ctrl+c":
		return tea.Quit
	case " ":
		switch m.list {
		case 0:
			m.contexts[m.item].checked = !m.contexts[m.item].checked
		}
	case "enter":
		return tea.Quit
	case "up":
		if m.item > 0 {
			m.item--
		} else if m.list > 0 {
			m.list--
			m.item = len(m.contexts) - 1
		}
	case "down":
		switch m.list {
		case 0:
			if m.item+1 < len(m.contexts) {
				m.item++
			} else {
				m.list++
				m.item = 0
			}
		}
	}
	return nil
}

func (m *model) View() string {
	curContext:= -1
	switch m.list {
		case 0:
			curContext = m.item
	}
	return m.renderList("choose contexts", m.contexts, curContext)
}

func (m *model) renderList(header string, items []item, selected int) string {
	out := "~ " + header + ":\n"
	for i, item := range items {
		sel := " "
		if i == selected {
			sel = ">"
		}
		check := " "
		if items[i].checked {
			check = "✓"
		}
		out += fmt.Sprintf("%s [%s] %s\n", sel, check, item.text)
	}
	return out
}

func loadContexts() []string {
	fmt.Println("loading contexts...")
	home := homedir.HomeDir();
	filePath := filepath.Join(home, ".kube", "config")
	config := k8s.GetConfig(filePath)
	contexts := k8s.FetchAllContexts(config)

	return contexts;
}

func LoadSelectedContexts() []string {
	contexts := loadContexts()
	return ContextMultiSelect(contexts)
}
