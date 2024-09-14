package golangStudy

import (
	"fmt"
)

func HelloGolangStudy(name string) string {
	if len(name) == 0 {
		name = `阿呆`
	}
	return fmt.Sprintf(`hello %s`, name)
}
