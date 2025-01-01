# SSH-Based Portfolio

This project is an SSH-based portfolio built using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss). It emulates a terminal-like interface for showcasing your portfolio, with a polished UI and smooth transitions.
THIS IS A WORK IN PROGESS WITH MUTIPLE FEATURES YET TO BE IMPLEMENTED
## Features

- **Interactive Tabs**: Navigate through sections like Portfolio, Blog, and Contact Me.
- **Loading Screen**: Displays a loading screen with centered text for two seconds on startup.
- **Customizable Themes**: Stylish tab designs and color schemes.
- **Clear Exit**: Cleanly exits and clears the terminal when 'q' is pressed.

## Requirements

- Go 1.19 or later

## Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:daschinmoy21/ssh-portfolio.git
   ```

2. Navigate to the project directory:
   ```bash
   cd ssh-portfolio
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Run the portfolio:
   ```bash
   go run main.go
   ```

2. Navigate between tabs using arrow keys or specific key bindings .

3. Press 'q' to exit and clear the terminal.

## Configuration

To customize colors, themes, or text, modify the relevant sections in the `main.go` file or any associated configuration files.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Charm Bracelet](https://charm.sh) for their amazing libraries: Bubble Tea and Lip Gloss.

