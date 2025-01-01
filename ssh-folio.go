package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

// Constants for sections
const (
	Portfolio = iota
	Blog
	Contact
	LoadingScreen
)

type model struct {
	currentView  int
	spinner      spinner.Model
	width        int
	height       int
	loadingDone  bool
}

var (
	headerStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("240")).
		Foreground(lipgloss.Color("229")).
		Padding(0, 1).
		Bold(true)

	activeTabStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("229")).
		Foreground(lipgloss.Color("240")).
		Padding(0, 1).
		Bold(true)

	inactiveTabStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Background(lipgloss.Color("0")).
		Padding(0, 1)

	footerStyle = lipgloss.NewStyle().
		MarginTop(1).
		Foreground(lipgloss.Color("E69821"))

	centeredStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("E69821")).
		Align(lipgloss.Center).
		Width(60).
		Height(30)

	terminalStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Bold(true).
		Align(lipgloss.Center)

	cursorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("166")).
		Bold(true)
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			if m.currentView == LoadingScreen {
				time.Sleep(2 * time.Second)
				return "loading-done"
			}
			return nil
		},
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			clearScreen()
			return m, tea.Quit
		case "p":
			m.currentView = Portfolio
		case "b":
			m.currentView = Blog
		case "c":
			m.currentView = Contact
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case string:
		if msg == "loading-done" {
			m.loadingDone = true
			m.currentView = Portfolio
		}
	}
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.currentView == LoadingScreen && !m.loadingDone {
		cursor := ""
		if time.Now().UnixNano()/int64(time.Millisecond)%1000 < 500 {
			cursor = cursorStyle.Render("█")
		}

		loadingText := terminalStyle.Render("chinmoy.dev" + cursor)
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			loadingText,
		)
	}

	tabs := fmt.Sprintf(
		"%s | %s | %s",
		m.tab("Portfolio", Portfolio),
		m.tab("Blog", Blog),
		m.tab("Contact Me", Contact),
	)

	var content string
	switch m.currentView {
	case Portfolio:
		content = "🛠️ Hey! I'm Chinmoy and Welcome to my Portfolio!\n\n" +
			"Here you can see details about my projects.\n\n" +
			m.spinner.View() + " Loading portfolio..."
	case Blog:
		content = "📖 Welcome to my Blog!\n\n" +
			"Fetching latest posts...\n\n" +
			m.spinner.View()
	case Contact:
		content = "📬 Contact Me\n\nEmail: daschinmoyy21@gmail.com\nGitHub: github.com/daschinmoy21"
	}

	footer := footerStyle.Render("Press 'p', 'b', or 'c' to navigate | 'q' to quit.")
	body := fmt.Sprintf("%s\n\n%s\n\n%s", headerStyle.Render(tabs), content, footer)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.NewStyle().Padding(1, 2).Render(body),
	)
}

func (m model) tab(name string, section int) string {
	if m.currentView == section {
		return activeTabStyle.Render(name)
	}
	return inactiveTabStyle.Render(name)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func main() {
	initialModel := model{
		currentView: LoadingScreen,
		spinner:     spinner.New(),
	}
	initialModel.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := tea.NewProgram(
		initialModel,
		tea.WithAltScreen(),
	)

	if err := p.Start(); err != nil {
		fmt.Printf("Error starting app: %v\n", err)
		os.Exit(1)
	}
}
