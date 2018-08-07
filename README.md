# tasks

## 说明
golang的多任务执行，基于job worker模式

## 安装
```sh
go get github.com/anerg2046/tasks
```

## 用法
你需要自己实现一个struct，并且必须有一个Exec方法

```go
package main

import (
	"anerg/tasks"
	"fmt"
	"os"
	"sync"

	"github.com/anerg2046/snowflake"
	"github.com/gin-gonic/gin"
)

var work tasks.Job

type Test struct {
	Num string
	mu  sync.Mutex
}

func (t *Test) Exec() error {
	t.mu.Lock()
	fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(t.Num + "\n")
	fd.Write(buf)
	fd.Close()
	t.mu.Unlock()
	return nil
}

func main() {
	node, err := snowflake.NewNode(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		work = &Test{Num: node.NextID().String()}
		tasks.JobQueue <- work
	})
	router.Run(":8080")
}

```