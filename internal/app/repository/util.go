package repository

import (
	"fmt"
	"strings"
)

func replacePlaceholders(query string) string {
	// Counter to keep track of the placeholder index
	index := 0

	// Split the query string by ?, and join it back with $n where n is the index
	parts := strings.Split(query, "?")
	convertedQuery := strings.Builder{}
	for i, part := range parts {
		// Append the current part
		convertedQuery.WriteString(part)
		// If there is a placeholder to replace, append $n
		if i < len(parts)-1 {
			index++
			convertedQuery.WriteString(fmt.Sprintf("$%d", index))
		}
	}

	return convertedQuery.String()
}
