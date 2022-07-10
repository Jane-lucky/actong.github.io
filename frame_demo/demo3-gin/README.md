# day1
原生`http`包
`http.HandleFunc`函数处理
`http.ListenAndServe`服务监听
```go
func SayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./frame_demo/demo3-gin/hello.txt")
	fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	//监听
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server faild,err:%v\n", err)
		return
	}
}
```

gin框架
```go
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func main() {
	//创建默认的路由引擎
	r := gin.Default()
	r.GET("/hello", sayHello)

	//启动服务
	r.Run(":9090")
}
```

TODO
1. r.Group()