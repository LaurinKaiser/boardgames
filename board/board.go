package board

import "fmt"

// Ein Spielfeld für Brettspiele.
type Board struct {
	rows [][]string
}

// New erstellt ein neues leeres Spielfeld mit der angegebenen Anzahl von Zeilen und Spalten.
// Die Zellen des Spielfelds werden mit Leerzeichen initialisiert.
func New(rows, cols int) *Board {
	// Hinweis:
	// Erstellen Sie eine 2D-Liste mit der angegebenen Anzahl von Zeilen und Spalten.
	// Sie können die Hilfsfunktion make verwenden, um die äußere Liste zu erstellen,
	// und dann eine Schleife, um die inneren Listen zu erstellen.
	// Die Schleife muss dabei die innere Liste mit Leerzeichen initialisieren
	// und dann in jedes Element ein Zeichen " " einfügen.
	b := &Board{
		rows: make([][]string, rows),
	}
	for i := range b.rows {
		b.rows[i] = make([]string, cols)
		for j := range b.rows[i] {
			b.rows[i][j] = " "
		}
	}
	return b
}

// NewWithNumbers erstellt ein neues Spielfeld mit der angegebenen Anzahl von Zeilen und Spalten.
// Füllt die Zellen des Spielfelds mit aufsteigenden Zahlen als Strings, beginnend bei "1".
func NewWithNumbers(rows, cols int) *Board {
	// Hinweis:
	// Gehen Sie analog zu New vor, aber anstatt die Zellen mit Leerzeichen zu initialisieren,
	// füllen Sie sie mit aufsteigenden Zahlen als Strings.
	// Verwenden Sie dafür zusätzlich eine Zählvariable.
	b := &Board{
		rows: make([][]string, rows),
	}
	num := 1
	for i := range b.rows {
		b.rows[i] = make([]string, cols)
		for j := range b.rows[i] {
			b.rows[i][j] = fmt.Sprintf("%d", num)
			num++
		}
	}
	return b
}

// Get liefert den Inhalt der Zelle an der angegebenen Position zurück.
// Liefert einen leeren String, falls die Position außerhalb des Spielfelds liegt.
func (b *Board) Get(row, col int) string {
	// Hinweis:
	// Sie können direkt auf das Array b.rows zugreifen, um den Inhalt der Zelle zu erhalten.
	// Überprüfen Sie aber vorher, ob die angegebenen Zeilen- und Spaltenindizes
	// innerhalb der Grenzen des Spielfelds liegen.
	if row < 0 || row >= b.Height() || col < 0 || col >= b.Width() {
		return ""
	}
	return b.rows[row][col]
}

// Set setzt den Inhalt der Zelle an der angegebenen Position auf den angegebenen Wert.
// Ignoriert die Anweisung, falls die Position außerhalb des Spielfelds liegt.
func (b *Board) Set(row, col int, value string) {
	// Hinweis:
	// Sie können direkt auf das Array b.rows zugreifen, um den Inhalt der Zelle zu setzen.
	// Überprüfen Sie aber vorher, ob die angegebenen Zeilen- und Spaltenindizes
	// innerhalb der Grenzen des Spielfelds liegen.
	if row < 0 || row >= b.Height() || col < 0 || col >= b.Width() {
		return
	}
	b.rows[row][col] = value
}
