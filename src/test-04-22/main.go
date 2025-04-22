package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/morari-roman/test-04-22/asana/internal/service"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("is running")

	just_test := service.New()

	go func() {
		defer wg.Done()
		err := just_test.GetData()
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	fmt.Println("is done")

}
