# drone插件

**制作原理：**参照变量的方式（PLUGIN_）

制作脚本文件——》封装镜像——》插件完成，编写drone文件验证正确性

## 基于bash的插件制作

1. 制作脚本文件

```shell
#! /bin/sh

curl -X ${PLUGIN_METHOD} -d ${PLUGIN_BODY} ${PLUGIN_URL}
```

2. 制作Dockerfile文件

```dockerfile
FROM alpine:3.14.6

COPY build.sh /bin/

RUN chmod +x /bin/build.sh

RUN apk -Uuv add curl ca-certificates
ENTRYPOINT /bin/build.sh
```

构建镜像：docker build -t test:v1 .

3. .drone.yml文件

```yml
kind: pipeline
type: docker
name: default

steps:
- name: greeting
  image: test:v1
  settings:
    url: http://hook.acme.com
    method: post
    body: |
      hello world

node:
  foo: bar
```

## 基于go语言的插件

1. 编写需求对应的go文件

```go
main.go
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o webhook
2. 构建镜像

参考链接：https://docs.drone.io/plugins/overview/

