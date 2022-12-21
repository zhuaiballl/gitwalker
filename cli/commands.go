package cli

import (
	"fmt"
	"github.com/otiai10/copy"
	"log"
	"os"
	"os/exec"
	"strings"
)

func (cli *CLI) countTag() {
	cmd := exec.Command("git", "tag")
	tagOut, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}
	tags := strings.Split(string(tagOut), "\n")
	fmt.Println(len(tags) - 1)
}

func (cli *CLI) walk(bare bool) {
	cmd := exec.Command("git", "rev-list", "--reverse", "HEAD")
	commitOut, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}
	commits := strings.Split(string(commitOut), "\n")
	commits = commits[:len(commits)-1]
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	cur, err := os.Getwd()
	for i, commit := range commits {
		fmt.Println(commit)
		cmd = exec.Command("git", "checkout", commit)
		err = cmd.Run()
		folderName := fmt.Sprintf("%04d", i)
		if !bare {
			folderName += "_" + commit
		}
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

func (cli *CLI) walkByTag(bare bool) {
	cmd := exec.Command("git", "tag", "-l", "--sort=v:refname")
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
		fmt.Println(tag)
		cmd = exec.Command("git", "checkout", tag)
		err = cmd.Run()
		folderName := fmt.Sprintf("%04d", i)
		if !bare {
			folderName += "_" + tag
		}
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
