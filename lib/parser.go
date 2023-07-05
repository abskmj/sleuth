package lib

import (
	"fmt"
)

func Parse(name string) {
	fmt.Println("Name:", name)

	var sanitized = sanitize(name)

	fmt.Println("Sanitized:", sanitized)
}
