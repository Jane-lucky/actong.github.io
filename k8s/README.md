### 容器化部署的问题
服务的横向扩展

容器宕机之后的恢复

版本发布之后更新操作（不影响业务）

监控容器状态

容器如何调度创建

数据安全行的保证



### Job
需求：

- 如何保证Podcast内进程正确结束

- 如何保证进程运行失败后重试

- 如何管理多个任务，且任务之间存在依赖关系

- 如何并发地运行任务，并管理任务的队列大小


1. 管理任务的控制器，创建一个或多个pod来指定pod数量，并监控它是否成功运行或终止。
2. 根据依赖关系，确保上一个任务运行完成之后再执行下一个任务
3. 根据pod的状态来给Job设置重置的方式及重试的次数
4. 控制任务的并行度

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: hello
  namespace: testk
spec:
  template:
    spec:
      containers:
      - name: hello
        image: busybox
        command: ["echo","hello k8s job"]
      restartPolicy: Never
  backoffLimit: 4
```

**并行job：**
需求：job运行之后，最大化并行——节点数有限制，不希望同时并行pod数过多。

job**失败**之后，会存在多个失败的pod：当pod启动时，容器失败退出，根据restartPolicy:Never，所有这个容器不会被重启，但Jod 的desired是1，所以controllerhi一直创建新的pod。 

### DaemonSet（守护进程控制器）
需求：

- 希望每个节点都运行同样一个pod
- 新节点加入是，需要及时去感知并去部署一个pod，同时初始化内容
- 存在节点退出，希望删除该pod
- pod异常，及时监控异常节点。

适用的场景：
集群存储进程、日志收集进程、需要每个节点运行的监控收集器


### Pod的配置管理

**可变配置:configmap**

连接方式：通过labels实现pod和配置文件的对应；

使用：用于挂载pod配置文件、环境变量、命令行参数等

优点：配置文件的灵活性，保证容器的可移植性。

**敏感配置:secret**

存储密码token的资源对象

1. 使用私有镜像库
  私有镜像仓库的信息可以通过Secret的方式存储在集群中,配置方式：

  - 通过pod的imagePullSecrets来指定
  - 通过serviceAccount设置，之后自动为该SA的pod注入信息

  ```yaml
  apiVersion: v1
  kind: ServiceAccount
  metadata:
    name: service1
    namespace: basic
  secrets:
  - name: service1-token-mp4qs
  imagePullSecrets:
  - name: docker-reader-secret
  ---
  apiVersion: v1
  kind: Pod
  metadata:
    name: private-tag
  spec:
    containers:
    - name: private-container
      image: image-name
    imagePullSecrets:
    - name: regcred
  ---
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: nginx-deployment
    labels:
      app: nginx
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: nginx
    template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx
          image: nginx:1.14.2
        serviceAccountName: service1
  ```

  

**身份认证**

用于解决pod集群中身份认证问题。

- pod应用访问k8s集群

  pod创建之后，会将对应的secret挂载到容器目录的/var/run/secrets/kubernetes.io/serviceaccount/下，当pod访问集群时，会采用token和ca.crt来校验服务端。

**资源配置：resources**

容器的一个资源配置管理：CPU、内存以及临时存储

一个request，一个limits。分别时对需要资源和资源临界进行的一个声明。

**安全管控	**

**前置校验：Initcontainer**

1. 较普通container先启动
2. 依次执行，而普通的container则是并发启动
3. 执行完成退出，而普通容器会一直执行。

用途：前置准备——容器的初始化，或者配置文件之类；或者说是前置校验——例如网络是否联通之类。

### k8s数据卷

**emptyDir：**空目录，生命周期与pod一致——可以在同一pod内的不同容器之间共享工作过程中产生的文件；

**HostDir：**挂载在宿主机上的指定目录

**NFS：**网络存储

