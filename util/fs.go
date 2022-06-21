package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Exist uses err returned from os.Stat to determine if a file/folder exists
func Exist(err error) bool {
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			log.Panic(err)
		}
	}
	return true
}

func Copy(from, to string) {
	fmt.Println("Copy from", from, "to", to)
	src, err := os.Open(from)
	if err != nil {
		log.Panic(err)
	}
	defer src.Close()
	dst, err := os.Create(to)
	if err != nil {
		log.Panic(err)
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Panic(err)
	}
}
