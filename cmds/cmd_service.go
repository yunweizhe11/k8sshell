package cmds

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var serviceName string

func init() {
	serviceGetCmd.Flags().StringVar(&serviceName, "name", "", "service name") //添加子命令
}

// Service Create 命令
var serviceCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Service Update 命令
var serviceUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Service Get 命令
var serviceGetCmd = cobra.Command{
	Use:     "service",
	Aliases: []string{"svc"},
	Short:   "service is used to manage kubernetes Services",
	Run: func(cmd *cobra.Command, args []string) {
		if namespace == "" {
			cmd.Help()
			return
		}
		k8sClient, err := createK8SClient()
		if err != nil {
			fmt.Println("Err:", err)
			return
		}
		ctx := context.Background()
		Listopt := metav1.ListOptions{}
		if serviceName != "" {
			Listopt.FieldSelector = fmt.Sprintf("metadata.name=%s", serviceName)
		}
		serviceList, err := k8sClient.CoreV1().Services(namespace).List(ctx, Listopt)
		if err != nil {
			fmt.Println("Err:", err)
			return
		}
		var serviceObject []map[string]string
		var table *table.Table
		for _, service := range serviceList.Items {
			// 格式化 ServicePort
			servicePorts := make([]string, 0, len(service.Spec.Ports))
			for _, p := range service.Spec.Ports {
				servicePorts = append(servicePorts, fmt.Sprintf("%d:%d/%s", p.Port, p.NodePort, p.Protocol))
			}

			// 格式化 External IPs
			externalIPs := make([]string, 0, len(service.Spec.ExternalIPs))
			for _, ip := range service.Spec.ExternalIPs {
				externalIPs = append(externalIPs, ip)
			}
			var externalIPsStr = "<none>"
			if len(externalIPs) > 0 {
				externalIPsStr = strings.Join(externalIPs, ",")
			}
			ingressHostList := GetIngress(k8sClient, namespace, service.Name)                                     //获取ingress 域名
			table, err = gotable.Create("NAME", "TYPE", "CLUSTER-IP", "EXTERNAL-IP", "PORT(S)", "Ingress", "AGE") //table title
			if err != nil {
				fmt.Println("Err:", err)
				return
			}
			serviceObject = append(serviceObject, map[string]string{"NAME": service.Name, "TYPE": string(service.Spec.Type), "CLUSTER-IP": service.Spec.ClusterIP, "EXTERNAL-IP": externalIPsStr,
				"PORT(S)": strings.Join(servicePorts, ","), "Ingress": strings.Join(ingressHostList, ","), "AGE": time.Now().Sub(service.GetCreationTimestamp().Time).String()})
		}
		goPrint(serviceObject, table) //打印数据
	},
}

func GetIngress(clientset *kubernetes.Clientset, namespace string, name string) []string {
	ctx := context.Background()
	Listopt := metav1.ListOptions{}
	// Listopt.FieldSelector = fmt.Sprintf("metadata.name=%s", podGetName)
	IngressInfo, err := clientset.NetworkingV1().Ingresses(namespace).List(ctx, Listopt)
	if err != nil {
		fmt.Println("Err:", err)
	}
	var ingressHostList []string
	for _, item := range IngressInfo.Items {
		for _, key := range item.Spec.Rules {
			for _, path := range key.HTTP.Paths {
				if path.Backend.Service.Name == name {
					ingressHostList = append(ingressHostList, key.Host) //匹配ingress域名与service name 一致
				}
			}
			// key.HTTP.Paths[0].Backend.Service.Name
			// ingressHostList = append(ingressHostList, key.Host)
		}
	}
	return ingressHostList
}

// Service Delete 命令
var serviceDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
