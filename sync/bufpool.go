/*** 
Implementing a buffer pool with a self adjusting buffer.

If you'd like to fix the buffer then use `make([]byte, 1024)` instead,

***/

package main

import (
	"sync"
	"bytes"

	"fmt" // for main function
)

type BufferPool struct {
	pool sync.Pool;
}

func NewBufferPool() (*BufferPool) {
	return &BufferPool {
		pool: sync.Pool {
			New : func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

func (b *BufferPool) Get() *bytes.Buffer{
	return b.pool.Get().(*bytes.Buffer)
}

func (b *BufferPool) Put(buf *bytes.Buffer) {
	buf.Reset();
	b.pool.Put(buf);
}

func main() {
	pool := NewBufferPool();	

	buf := pool.Get();
	buf.WriteString("Hello World!");
	pool.Put(buf);

	buf.WriteString("Hello Again!");
	fmt.Println(buf);
}

