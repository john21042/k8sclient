package k8sclient

import (
	"log"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func InitK8sClient() kubernetes.Interface {  // *kubernetes.Clientset
	log.Printf("InitK8sClient start")
	// setup k8s client
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatal("Failed to load KUBECONFIG -> ", err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	log.Printf("InitK8sClient done")
	return clientset
}