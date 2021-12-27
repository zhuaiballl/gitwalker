package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func main() {
	for {
		cmd := exec.Command("git", "checkout", "HEAD^")
		err := cmd.Run()
		if err != nil {
			log.Panic(err)
		}
		head, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		if err != nil {
			log.Panic(err)
		}
		h := fmt.Sprint(string(head[:7]))
		fmt.Println("commitID:", h)
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Panic(err)
		}
		dir := homedir + "/.gitwalker/" + h
		//fmt.Println(dir)
		err = os.MkdirAll(dir, os.ModePerm)
		// err = exec.Command("cp", ".", "dir").Run()
		cur, err := os.Getwd()
		if err != nil {
			log.Panic(err)
		}
		skipList := map[string]bool{
			"gitwalker": true,
			".git":      true,
			".idea":     true,
			".DS_Store": true,
		}
		curLen := len(cur)
		err = filepath.Walk(cur, func(path string, info fs.FileInfo, err error) error {
			fmt.Println("current path is", path)
			if err != nil {
				return err
			}
			if _, fd := skipList[info.Name()]; fd {
				if info.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
			destPath := dir + path[curLen:]
			_, err = os.Stat(destPath)
			if Exist(err) {
				return nil
			} else {
				if info.IsDir() {
					err = os.Mkdir(destPath, os.ModePerm)
					if err != nil {
						log.Panic(err)
						return err
					}
				} else {
					Copy(path, destPath)
				}
			}
			return nil
		})
	}
	fmt.Println("Hello world!")
}
