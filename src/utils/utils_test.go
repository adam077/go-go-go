package utils

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	t2()
}

func t1() {
	fmt.Println(GetUUID())
}

func t2() {
	fmt.Println(FindBetween([]byte("asdfgyhjkfsdj"), "f", "j"))
	// fghj   fsdj
}
