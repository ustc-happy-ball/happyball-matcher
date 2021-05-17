package dgs

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"strconv"
)

var DgsAddr []*Address
var done bool

type Address struct {
	InternalIP string
	InternalPort string
	ExternalIP string
	ExternalPort string
}

// fakeAddr return usable address for testing
func fakeAddr() {
	var addrs []*Address
	addr1 := Address{
		InternalIP:   "172.16.0.35",
		InternalPort: "32003",
		ExternalIP:   "1.15.79.161",
		ExternalPort: "32001",
	}

	addr2 := Address{
		InternalIP:   "172.16.0.6",
		InternalPort: "32003",
		ExternalIP:   "1.15.135.248",
		ExternalPort: "32001",
	}

	addr3 := Address{
		InternalIP:   "172.16.0.73",
		InternalPort: "32003",
		ExternalIP:   "1.15.221.112",
		ExternalPort: "32001",
	}

	addrs = append(addrs, &addr1,&addr2,&addr3)
	DgsAddr = addrs
}

func InitKube(stsName string) error{
	if !done {
		fakeAddr()
		return nil
	} else {
		config, err := clientcmd.BuildConfigFromFlags("", "/Users/randy/.kube/tke/cls-iw9xj6c7-config")
		if err != nil {
			log.Fatalln(err)
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalln(err)
		}

		// get pod name
		stss, err := clientset.AppsV1().StatefulSets("default").Get(context.TODO(), stsName, v1.GetOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		name, num := stss.Name, int(*stss.Spec.Replicas)
		m := make(map[string]bool)
		for i := 0; i < num; i++ {
			ip := name + "-" + strconv.Itoa(i)
			m[ip] = true
			//addr.InternalIP = ip
			//DgsAddr = append(DgsAddr, &addr)
		}

		// access the API to list pods
		pods, _ := clientset.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
		for _, pod := range pods.Items {
			pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), pod.Name, v1.GetOptions{})
			if err != nil {
				log.Fatalln(err)
			}

			if pod.Name == os.Getenv("HOSTNAME") && m[pod.Name] == true {
				var addr Address
				addr.InternalIP = pod.Status.PodIP
				addr.InternalPort = pod.Status.HostIP

			}
			//if
			//fmt.Println(pod.Status.HostIP)
		}

		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
		_ = nodes

		return nil
	}
}


