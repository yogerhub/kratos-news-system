package data

import (
	"fmt"
)

func demoKey(id int64) string {
	return fmt.Sprintf("demo:id:%d", id)
}
