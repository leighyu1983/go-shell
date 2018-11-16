package  util

import (
	"os"
    "runtime/debug"
	"time"
    "fmt"
    "github.com/gin-gonic/gin"
)


func PanicHttpHandler(c *gin.Context) {
    if err := recover(); err != nil {
        errMsg :=  fmt.Sprintf("%s", err) + "\n" + string(debug.Stack())
        fmt.Println(errMsg)
        c.JSON(200, gin.H{"code": 1, "data": "", "message": errMsg}) 
    }
}


func PanicHandler() {
    exeName := os.Args[0] //获取程序名称

    now := time.Now()  //获取当前时间
    pid := os.Getpid() //获取进程ID

    time_str := now.Format("20060102150405")                          //设定时间格式
    fname := fmt.Sprintf("%s-%d-%s-dump.log", exeName, pid, time_str) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
    fmt.Println("dump to file ", fname)

    f, err := os.Create(fname)
    if err != nil {
        return
    }
    defer f.Close()

    if err := recover(); err != nil {
        f.WriteString(fmt.Sprintf("%v\r\n", err)) //输出panic信息
        f.WriteString("========\r\n")
    }

    f.WriteString(string(debug.Stack())) //输出堆栈信息
}