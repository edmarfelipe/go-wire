package main

import (
	"fmt"
)

func main() {
	usecase, _ := InitializeUsecase()

	err := usecase.Update()
	if err != nil {
		fmt.Println(err)
	}
}
