package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// create the in-cluster config
	inClusterConfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// create a clientSet
	clientset, err := kubernetes.NewForConfig(inClusterConfig)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(pods.Items)
}
