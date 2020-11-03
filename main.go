package main

import (
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err.Error())
	}
	clientset, _ := kubernetes.NewForConfig(config)
	command := []string{
		"sh",
		"-c",
		"ls -a",
		"echo $(date +%s) > heure",
		"echo $heure"
	}
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name("nexus-sonatype-nexus-6fb649db5c-xshlj").
		Namespace("default").
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: "nexus",
			Command:   command,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println()

}
