package doc

import (
	"strings"
)

// CreateLong ... Create the long felid for the command
func CreateLong(asciiTitle string) string {
	indentation := "  "
	indentedTitle := strings.Join(strings.Split(asciiTitle, "\n"), "\n"+indentation)

	// Creating underline
	dots := []string{}
	titleLines := strings.Split(asciiTitle, "\n")
	for i := 0; i < len(titleLines[len(titleLines)-2]); i++ {
		dots = append(dots, "⋯")
	}

	return indentedTitle + strings.Join(dots, "") + `

  🐙 GitHub: https://github.com/Matt-Gleich/btask
  📜 Produced under the MIT license
  🛠  Authors:
	- Matthew Gleich (https://github.com/Matt-Gleich)

`
}
