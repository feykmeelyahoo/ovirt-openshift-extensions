/*
Copyright 2019 oVirt-maintainers

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/util/file"
//	"k8s.io/client-go/rest"
)

func getSystemUUIDByNodeName(nodeName string) (string, error) {
	fmt.Printf("SERKAN -- getSystemUUIDByNodeName - node name %s. \n", nodeName)
	nodes, e := getKubeNodes()
	if e != nil {
		return "", e
	}
	for _, n := range nodes {
		if n.Name == nodeName {
			return n.Status.NodeInfo.SystemUUID, nil
		}
	}
	return "", fmt.Errorf("node name %s was not found", nodeName)
}

func getKubeNodes() ([]v1.Node, error) {
	kubeconfig, err := sblocateKubeConfig()


        fmt.Printf("SERKAN -- getKubeNodes !!!")
        fmt.Printf("SERKAN -- getKubeNodes  config => %s. \n", kubeconfig)

	if err != nil {
		return nil, err
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodes.Items, nil
}

/////////////////////////////// SERKAN ////////////////
func sblocateKubeConfig() (string, error) {
	fmt.Printf("SERKAN -- sblocateKubeConfig !!!. \n")
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	fmt.Printf("SERKAN -- sblocateKubeConfig Cıkıyor  !!!. %s ----  \n", kubeconfig)

	return *kubeconfig, nil
}
/////////////////////////////// SERKAN ////////////////

func locateKubeConfig() (string, error) {
	fmt.Printf("SERKAN -- locateKubeConfig !!!")

	defaultKubeConfig := "/etc/origin/master/admin.kubeconfig"
	var err = os.ErrNotExist
	var ok bool
	if ok, err = file.FileOrSymlinkExists(defaultKubeConfig); ok {
                fmt.Printf("SERKAN -- locateKubeConfig  defaultKubeConfig  => %s. \n", defaultKubeConfig)
		return defaultKubeConfig, nil
	}

	if k := os.Getenv("KUBECONFIG"); k != "" {
		if ok, err = file.FileOrSymlinkExists(k); ok {
			return k, nil
		}
	}

	if home := homeDir(); home != "" {
		kubeconfig := filepath.Join(home, ".kube", "config")
		if ok, err = file.FileOrSymlinkExists(kubeconfig); ok {
			return kubeconfig, nil
		}
	}

	return "", err
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
