package greeting

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	if name == "" {
		return "Hello, my friend."
	}

	result := fmt.Sprintf("Hello, %s.", name)

	if name == strings.ToUpper(name) {
		result = strings.ToUpper(result)
	}

	return result
}
