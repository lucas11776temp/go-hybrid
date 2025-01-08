package path

import (
	"os"
)

func FileExits(filename string) bool {
	_, err := os.Open(filename)

	if err != nil {
		return false
	}

	return true
}
