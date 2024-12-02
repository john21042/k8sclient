package cmd

import (
	"fmt"
	"log"
	k8sclient "my/k8sclient"
)

var getInitK8sClient = k8sclient.InitK8sClient

func AdRemoveCmd() {
	fmt.Println("adRemoveCmd begin")
	clientset := getInitK8sClient()
	if clientset == nil {
		log.Fatal("adRemoveCmd getInitK8sClient Failed")
	} else {
		log.Printf("adRemoveCmd getInitK8sClient OK")
	}
}