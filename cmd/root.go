package cmd

import (
	"fmt"
	"kakau/components"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kakao",
	Short: "kakao is a multi context sql query execution tool",
	Run: func(cmd *cobra.Command, args []string) {
		Welcome()
		Kakao()
	},
}
  
func Kakao() {	
	selectedContexts := components.LoadSelectedContexts()
	fmt.Println(selectedContexts)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}		