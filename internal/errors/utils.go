package errors

import (
	"fmt"
)

// errorFmt formats an error string by showing its type as prefix
func errorFmt(t string, err string) string {
	return fmt.Sprintf(
		`[Neo4Go %s Error]: %s`,
		t,
		err,
	)
}
