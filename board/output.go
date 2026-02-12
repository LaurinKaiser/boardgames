package board

import (
	"fmt"
	"strings"
)

// String liefert eine menschenlesbare Darstellung des Spielfelds.
func (b *Board) String() string {
	// Hinweis:
	// - Berechnen sie einen für die oberen/unteren Ränder und Trennlinien.
	// - Für größere Spielfelder sollten die Zellen automatisch breiter werden, damit der Inhalt lesbar bleibt.
	//   Verwenden Sie die Methode MaxCellWidth, um die Breite der Zellen zu bestimmen.
	// - Erstellen Sie eine Liste von Zeilen, die die Zeilen des Spielfelds repräsentieren.
	// - Bauen Sie in einer Schleife jeweils die aktuelle Zeile auf und fügen Sie dann
	//   die Zeile sowie die Trennlinie zur Zeilen-Liste hinzu.
	// - Am Ende können Sie die Zeilen mit einem Zeilenumbruch zusammenfügen und zurückgeben.

	cell_width := b.MaxCellWidth()
	v_border := strings.Repeat(fmt.Sprintf("+-%s-", strings.Repeat("-", cell_width)), b.Width()) + "+"

	lines := []string{v_border}
	for _, row := range b.rows {
		line := "|"
		for _, cell := range row {
			line += fmt.Sprintf(" %s |", fmt.Sprintf("%*s", cell_width, cell))
		}
		lines = append(lines, line)
		lines = append(lines, v_border)
	}
	return strings.Join(lines, "\n")
}
