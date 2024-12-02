package k8sclient

import (
	"log"
	// "testing"
	// "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func MockInitK8sClient() kubernetes.Interface {
	log.Printf("MockInitK8sClient called")
	clientset := fake.NewSimpleClientset()
	return clientset
}