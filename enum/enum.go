package enum

import "fmt"

type enumer interface {
	fmt.Stringer
	Invalid() bool
}

func DefaultInvalid(e int, invalid int) bool {
	return e >= invalid
}
