package cmds

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/liushuochen/gotable/table"
	"github.com/spf13/cobra"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 以下这些值在每次实验中因为环境重建问题可能发现变化，需要自行填写
// 获取这些值的方法参考第三节我们讲解创建 ServiceAccount 名称为
// shiyanlou-admin的地方。
var (
	// K8SCertificateData 表示 Kubernetes 服务端证书
	K8SCertificateData = `-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE4MDUyMTA4NTgwOFoXDTI4MDUxODA4NTgwOFowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKi6
iiiSNO8TK3eYFsaCJF3ZlZreno3z/4cFHjY7C5Bct1ZaEUB15xO8MlBt3puEpb5o
o2fZ2+IIL1NRWyvjAsexJ4oyHk0xBP0a9KjEseypiw+m5lxe826GKLUX18BguaPX
Dge8qIh7b3/zWEfYkb7G/tLjtNKIDYDN+OOt6tjohjZ7FSP8G1qXEuj7MaqjFq4a
LB3uSkzITZ4aOPP0Yrpa9dzSjq+hHWz6H88Tg98oZL7PRIrrHOPXhMYXZwfoEtGM
dr73Abze5+2tLdN5Nv+5mtXbLxpVL+x6mBxwGEQ2bqUspuJ/SgHBrbV5ylmPzffp
vk1WfNy8k6ZDf2I+DU0CAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBABq1HL188anTft34C9HG9i7rMqmO
Uyzv1YiFeQNaO0jp7IlbzN8RO6PV0tT/SJbCh3MmwP/SatrrpNTI4snwm9aTIqRv
bSmOo63uu2yPQPosUvPtsNF8XmzXVF3vqMcdy/J/w9hcwRZVJ53K/M8noc3rwZ9d
O2k2wXa4F7Q3bIIgcQhXgHiWT2iGi1n61Rnci+PAePZNkX1X1DhNPz/6UVToxxur
i0v7L8KEvTjprdNYML/aZCHVwamvz9y6dFJ3SIGok34EGSWWd6SoJPD3dd7Now8V
oByo9DPJ5u6y45iAb7kj7NEc15P1Qq94srB6Zs5B4qzAW2uzJL0YhThpRWA=
-----END CERTIFICATE-----`

	// K8SAPIServer 表示 Kubernetes API Server 地址
	K8SAPIServer = "https://10.192.0.2:6443"

	// K8SAPIToken 表示 ServiceAccount shiyanlou-admin 的 Secret 对应的 Token
	K8SAPIToken = `eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InNoaXlhbmxvdS1hZG1pbi10b2tlbi1mcWo5dyIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJzaGl5YW5sb3UtYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI3Mjk4NmRjMy0xZTM1LTExZTktODU4My1jYTQwZWVlNDBmOTgiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6ZGVmYXVsdDpzaGl5YW5sb3UtYWRtaW4ifQ.T6SKaRlHo3G_DItrRxNbKS6UHbe3-HmxjmSkgUd6JQWhWll5Y0VUX8AFCHjATzPFSjmKNmE9va_ljemm4H06O8a0AF4NTUZoliodhghquomtuS0nVbREMRli08hUfXg_uq2PDJqls0XKj6hmtq8pZfGQ8vUJTR8yEkX-1bPel0aT_6qwf2-D1KjLm-JGGCp4XWWLP89C-9jbcFEPbHOMuEbnh5jmXm8tBXp2tcMnFDVWRNvjMU8VtwuRLX4Vr7yqNrIIwSsRlBA9N228fgbTZ81CKg5wxRmC4Emli5YpSDD_TmBj7VVEmpVU9C82Y19FdyeHCBdwoAX6VdK2vrIxZg`

	// K8SAPITimeout 表示超时时间
	K8SAPITimeout = 30
)

var namespace string
var version bool
var outputType string

// GetRootCommand 返回组装好的根命令
func GetRootCommand() *cobra.Command {
	// 定义根命令
	rootCmd := cobra.Command{
		Use: "kubectl-info",
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				restclient, err := createK8SClient()
				if err != nil {
					fmt.Println("Err:", err)
					return
				}
				// 通过 ServerVersion 方法来获取版本号
				versionInfo, err := restclient.ServerVersion()
				if err != nil {
					fmt.Println("Err:", err)
					return
				}
				fmt.Println("Kubernetes Version:", versionInfo.String())
			}
		},
	}
	// 添加全局选项参数
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "namespace")
	rootCmd.PersistentFlags().StringVarP(&outputType, "outputType", "o", "", "output json/table/yaml  default table")

	// 添加显示版本的信息
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "kubernetes version")

	// 添加子命令
	addCommands(&rootCmd)
	return &rootCmd
}

// addCommands 将各个命令拼装在一起
func addCommands(rootCmd *cobra.Command) {
	//get
	GetCmd.AddCommand(&podGetCmd)
	GetCmd.AddCommand(&serviceGetCmd)
	// GetCmd.AddCommand(&ingressGetCmd)
	// GetCmd.AddCommand(&deploymentGetCmd)
	// GetCmd.AddCommand(&secretGetCmd)
	//

	// 组装命令
	rootCmd.AddCommand(&GetCmd)
}

// get 命令
var GetCmd = cobra.Command{
	Use:   "get",
	Short: "get kubernetes resources",
}

func goPrint(serviceObject []map[string]string, table *table.Table) {
	for _, value := range serviceObject {
		err := table.AddRow(value)
		if err != nil {
			fmt.Println("Err:", err)
			return
		}
	}
	if outputType == "" {
		table.CloseBorder()
		table.PrintTable()
	}
	if outputType == "json" {
		jsonString, err := table.Json(4)
		if err != nil {
			fmt.Printf("Err", err)
		}
		fmt.Println(jsonString)
	}
	if outputType == "yaml" {
		fmt.Println("yaml is not supported")
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// createK8SClient 根据鉴权信息创建 Kubernetes 的连接客户端
func createK8SClient() (k8sClient *k8s.Clientset, err error) {
	// cfg := restclient.Config{}
	// cfg.Host = K8SAPIServer
	// cfg.CAData = []byte(K8SCertificateData)
	// cfg.BearerToken = K8SAPIToken
	// cfg.Timeout = time.Second * time.Duration(K8SAPITimeout)
	// k8sClient, err = k8s.NewForConfig(&cfg)
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	k8sClient, err = k8s.NewForConfig(config)
	return
}
