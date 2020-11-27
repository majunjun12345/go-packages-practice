### 源码分析
[深度解密Go语言之sync.map](https://mp.weixin.qq.com/s/tYlTE-0cs0s5DWMyVwZPYQ)

```go
type Map struct {
	mu Mutex

	read atomic.Value

	dirty map[interface{}]*entry

	misses int
}
```
- read
  是 atomic.Value 类型，可以并发地读，实际存储的是一个结构体


- dirty
  非线程安全的原始 map;
  包含新写入的 key，并且包含 read 中的所有未被删除的 key。