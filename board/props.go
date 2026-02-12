package board

// MaxCellWidth bestimmt die maximale Breite des Inhalts einer Zelle.
// Dies ist nützlich, um die Spaltenbreite für die Ausgabe zu berechnen.
func (b *Board) MaxCellWidth() int {
	// Hinweis:
	// Gehen Sie durch alle Zellen des Spielfelds und bestimmen Sie die maximale Länge des Inhalts.
	// Prüfen Sie dabei für jede Zelle, ob die Länge des Inhalts größer ist als die bisher gefundene maximale Länge.
	maxWidth := 0
	for _, row := range b.rows {
		for _, cell := range row {
			if len(cell) > maxWidth {
				maxWidth = len(cell)
			}
		}
	}
	return maxWidth
}

// / Width bestimmt die Breite des Spields. D.h. die Anzahl der Spalten.
func (b *Board) Width() int {
	// Hinweis:
	// Die Breite des Spielfelds entspricht der Anzahl der Spalten in einer Zeile.
	// Da wir annehmen, dass das Spielfeld rechteckig ist, können Sie die Anzahl
	// der Spalten in der ersten Zeile zurückgeben.
	// Überprüfen Sie aber vorher, ob das Spielfeld überhaupt Zeilen enthält, um einen Fehler zu vermeiden.
	if len(b.rows) == 0 {
		return 0
	}
	return len(b.rows[0])
}

// Height bestimmt die Höhe des Spields. D.h. die Anzahl der Zeilen.
func (b *Board) Height() int {
	// Hinweis:
	// Die Höhe des Spielfelds entspricht der Anzahl der Zeilen, also der Länge des Arrays b.rows.
	return len(b.rows)
}

// Row liefert die Zeile an der angegebenen Position zurück.
func (b *Board) Row(index int) []string {
	// Hinweis:
	// Sie können direkt auf das Array b.rows zugreifen, um die Zeile zurückzugeben.
	return b.rows[index]
}

// Col liefert die Spalte an der angegebenen Position zurück.
func (b *Board) Col(index int) []string {
	// Hinweis:
	// Erstellen Sie ein neues Array, das die Spalte repräsentiert, und füllen Sie es mit den entsprechenden Werten aus den Zeilen.
	col := make([]string, len(b.rows))
	for i, row := range b.rows {
		col[i] = row[index]
	}
	return col
}

// DiagDownRight liefert eine Diagonale von oben links nach unten rechts zurück.
func (b *Board) DiagDownRight(startCol int) []string {
	// Hinweis:
	// Erstellen Sie ein neues Array, das die Diagonale repräsentiert, und füllen Sie es mit den entsprechenden Werten aus den Zeilen.
	diag := []string{}
	for i, j := 0, startCol; i < b.Height(); i, j = i+1, j+1 {
		if j >= 0 && j < b.Width() {
			diag = append(diag, b.rows[i][j])
		}
	}
	return diag
}

// DiagDownLeft liefert eine Diagonale von oben rechts nach unten links zurück.
func (b *Board) DiagDownLeft(startCol int) []string {
	// Hinweis:
	// Erstellen Sie ein neues Array, das die Diagonale repräsentiert, und füllen Sie es mit den entsprechenden Werten aus den Zeilen.
	diag := []string{}
	for i, j := 0, startCol; i < b.Height(); i, j = i+1, j-1 {
		if j >= 0 && j < b.Width() {
			diag = append(diag, b.rows[i][j])
		}
	}
	return diag
}
