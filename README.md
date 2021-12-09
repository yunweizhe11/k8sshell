# k8sshell
为了更好的展示我们所需要的信息
## Prepare

1. 执行 build.sh
  + 拷贝 编译后的二进制到 kubectl 所在目录。

## examples
```bash
kevindeMacBook-Air:src kevin$ kubectl info get pod -n namespace
             NAME                        READY STATUS  RESTARTS         AGE         Containers      CPU       Memory
         srv-pot-worker-75b6b-hcvz4           1/1  Running    1     3406h31m51.048323s      1        50m/1      1Gi/2Gi
         srv-pollution-fker6b-rw9gl           1/1  Running    1     2930h48m36.048334s      1        50m/1      1Gi/2Gi


kevindeMacBook-Air:src kevin$ kubectl info get pod -n namespaces -o json
Method `Json` will no longer supported. You can use the `JSON` method instead of `Json` method. This method will be removed in version gotable 5.0.
[
       {
              "AGE": "3406h34m7.638246s",
              "CPU": "50m/1",
              "Containers": "1",
              "Memory": "1Gi/2Gi",
              "NAME": "srv-potion-f5b6b-hcvz4",
              "READY": "1/1",
              "RESTARTS": "1",
              "STATUS": "Running"
       }
]
```
## Usage 

```bash
$ kubectl info -h 

展示pod 资源限制明细以及pod内容器数量
Usage:
  kubectl-info [flags]
  kubectl-info [command]
Available Commands:
  completion  generate the autocompletion script for the specified shell
  get         get kubernetes resources
  help        Help about any command
Flags:
  -h, --help                help for kubectl-info
  -n, --namespace string    namespace
  -o, --outputType string   output json  default table
  -v, --version             kubernetes version
Use "kubectl-info [command] --help" for more information about a command.
```

## Todo

**已实现**
+ [x] 通过GET pod 展示 资源limit
**后续计划**
+ [ ] 增加 GET service ingress 之间关联关系