package tictactoe

import (
	"boardgames/board"
	"fmt"
)

// Show zeigt den aktuellen Spielstand an.
func (g *Game) Show() {
	fmt.Println(g.board)
}

// Move fragt den Spieler nach einem Zug und führt diesen Zug aus, falls er erlaubt ist.
// Falls der Zug nicht erlaubt ist, wird der Spieler erneut nach einem Zug gefragt.
func (g *Game) Move(player string) {
	// Hinweis:
	// - Zeigen Sie zuerst das Spielfeld mit den nummerierten Zellen an, damit der Spieler weiß, welche Zahl welcher Zelle entspricht.
	// - Fragen Sie den Spieler dann nach seiner Eingabe.
	// - Für die Überprüfung der Eingabe können Sie sie als String einlesen und dann überprüfen, ob dieser
	//   String eine gültige Zahl zwischen "1" und "9" ist.
	//   Falls nicht, geben Sie eine Fehlermeldung aus und rufen Sie Move erneut auf.
	// - Wandeln Sie dann die eingegebene Zahl in die entsprechenden Zeilen- und Spaltenindizes um.
	// - Die Zeilen- und Spaltenindizes können Sie mit einfachen mathematischen Operationen berechnen.
	// - Überprüfen Sie auch diese Indizes erneut auf Gültigkeit, bevor Sie den Zug ausführen.
	//   Falls der Zug nicht erlaubt ist, geben Sie eine Fehlermeldung aus und rufen Sie Move erneut auf.
	fmt.Println(board.NewWithNumbers(3, 3))
	fmt.Printf("Spieler %s, bitte gib deinen Zug ein: ", player)

	var input string
	fmt.Scanln(&input)
	if len(input) != 1 || input[0] < '1' || input[0] > '9' {
		fmt.Println("Ungültige Eingabe. Bitte gib eine Zahl zwischen 1 und 9 ein.")
		g.Move(player)
		return
	}

	row, col := (input[0]-'1')/3, (input[0]-'1')%3
	if !g.MoveAllowed(int(row), int(col)) {
		fmt.Println("Dieser Zug ist nicht erlaubt. Bitte wähle eine andere Position.")
		g.Move(player)
		return
	}

	g.Set(int(row), int(col), player)
}
