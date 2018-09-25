package main

import (
	"flag"
	"fmt"
	"os"

	clientset "github.com/heptio/contour/apis/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
	inCluster  bool
)

func init() {
	flag.StringVar(&kubeconfig, "kubecfg-file", "", "Location of kubecfg file for access to gimbal system kubernetes api, defaults to service account tokens")
	flag.BoolVar(&inCluster, "in-cluster", false, "use in cluster configuration.")
	flag.Parse()
}

func main() {
	contourClient := newClient(kubeconfig, inCluster)
	routes, err := contourClient.ContourV1beta1().IngressRoutes("default").List(metav1.ListOptions{})
	check(err)

	for _, v := range routes.Items {
		fmt.Println("Got Routes: ", v)
	}
}

func newClient(kubeconfig string, inCluster bool) *clientset.Clientset {
	var err error
	var config *rest.Config
	if kubeconfig != "" && !inCluster {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		check(err)
	} else {
		config, err = rest.InClusterConfig()
		check(err)
	}

	contourClient, err := clientset.NewForConfig(config)
	check(err)
	return contourClient
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
