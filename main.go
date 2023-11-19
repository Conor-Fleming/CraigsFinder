package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := []list.Item{
		cli.item("Ramen"),
		item("Tomato Soup"),
		item("Hamburgers"),
		item("Cheeseburgers"),
		item("Currywurst"),
		item("Okonomiyaki"),
		item("Pasta"),
		item("Fillet Mignon"),
		item("Caviar"),
		item("Just Wine"),
	}

	const defaultWidth = 20

	l := list.New(items, cli.itemDelegate{}, defaultWidth, cli.listHeight)
	l.Title = "What do you want for dinner?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = cli.titleStyle
	l.Styles.PaginationStyle = cli.paginationStyle
	l.Styles.HelpStyle = cli.helpStyle

	m := cli.model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
