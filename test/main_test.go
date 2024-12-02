package main

import (
	//"fmt"
	"log"
	k8sclient "my/k8sclient"
	//"k8s.io/client-go/kubernetes/fake"
	// cmd "my/cmd"
	"testing"
)

//var realInitK8sClient = k8sclient.InitK8sClient  // returns *kubernetes.Clientset
var mockInitK8sClient = k8sclient.MockInitK8sClient

func TestInitK8sClient(t *testing.T) {
	log.Println("k8sclient/test/main_test.go")
	kubeClient := mockInitK8sClient()
	if kubeClient == nil {
		log.Fatal("mockInitK8sClient Failed")
	}
	// assert.NotNil(kubeClient, "should not be nil")
}