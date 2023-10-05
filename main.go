package main

import (
	"errors"
	"fmt"

	"gitea.com/lunny/log"
)

func getName() string {
	log.Errorf("i am err: %v", errors.New("hello"))
	return "alex"
}

func main() {

	fmt.Println(getName())
}
