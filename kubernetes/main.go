package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	podName        = "test-deployment-xyz" //can be got from env variable created with field ref
	namespace      = "test-ns"             //can be got from env variable created with field ref
	deploymentName = "test-deployment"     //can be got from added label
	nodeName       = "node1"               //node1 is the test node name to be used as field selector for pods
)

var (
	kubeClient *kubernetes.Clientset
	podSpec    *v1.Pod //pod details availed from api
)

func newClientWithConfig() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal().Msg("Error while getting in cluster kube config : " + err.Error())
	}
	kubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal().Msg("Error while creating kubernetes client set : " + err.Error())
	}
}

func main() {
	log.Info().Msg("######## Getting kubeConfig and creating kube client #########")
	newClientWithConfig()

	log.Info().Msg("######### Getting pod details ########")
	getPodSpec()

	log.Info().Msg("############## Getting pod details list with criteria #########")
	listPods()

	log.Info().Msg("######### deleting pod ########")
	deletePod()

	log.Info().Msg("######### Getting deployment details ########")
	//operations almost same for statefulsets

	//getting deployment details
	getDeployment()
}

func getPodSpec() {
	var err error
	podSpec, err = kubeClient.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		log.Error().Msg("Error while getting pod info : " + err.Error())
		return
	}

	//getting pod name, namespace and nodename
	log.Info().Msg("pod name : " + podSpec.Name + " namespace : " + podSpec.Namespace + " nodename : " + podSpec.Spec.NodeName)

	//getting label value for label with name "test-label" from pod spec
	labels := podSpec.GetLabels()
	if labelVal, ok := labels["test-label"]; ok {
		log.Info().Msg("label test-label has value : " + labelVal)
	}

	//getting uid for pod
	uid := podSpec.GetUID()
	log.Info().Msg("pod uid : " + string(uid))

	//getting podIP
	podIP := podSpec.Status.PodIP
	log.Info().Msg("podIP : " + podIP)

	//getting pod phase i.e., if pod is Pending, Running, Succeeded, Failed or Unknown
	podPhase := string(podSpec.Status.Phase)
	log.Info().Msg("pod phase : " + podPhase)

	//getting annotation value for annotation with name "test-annotation" from pod spec
	annotations := podSpec.GetAnnotations()
	if annotationVal, ok := annotations["test-annotation"]; ok {
		log.Info().Msg("annotation test-annotation has value : " + annotationVal)
	}
}

func listPods() {
	//getting pods in a specific namespace
	podList, err := kubeClient.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Error().Msg("Error while listing pods : " + err.Error())
		return
	}
	for _, pod := range podList.Items {
		log.Info().Msg("Pod in " + namespace + " : " + pod.Name)
	}

	//getting pods with a specific label selector
	podList, err = kubeClient.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		LabelSelector: "app=test-deployment",
		//LabelSelector: labels.Set{"app": "test-depployment"}.AsSelector().String(),
	})
	if err != nil {
		log.Error().Msg("Error while listing pods : " + err.Error())
		return
	}
	for _, pod := range podList.Items {
		log.Info().Msg("Pod with label selector app=test-deployment : " + pod.Name)
	}
	//another way
	podList, err = kubeClient.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set{"app": "test-depployment"}.AsSelector().String(),
	})
	if err != nil {
		log.Error().Msg("Error while listing pods : " + err.Error())
		return
	}
	for _, pod := range podList.Items {
		log.Info().Msg("Pod with label selector app=test-deployment : " + pod.Name)
	}

	//getting pods running on a specific worker node
	podList, err = kubeClient.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + nodeName,
	})
	if err != nil {
		log.Error().Msg("Error while listing pods : " + err.Error())
		return
	}
	for _, pod := range podList.Items {
		log.Info().Msg("Pod running on worker node " + nodeName + " : " + pod.Name)
	}
}

func deletePod() {
	//deleting a specific pod in a specific namespace
	err := kubeClient.CoreV1().Pods(namespace).Delete(context.Background(), podName, metav1.DeleteOptions{})
	if err != nil {
		log.Error().Msg("Error while deleting pod : " + err.Error())
	}

	//getting pod deletion grace period deleting pod with grace period
	deletionGracePeriod := podSpec.GetDeletionGracePeriodSeconds()
	log.Info().Msg(fmt.Sprintf("Grace period : %d", *deletionGracePeriod))

	err = kubeClient.CoreV1().Pods(namespace).Delete(context.Background(), podName, metav1.DeleteOptions{GracePeriodSeconds: deletionGracePeriod})
	if err != nil {
		log.Error().Msg("Error while deleting pod with grace period : " + err.Error())
	}
}

func getDeployment() {
	var err error
	deploymentSpec, err := kubeClient.AppsV1().Deployments(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	//statefulsetSpec, err = kubeClient.AppsV1().StatefulSets(namespace).Get(context.Background(), statefulsetName, metav1.GetOptions{})
	if err != nil {
		log.Error().Msg("Error while getting deployment details : " + err.Error())
	}
	log.Info().Msg("Deployment name : " + deploymentSpec.Name)
	//getting no. of replicas
	numReplicas := int(*deploymentSpec.Spec.Replicas)
	log.Info().Msgf("no. of replicas in deployment : %d", numReplicas)
}
