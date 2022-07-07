# cobra
命令行程序，可以编写命令行程序

### 实现命令行程序git
仅是模拟命令行,最终也是通过os/exec库调用外部程序执行真正的git命令，返回结果

每一个cobra程序都有一个根命令，可以添加任意的子命令：
rootcmd.AddCommands(cmd)

建立在命令、参数、标志的结构之上

# go
如果程序不是单文件程序，则需要`go run .`，或者`go build -o main.exe`