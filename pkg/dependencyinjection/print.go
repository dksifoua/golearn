package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(w io.StringWriter, name string) (err error) {
	_, err = w.WriteString(fmt.Sprintf("Hello, %s!\n", name))
	return
}
