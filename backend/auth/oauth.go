package auth

import (
	"fmt"

	"golang.org/x/oauth2"
)

// just so my linter doesn't delete the import
func fu() {
	var x oauth2.Config
	fmt.Println(x)
}
