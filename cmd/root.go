package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
)

var ProgramInfo struct {
	PlatForm  string
	Version   string
	GoVersion string
	Author    string
	BuildTime string
}
var (
	version         bool
	versionTemplate = `version:fopt-{{.Version}}-{{.PlatForm}}
go version:{{.GoVersion}}
author:{{.Author}}
build time :{{.BuildTime}}
`
)

var rootCmd = &cobra.Command{
	Use:   "fopt",
	Short: "operate your file",
	Long:  `fopt is a file operator, delete, rename and so on`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			setVersion()
		}
	},
}

func setVersion() {
	vt := template.New("versionTpl")
	parse, err := vt.Parse(versionTemplate)
	if err != nil {
		panic(err)
	}
	err = parse.Execute(os.Stdout, ProgramInfo)
	panic(err)
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
