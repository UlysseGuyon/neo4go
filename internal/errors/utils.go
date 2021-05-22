package errors

import "fmt"

func errorFmt(t string, err string) string {
	return fmt.Sprintf(
		`[Neo4Go %s Error]: %s`,
		t,
		err,
	)
}
