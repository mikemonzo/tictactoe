package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	Empty = iota
	PlayerX
	PlayerO
)

var (
	board         [3][3]int
	currentPlayer = PlayerX
	buttons       [3][3]*widget.Button
	statusLabel   *widget.Label
)

func main() {
	application := app.New()
	window := application.NewWindow("Tic Tac Toe")

	initializeBoard()
	statusLabel = widget.NewLabel("Player X's turn")

	grid := container.NewGridWithColumns(3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			button := widget.NewButton("", createCellHandler(i, j))
			buttons[i][j] = button
			grid.Add(button)
		}
	}
	window.SetContent(container.NewVBox(
		statusLabel,
		grid,
		widget.NewButton("Reset", resetGame),
	))

	window.Resize(fyne.NewSize(300, 300))
	window.ShowAndRun()
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = Empty
		}
	}
}

func createCellHandler(i, j int) func() {
	return func() {
		if board[i][j] == Empty {
			board[i][j] = currentPlayer
			updateButtonLabel(i, j)
			if checkWin() {
				statusLabel.SetText("Player " + getPlayerLabel(currentPlayer) + " wins!")
				disableBoard()
			} else {
				switchPlayer()
			}
		}
	}
}

func updateButtonLabel(i, j int) {
	switch board[i][j] {
	case PlayerX:
		buttons[i][j].SetText("X")
	case PlayerO:
		buttons[i][j].SetText("O")
	}
}

func switchPlayer() {
	if currentPlayer == PlayerX {
		currentPlayer = PlayerO
	} else {
		currentPlayer = PlayerX
	}
	statusLabel.SetText("Player " + getPlayerLabel(currentPlayer) + "'s turn")
}

func getPlayerLabel(player int) string {
	if player == PlayerX {
		return "X"
	}
	return "O"
}

func checkWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] != Empty && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return true
		}
		if board[0][i] != Empty && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return true
		}
	}
	if board[0][0] != Empty && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return true
	}
	if board[0][2] != Empty && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return true
	}
	return false
}

func disableBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			buttons[i][j].Disable()
		}
	}
}

func resetGame() {
	initializeBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			buttons[i][j].SetText(" ")
			buttons[i][j].Enable()
		}
	}
	statusLabel.SetText("Player X's turn")
	currentPlayer = PlayerX
}
