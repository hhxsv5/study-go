package main

import (
	"bytes"
	"fmt"
	"sync"
)

//公共变量多线程时同步访问
type PublicInfo struct {
	mu      sync.Mutex
	content string
}

func Update(info *PublicInfo) {
	//加排它锁：不让其他读和写
	info.mu.Lock()

	//修改数据
	info.content = "Hello Mutex"

	//解锁
	info.mu.Unlock()
}

//共享缓存区：允许多个线程读，一个线程写
type SyncedBuffer struct {
	lock   sync.RWMutex
	buffer bytes.Buffer
}

//多个线程读
func ReadBuffer(buffer *SyncedBuffer) {
	fmt.Println(buffer.buffer.String())
}

//一个线程写
func WriteBuffer(buffer *SyncedBuffer) {
	//上共享锁
	buffer.lock.RLock()

	//修改buffer
	buffer.buffer.WriteString("buffer data1\n")

	buffer.buffer.WriteString("buffer data2\n")

	//解锁
	buffer.lock.RUnlock()

}

func main() {

	//info := PublicInfo{}
	//info.mu = sync.Mutex{}
	//info.content = "c1"
	//
	//Update(&info)
	//fmt.Println(info)

	buffer := SyncedBuffer{}

	ReadBuffer(&buffer)
	WriteBuffer(&buffer)
	ReadBuffer(&buffer)
}
