package checing

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {

		fmt.Println("ERROR:", e)
		os.Exit(0)

	}
}