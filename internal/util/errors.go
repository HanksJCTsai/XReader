package util

import (
	"fmt"
)

func WrapError(operation string, err error) error {
	return fmt.Errorf("[XReader] %s Operation failed: %v", operation, err)
}
