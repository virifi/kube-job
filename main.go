package main

import (
	"fmt"
	"os"

	"github.com/h3poteto/kube-job/cmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // Support GKE
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
