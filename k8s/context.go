package k8s

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetConfig(filePath string) *api.Config {	
	config, err := clientcmd.LoadFromFile(filePath)
	if err != nil {
		panic(err)
	}
	return config;
}

func FetchAllContexts(config *api.Config) []string {
	var contexts []string
	for name := range config.Contexts {
		contexts = append(contexts, name)
	}
	return contexts
}