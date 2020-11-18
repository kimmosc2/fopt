// TODO: refactor me
package internal

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// RegWalk call filepath.Walk to delete file by reg,
// if safe is true, RegWalk just find them but nothing to do
func RegWalk(fileDir string, reg *regexp.Regexp, unsafe bool) error {
	if unsafe {
		log.Println("[Waring]: you are now in unsafe mode, use \"fopt delete --help\" learn more")
	} else {
		log.Println("[Info]:you are now in unsafe mode, use \"fopt delete --help\" learn more")
	}
	return filepath.Walk(fileDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if reg.MatchString(path) {
			log.Printf("find %s", path)
			// unsafe mode
			if unsafe {
				if err := os.Remove(path); err != nil {
					log.Println("delete file error: ", err)
					return nil
				}
				log.Printf("%s has been deleted", path)
			}
			return nil
		}
		return nil
	})
}

// SuffixWalk is a copy of RegWalk,it has a very simply implementation,
// call strings.HasSuffix()
func SuffixWalk(fileDir string, suffix string, unsafe bool) error {
	if unsafe {
		log.Println("[Waring]: you are now in unsafe mode, use \"fopt delete --help\" learn more")
	} else {
		log.Println("[Info]:you are now in unsafe mode, use \"fopt delete --help\" learn more")
	}
	return filepath.Walk(fileDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, suffix) {
			log.Printf("find %s", path)
			// unsafe mode
			if unsafe {
				if err := os.Remove(path); err != nil {
					log.Println("delete file error: ", err)
					return nil
				}
				log.Printf("%s has been deleted", path)
			}
			return nil
		}
		return nil
	})
}
