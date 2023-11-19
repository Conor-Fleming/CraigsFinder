package main

import (
	"fmt"
	"os"

	"github.com/Conor-Flemign/CraigsFinder/cli"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tabs := []string{"New Search", "Saved Search", "Results", "Saved Results"}
	tabContent := []string{"Search Tab", "Saved Searches", "list of results", "Saved Results"}
	m := cli.Model{Tabs: tabs, TabContent: tabContent}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
