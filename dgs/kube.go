package dgs

import (
	"context"
	"happyball-matcher/configs"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"reflect"
	"strconv"
	"time"
)

var GlobalDgsInfo *DgsAddress

//Address represents the info of a single dgs pod
type Address struct {
	InternalIP string
	InternalPort string
	ExternalIP string
	ExternalPort string
}

//DgsAddress represent all dgs pod address
type DgsAddress struct {
	Address []*Address
	Die 	chan struct{}
	NameSpace string
	DgsName string
	clientset *kubernetes.Clientset
}

// NewDgsAddress return an DgsAddress. It will update address info every 1 hour.
func NewDgsAddress(nameSpace string, dgsName string) *DgsAddress {
	clientSet := getClientSet()
	addrs,err := getDgsAddress(clientSet, nameSpace, dgsName)
	if err != nil {
		log.Fatalln(err)
	}

	dgs := DgsAddress{
		Address: addrs,
		Die:     make(chan struct{}),
		clientset: clientSet,
		NameSpace: nameSpace,
		DgsName: dgsName,
	}

	go dgs.maintainAddress()
	return &dgs
}

// getClientSet return clientSet instance
func getClientSet() *kubernetes.Clientset{
	// in-cluster access
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// out-of-cluster access
	//config, err := clientcmd.BuildConfigFromFlags("", "/Users/randy/.kube/tke/cls-iw9xj6c7-config")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// creates the clientset
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	return clientset
}

// getDgsAddress return all dgs pod ip address info.
func getDgsAddress(clientset *kubernetes.Clientset, nameSpace string, dgsName string) ([]*Address,error) {
	var dgsAddr []*Address
	svc,err := clientset.CoreV1().Services(nameSpace).Get(context.TODO(),dgsName,v1.GetOptions{})
	if err != nil {
		return nil,err
	}

	//var internalPort,nodePort string
	var nodePort string
	for _, port := range svc.Spec.Ports {
		//internalPort = strconv.Itoa(int(port.Port))
		nodePort = strconv.Itoa(int(port.NodePort))
	}

	set := labels.Set(svc.Spec.Selector)
	if pods,err := clientset.CoreV1().Pods("default").List(context.TODO(),v1.ListOptions{LabelSelector: set.AsSelector().String()}); err == nil {
		for _, pod := range pods.Items {
			nodePubIP := getNodePublicIP(clientset,pod.Spec.NodeName)
			podIP := pod.Status.PodIP

			dgsAddr = append(dgsAddr,&Address{
				InternalIP:   podIP,
				InternalPort: configs.DgsRPCPort,
				ExternalIP:   nodePubIP,
				ExternalPort: nodePort,
			})
		}
	} else {
		log.Fatalln(err)
	}

	return dgsAddr,nil
}

func (d *DgsAddress) updateAddress() {
	addrs,err := getDgsAddress(d.clientset,d.NameSpace,d.DgsName)
	if err != nil {
		log.Fatalln(err)
	}

	if !reflect.DeepEqual(addrs,d.Address) {
		d.Address = addrs
	}
}

// maintainAddress will update dgs ip info every hour
func (d *DgsAddress)maintainAddress() {
	tick := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <- tick.C:
			d.updateAddress()
		case <- d.Die:
			return
		}
	}
}

// PrintAddress is a help function for debugging
func (d *DgsAddress)PrintAddress() {
	for _,addr := range d.Address {
		log.Printf("InternalIP: %s, internalPort: %s, externalIP: %s, externalPort: %s\n",addr.InternalIP,addr.InternalPort,addr.ExternalIP,addr.ExternalPort)
	}
}

// getNodePublicIP return node Public IP address according to its nodeName
func getNodePublicIP(clientset *kubernetes.Clientset, nodeName string) string {
	var nodeIP string
	node,err  := clientset.CoreV1().Nodes().Get(context.TODO(),nodeName,v1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, addr := range node.Status.Addresses {
		if addr.Type == "ExternalIP" {
			nodeIP = addr.Address
			break
		}
	}

	return nodeIP
}


// fakeAddr return usable address for testing
//func fakeAddr() {
//	var addrs []*Address
//	addr1 := Address{
//		InternalIP:   "172.16.0.35",
//		InternalPort: "32003",
//		ExternalIP:   "1.15.79.161",
//		ExternalPort: "32001",
//	}
//
//	addr2 := Address{
//		InternalIP:   "172.16.0.6",
//		InternalPort: "32003",
//		ExternalIP:   "1.15.135.248",
//		ExternalPort: "32001",
//	}
//
//	addr3 := Address{
//		InternalIP:   "172.16.0.73",
//		InternalPort: "32003",
//		ExternalIP:   "1.15.221.112",
//		ExternalPort: "32001",
//	}
//
//	addrs = append(addrs, &addr1,&addr2,&addr3)
//	DgsAddr = addrs
//}


