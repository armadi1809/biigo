package langerror

import (
	"fmt"
)

type LangError struct {
	Message string
	Where   string
	Line    int
}

func (e *LangError) Error() string {
	return fmt.Sprintf("[Line %d] Error%s: %s\n", e.Line, e.Where, e.Message)
}
