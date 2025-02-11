package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rubenhoenle/sangam-quiz/quiz"
	"os"
)

func main() {
	m := quiz.InitializeModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
