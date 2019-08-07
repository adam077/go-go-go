package utils

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	t1()
}

func t1() {
	fmt.Println(GetUUID())
}
