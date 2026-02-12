package tictactoe

import "boardgames/board"

type Game struct {
	board *board.Board
}

// NewGame erstellt ein neues Spiel mit einem leeren 3x3-Spielfeld.
func NewGame() *Game {
	return &Game{
		board: board.New(3, 3),
	}
}

// Get liefert den Inhalt der Zelle an der angegebenen Position zurück.
// Liefert einen leeren String, falls die Position außerhalb des Spielfelds liegt.
func (g *Game) Get(row, col int) string {
	// Hinweis:
	// Sie können direkt board.Get verwenden, da die Überprüfung der Grenzen bereits dort implementiert ist.
	return g.board.Get(row, col)
}

// Set setzt den Inhalt der Zelle an der angegebenen Position auf den angegebenen Wert.
// Ignoriert die Anweisung, falls der Zug nicht erlaubt ist.
func (g *Game) Set(row, col int, value string) {
	// Hinweis:
	// Sie können direkt board.Set verwenden, aber überprüfen Sie vorher mit MoveAllowed, ob der Zug erlaubt ist.
	if g.MoveAllowed(row, col) {
		g.board.Set(row, col, value)
	}
}
