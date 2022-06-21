package cli

import (
	"fmt"
	"github.com/otiai10/copy"
	"log"
	"os"
	"os/exec"
	"strings"
)

func (cli *CLI) walk() {
	depth := 1
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
		h := fmt.Sprintf("%04d", depth) + fmt.Sprint(string(head[:7]))
		depth++
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
		//skipList := map[string]bool{
		//	"gitwalker": true,
		//	".git":      true,
		//	".idea":     true,
		//	".DS_Store": true,
		//}
		copyOpt := copy.Options{
			Skip: func(src string) (bool, error) {
				return strings.HasSuffix(src, ".git") || strings.HasSuffix(src, "gitwalker"), nil
			},
		}
		err = copy.Copy(cur, dir, copyOpt)
		if err != nil {
			log.Panic(err)
		}
		//curLen := len(cur)
		//err = filepath.Walk(cur, func(path string, info fs.FileInfo, err error) error {
		//	fmt.Println("current path is", path)
		//	if err != nil {
		//		return err
		//	}
		//	if _, fd := skipList[info.Name()]; fd {
		//		if info.IsDir() {
		//			return filepath.SkipDir
		//		} else {
		//			return nil
		//		}
		//	}
		//	destPath := dir + path[curLen:]
		//	_, err = os.Stat(destPath)
		//	if util.Exist(err) {
		//		return nil
		//	} else {
		//		if info.IsDir() {
		//			err = os.Mkdir(destPath, os.ModePerm)
		//			if err != nil {
		//				log.Panic(err)
		//				return err
		//			}
		//		} else {
		//			util.Copy(path, destPath)
		//		}
		//	}
		//	return nil
		//})
	}
}

func (cli *CLI) walkByTag() {
	cmd := exec.Command("git", "tag")
	tagOut, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}
	tags := strings.Split(string(tagOut), "\n")
	tags = tags[:len(tags)-1]
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	cur, err := os.Getwd()
	for i, tag := range tags {
		fmt.Println(tag, "QAQ")
		cmd = exec.Command("git", "checkout", tag)
		folderName := fmt.Sprintf("%04d_", i) + tag
		copyOpt := copy.Options{
			Skip: func(src string) (bool, error) {
				return strings.HasSuffix(src, ".git") || strings.HasSuffix(src, "gitwalker"), nil
			},
		}
		dir := homedir + "/.gitwalker/" + folderName
		err = copy.Copy(cur, dir, copyOpt)
		if err != nil {
			log.Panic(err)
		}
	}

}
