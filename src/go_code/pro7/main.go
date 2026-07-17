package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// 生成一个新的 UUID
	id := uuid.New()
	fmt.Println("生成的 UUID 是:", id.String())
}
