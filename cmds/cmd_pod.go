package cmds

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	podCreateCmd.Flags().StringVar(&podCreateName, "name", "", "pod name")
	podCreateCmd.Flags().StringVar(&podCreateImage, "image", "", "image name")
	podGetCmd.Flags().StringVar(&podGetName, "name", "", "pod name")
}

// Pod Create 命令
var podCreateCmd = cobra.Command{
	Use:   "create",
	Short: "create a new pod",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Pod Update 命令
var podUpdateCmd = cobra.Command{
	Use:   "update",
	Short: "update a pod",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Pod Get 命令
var podGetCmd = cobra.Command{
	Use:   "pod",
	Short: "pod is used to manage kubernetes Pods",
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
		var serviceObject []map[string]string
		var table *table.Table
		ctx := context.Background()
		Listopt := metav1.ListOptions{}
		// listOption := metav1.ListOptions{}
		if podGetName != "" {
			Listopt.FieldSelector = fmt.Sprintf("metadata.name=%s", podGetName)
		}
		podList, err := k8sClient.CoreV1().Pods(namespace).List(ctx, Listopt)
		if err != nil {
			fmt.Println("Err:", err)
			return
		}
		table, err = gotable.Create("NAME", "READY", "STATUS", "RESTARTS", "AGE", "IP", "NODE", "Containers", "CPU", "Memory") //table title
		for _, pod := range podList.Items {
			// containerAllCount := len(pod.Status.ContainerStatuses)
			containerAllCount := len(pod.Spec.Containers) //获取当前pod 容器数量
			containerReadyCount := 0
			for _, cs := range pod.Status.ContainerStatuses {
				if cs.State.Running != nil {
					containerReadyCount++
				}
			}
			node_ip := pod.Status.HostIP
			pod_ip := pod.Status.PodIP
			cpu_requests := ""
			cpu_limit := ""
			memory_requests := ""
			memory_limit := ""
			Containers := 1
			restartCount := 0
			starttime := ""
			if string(pod.Status.Phase) == "Running" {
				for _, item := range pod.Status.ContainerStatuses {
					restartCount = int(item.RestartCount) + restartCount //统计容器重启次数 pod内所有容器重启之和
				}
				starttime = time.Now().Sub(pod.Status.StartTime.Time).String()
			} else {
				starttime = time.Now().Sub(pod.Status.Conditions[0].LastTransitionTime.Time).String() //不在运行中pod 获取时间字段不一致
			}
			for index, key := range pod.Spec.Containers {
				Containers = index + Containers
				cpu_requests = key.Resources.Requests.Cpu().String()
				memory_requests = key.Resources.Requests.Memory().String()
				memory_limit = key.Resources.Limits.Memory().String()
				cpu_limit = key.Resources.Limits.Cpu().String()
			}
			if err != nil {
				fmt.Println("Err:", err)
				return
			}
			serviceObject = append(serviceObject, map[string]string{"NAME": pod.Name, "READY": strconv.Itoa(containerReadyCount) + "/" + strconv.Itoa(containerAllCount), "STATUS": string(pod.Status.Phase), "RESTARTS": strconv.Itoa(restartCount),
				"Containers": strconv.Itoa(Containers), "CPU": cpu_requests + "/" + cpu_limit, "Memory": memory_requests + "/" + memory_limit, "AGE": starttime, "IP": pod_ip, "NODE": node_ip})
		}
		goPrint(serviceObject, table) //打印数据
	},
}

// Pod Delete 命令
var podDeleteCmd = cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "delete a pod",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var podCreateName string
var podCreateImage string
var podGetName string

// var namespace string
