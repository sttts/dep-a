package a

import (
	"fmt"

	"github.com/sttts/dep-b"
)

func A() string {
	return fmt.Sprintf("release of dep-a calling %s of dep-b", b.B())
}
