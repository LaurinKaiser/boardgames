package tictactoe

// PlayerWins prüft, ob der angegebene Spieler das Spiel gewonnen hat.
// Der Spieler wird dabei durch das Zeichen repräsentiert,
// das er auf dem Spielfeld verwendet (z.B. "X" oder "O").
func (g *Game) PlayerWins(player string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion BoardContainsChain mit einer Länge von 3.
	return g.board.BoardContainsChain(player, 3)
}

// GameOver prüft, ob das Spiel vorbei ist, d.h. ob ein Spieler gewonnen hat oder das Spielfeld voll ist.
func (g *Game) GameOver() bool {
	// Hinweis:
	// Verwenden Sie PlayerWins für beide Spieler und BoardContains, um zu überprüfen, ob das Spielfeld voll ist.
	return g.PlayerWins("X") || g.PlayerWins("O") || !g.board.BoardContains(" ")
}

// IsDraw prüft, ob das Spiel unentschieden ist, d.h. ob das Spielfeld voll ist und kein Spieler gewonnen hat.
func (g *Game) IsDraw() bool {
	// Hinweis:
	// Verwenden Sie die bisherigen Logik-Funktionen von Game.
	return !g.board.BoardContains(" ") && !g.PlayerWins("X") && !g.PlayerWins("O")
}

// MoveAllowed prüft, ob ein Zug an der angegebenen Position erlaubt ist.
// Ein Zug ist erlaubt, wenn die Position innerhalb des Spielfelds liegt und die Zelle an dieser Position leer ist.
func (g *Game) MoveAllowed(row, col int) bool {
	// Hinweis:
	// Die Prüfungslogik für row und col ist bereits in Get implementiert.
	// Sie müssen hier also nur noch überprüfen, was Get geliefert hat.
	return g.board.Get(row, col) == " "
}
