package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: square root of negative number")
	}

	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}

	return nil
}

func main() {
	res, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	res1, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res1)

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Data proceesed successfully!")
}
