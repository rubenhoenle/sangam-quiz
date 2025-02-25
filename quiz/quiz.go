package quiz

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	daomodel "github.com/rubenhoenle/sangam-quiz/model"
	jsonprovider "github.com/rubenhoenle/sangam-quiz/provider/json-provider"
	"github.com/rubenhoenle/sangam-quiz/service"
	"github.com/rubenhoenle/sangam-quiz/util"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	list         list.Model
	searchedItem daomodel.SangamItem
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("H", "toggle help"),
		)):
			if m.list.SelectedItem().FilterValue() == m.searchedItem.Name {
				return m, tea.Quit
			}
			break
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func InitializeModel() model {
	var sangamItemProvider service.SangamItemProvider = jsonprovider.JsonSangamItemProvider{}
	var items []list.Item
	sangamItems, err := sangamItemProvider.GetSangamItems()
	if err != nil {
		// TODO: error handling
	}
	for _, si := range sangamItems {
		items = append(items, si)
	}

	quizItem := util.GetRandomSangamItem(sangamItems)

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0), searchedItem: quizItem}
	m.list.Title = fmt.Sprintf("Sangam Quiz - Please select the menu with this id: %s", quizItem.Id)

	return m
}
