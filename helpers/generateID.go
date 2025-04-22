package helpers

import (
	"fmt"

	"github.com/sony/sonyflake"
)

func GenerateID() uint {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		fmt.Println("Sonyflake not created")
		return 0
	}

	id, err := sf.NextID()
	if err != nil {
		fmt.Println("Failed to generate ID:", err)
		return 0
	}
	return uint(id)
}
