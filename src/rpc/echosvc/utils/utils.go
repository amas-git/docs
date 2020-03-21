package utils

import (
	"fmt"

	"amas.org/echosvc/model"
)

func HelloWorld() {
	fmt.Println("HELLO WOLRD!")
	x := model.Msg_HIGH
	fmt.Println(x)
}
