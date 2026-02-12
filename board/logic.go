package board

import "boardgames/util"

// RowContains prüft, ob die angegebene Zeile den gesuchten Wert enthält.
func (b *Board) RowContains(i int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.Contains und die Methode Row.
	return util.Contains(b.Row(i), value)
}

// RowContainsChain prüft, ob die angegebene Zeile eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) RowContainsChain(rowIndex int, value string, length int) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsChain und die Methode Row.
	return util.ContainsChain(b.Row(rowIndex), value, length)
}

// RowContainsOnly prüft, ob die angegebene Zeile ausschließlich den gesuchten Wert enthält.
func (b *Board) RowContainsOnly(rowIndex int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsOnly und die Methode Row.
	return util.ContainsOnly(b.Row(rowIndex), value)
}

// ColContains prüft, ob die angegebene Spalte den gesuchten Wert enthält.
func (b *Board) ColContains(colIndex int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.Contains und die Methode Col.
	return util.Contains(b.Col(colIndex), value)
}

// ColContainsChain prüft, ob die angegebene Spalte eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) ColContainsChain(colIndex int, value string, length int) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsChain und die Methode Col.
	return util.ContainsChain(b.Col(colIndex), value, length)
}

// ColContainsOnly prüft, ob die angegebene Spalte ausschließlich den gesuchten Wert enthält.
func (b *Board) ColContainsOnly(colIndex int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsOnly und die Methode Col.
	return util.ContainsOnly(b.Col(colIndex), value)
}

// DiagDownRightContains prüft, ob die angegebene Diagonale von oben links nach unten rechts den gesuchten Wert enthält.
func (b *Board) DiagDownRightContains(startCol int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.Contains und die Methode DiagDownRight.
	return util.Contains(b.DiagDownRight(startCol), value)
}

// DiagDownRightContainsChain prüft, ob die angegebene Diagonale von oben links nach unten rechts eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) DiagDownRightContainsChain(startCol int, value string, length int) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsChain und die Methode DiagDownRight.
	return util.ContainsChain(b.DiagDownRight(startCol), value, length)
}

// DiagDownRightContainsOnly prüft, ob die angegebene Diagonale von oben links nach unten rechts ausschließlich den gesuchten Wert enthält.
func (b *Board) DiagDownRightContainsOnly(startCol int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsOnly und die Methode DiagDownRight.
	return util.ContainsOnly(b.DiagDownRight(startCol), value)
}

// DiagDownLeftContains prüft, ob die angegebene Diagonale von oben rechts nach unten links den gesuchten Wert enthält.
func (b *Board) DiagDownLeftContains(startCol int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.Contains und die Methode DiagDownLeft.
	return util.Contains(b.DiagDownLeft(startCol), value)
}

// DiagDownLeftContainsChain prüft, ob die angegebene Diagonale von oben rechts nach unten links eine ununterbrochene Kette des gesuchten Werts enthält.
func (b *Board) DiagDownLeftContainsChain(startCol int, value string, length int) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsChain und die Methode DiagDownLeft.
	return util.ContainsChain(b.DiagDownLeft(startCol), value, length)
}

// DiagDownLeftContainsOnly prüft, ob die angegebene Diagonale von oben rechts nach unten links ausschließlich den gesuchten Wert enthält.
func (b *Board) DiagDownLeftContainsOnly(startCol int, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion util.ContainsOnly und die Methode DiagDownLeft.
	return util.ContainsOnly(b.DiagDownLeft(startCol), value)
}

// BoardContains prüft, ob das Spielfeld den gesuchten Wert enthält.
func (b *Board) BoardContains(value string) bool {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um jede Zeile des Spielfelds zu überprüfen.
	for i := range b.rows {
		if b.RowContains(i, value) {
			return true
		}
	}
	return false
}

// BoardContainsChain prüft, ob das Spielfeld eine ununterbrochene Kette des gesuchten Werts enthält.
// Dabei werden alle Zeilen, Spalten und Diagonalen überprüft.
func (b *Board) BoardContainsChain(value string, length int) bool {
	// Hinweis:
	// Verwenden Sie for-Schleifen, um jede Zeile, Spalte und Diagonale des Spielfelds zu überprüfen.
	for i := range b.rows {
		if b.RowContainsChain(i, value, length) {
			return true
		}
	}
	for i := range b.rows[0] {
		if util.ContainsChain(b.Col(i), value, length) {
			return true
		}
	}
	for i := -b.Height() + 1; i < b.Width(); i++ {
		if util.ContainsChain(b.DiagDownRight(i), value, length) {
			return true
		}
	}
	for i := -b.Height() + 1; i < b.Width(); i++ {
		if util.ContainsChain(b.DiagDownLeft(i), value, length) {
			return true
		}
	}
	return false
}

// BoardContainsOnly prüft, ob das Spielfeld ausschließlich den gesuchten Wert enthält.
func (b *Board) BoardContainsOnly(value string) bool {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um jede Zeile des Spielfelds zu überprüfen.
	// Wenn eine Zeile nicht ausschließlich den gesuchten Wert enthält, können Sie sofort false zurückgeben.
	for i := range b.rows {
		if !b.RowContainsOnly(i, value) {
			return false
		}
	}
	return true
}
