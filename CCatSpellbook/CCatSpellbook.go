package main

import (
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CopyFile(src string, dst string) {
	// Reads src and writes it to dst.
	data, err := os.ReadFile(src)
	CheckErr(err)
	err = os.WriteFile(dst, data, 0644)
	CheckErr(err)
}
