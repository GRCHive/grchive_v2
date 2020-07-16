package gin_backend_utility

import "fmt"

type WebappError struct {
	Err     error
	Context string
}

func (w *WebappError) Error() string {
	if w.Err != nil {
		return fmt.Sprintf("ERR: %s - [%s]", w.Context, w.Err.Error())
	} else {
		return fmt.Sprintf("ERR: %s", w.Context)
	}
}
