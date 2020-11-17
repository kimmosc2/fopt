package cmd

import (
	"errors"
	"fmt"
	"github.com/kimmosc2/file-operator/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
	"regexp"
)

var (
	// regular expression
	reg string
	// file directory
	fileDir string
	// 	TODO: file suffix
	suffix string
	// readonly, not operate
	unsafeMode bool
)

// deleteCmd is a cobra command
var deleteCmd = &cobra.Command{
	Use:   "delete ",
	Short: "delete the specified file",
	Long:  "delete file, support regular expression",
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkDeleteParameter(); err != nil {
			fmt.Println("Error:", err.Error()+"\n")
			cmd.Usage()
			os.Exit(1)
		}

		compile, err := regexp.Compile(reg)
		if err != nil {
			log.Fatalf("compile regular expression error:%s", err)
		}
		if err := internal.Walk(fileDir, compile, unsafeMode); err != nil {
			log.Fatalf("Walk error: %s", err)
		}
	},
}

// checkDeleteParameter check deleteCmd parameter,
// if there is an error,it will be return an custom error
func checkDeleteParameter() error {
	if reg == "" {
		return errors.New("empty expression")
	}
	if fileDir == "" {
		return errors.New("no specified directory name")
	}
	if stat, _ := os.Stat(fileDir); !stat.IsDir() {
		return errors.New(fileDir + " is not a directory")
	}
	return nil
}

func init() {
	// regular expression
	deleteCmd.Flags().StringVarP(&reg, "regexp", "r", "", "regular expression")
	// file directory
	deleteCmd.Flags().StringVarP(&fileDir, "dir", "d", "", "target directory")
	// file suffix
	// deleteCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "file suffix,suffix and regexp cannot be set at the same time")
	// unsafe mode
	deleteCmd.Flags().BoolVarP(&unsafeMode, "unsafe", "u", false, "unsafe mode, if use this flag, the delete operation will be performed")
}

func trimQuote(s string) string {
	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-2]
	}
	return s
}
