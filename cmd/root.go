package cmd

import (
	"fmt"
	"kakau/k8s"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
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
	home := homedir.HomeDir();
	filePath := filepath.Join(home, ".kube", "config")


	config := k8s.GetConfig(filePath)
	contexts := k8s.FetchAllContexts(config)

	fmt.Println("available contexts:")
	for _, context := range contexts {
		fmt.Println(context)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}		