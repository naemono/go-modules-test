package main

import "os"

func main() {
	err := Public()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func Public() error {
	return nil
}
