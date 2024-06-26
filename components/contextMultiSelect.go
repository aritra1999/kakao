package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func ContextMultiSelect(contexts []string) []string {
	var contextOptions []item
	for _, context := range contexts {
		contextOptions = append(contextOptions, item{text: context, checked: false})
	}

	m := &model{contexts: contextOptions}
	
	if err := tea.NewProgram(m).Start(); err != nil {
		panic(fmt.Sprintf("failed to run program: %v", err))
	}

	var selectedContexts []string
	for _, context := range m.contexts {
		if context.checked {
			selectedContexts = append(selectedContexts, context.text)
		}
	}

	return selectedContexts
}

