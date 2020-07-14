package doc

import (
	"testing"
)

func TestCreateLong(t *testing.T) {
	result1 := CreateLong(
		`
____ ____ _  _ ____ _ ____
|    |  | |\ | |___ | | __
|___ |__| | \| |    | |__]
`,
	)
	expected1 := "\n  ____ ____ _  _ ____ _ ____\n  |    |  | |\\ | |___ | | __\n  |___ |__| | \\| |    | |__]\n  ⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯\n\n  🐙 GitHub: https://github.com/Matt-Gleich/btask\n  📜 Produced under the MIT license\n  🛠  Authors:\n\t- Matthew Gleich (https://github.com/Matt-Gleich)\n\n"
	if result1 != expected1 {
		t.Errorf(result1, "!= ", expected1)
	}

	result2 := CreateLong(
		`
___  ___ ____ ____ _  _
|__]  |  |__| [__  |_/
|__]  |  |  | ___] | \_
`,
	)
	expected2 := "\n  ___  ___ ____ ____ _  _\n  |__]  |  |__| [__  |_/\n  |__]  |  |  | ___] | \\_\n  ⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯\n\n  🐙 GitHub: https://github.com/Matt-Gleich/btask\n  📜 Produced under the MIT license\n  🛠  Authors:\n\t- Matthew Gleich (https://github.com/Matt-Gleich)\n\n"
	if result2 != expected2 {
		t.Errorf(result2, "!= ", expected2)
	}
}
