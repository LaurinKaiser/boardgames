package util

// Contains prüft, ob eine Liste den gesuchten Wert enthält.
func Contains(list []string, value string) bool {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um jedes Element der Liste zu überprüfen.
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// ContainsChain prüft, ob eine Liste eine ununterbrochene Kette des gesuchten Werts enthält.
func ContainsChain(list []string, value string, length int) bool {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um jedes Element der Liste zu überprüfen.
	// Zählen Sie dabei, wie viele aufeinanderfolgende Elemente dem gesuchten Wert entsprechen.
	// Sobald Sie einmal ein Element finden, das nicht dem gesuchten Wert entspricht, setzen Sie den Zähler zurück.
	count := 0
	for _, item := range list {
		if item == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// ContainsOnly prüft, ob eine Liste ausschließlich den gesuchten Wert enthält.
func ContainsOnly(list []string, value string) bool {
	// Hinweis:
	// Verwenden Sie die Funktion ContainsChain mit einer passenden Länge.
	return ContainsChain(list, value, len(list))
}
