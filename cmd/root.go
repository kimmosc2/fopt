package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	version bool
)

var (
	Version = `fopt-0.1.1-Alpha`
	Author  = `BuTn<http://github.com/kimmosc2>`
)

var rootCmd = &cobra.Command{
	Use:   "fopt",
	Short: "operate your file",
	Long:  `fopt is a file operator, delete, rename and so on`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			setVersion()
			os.Exit(0)
		}
		cmd.Usage()
		os.Exit(0)
	},
}

func setVersion() {
	fmt.Printf("version:%s\n"+
		"author:%s\n", Version, Author)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "fopt info")
	rootCmd.Flags().BoolVarP(&version, "info", "i", false, "fopt info")
	rootCmd.AddCommand(deleteCmd)
}
