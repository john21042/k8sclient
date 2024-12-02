package main

import (
	"fmt"
	"log"
	k8sclient "my/k8sclient"
	cmd "my/cmd"
)

var myInitK8sClient = k8sclient.InitK8sClient  // returns *kubernetes.Clientset
var myAdRemoveCmd = cmd.AdRemoveCmd

func main() {
	fmt.Println("k8sclient/test/main.go")
	clientset := myInitK8sClient()
	if clientset == nil {
		log.Fatal("main InitK8sClient Failed")
	} else {
		log.Printf("main InitK8sClient OK")
	}

	myAdRemoveCmd()
}

