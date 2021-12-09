# k8sshell
更好的输出所需要的信息

## Prepare

1. 执行 build.sh
  + 拷贝 编译后的二进制到 kubectl 所在目录。

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

**后续计划**
+ [ ] 增加 GET service ingress 之间关联关系