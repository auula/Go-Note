### map声明

```go
m1 := make(map[string]int)
m2 := map[string]int{}

var m3 map[string]int
m3 = make(map[string]int)
```

### 判断某个键是否存在
```go
val, ok := map[key]
```

### map遍历
```go
func main() {
	m := map[string]int{"one":1,"two":2,"three":3}
	for k, v := range m {
	    fmt.Println(k,v)	
    }
}
```
map是无序的，遍历map时的元素顺序与添加键值对的顺序无关。

### map删除键值对
```go
delete(map, key)
```

### map使用时需要注意的点

```go
var m map[int]int
m[1] = 1    //panic: assignment to entry in nil map
```
这样会引发panic，原因是map底层是一个哈希表，var只是声明m的类型，并未申请内存。


### map的实现

hmap源码：
```go
type hmap struct {
	
	count     int       // 当前哈希表的元素个数，len(map)返回的就是该字段 
	flags     uint8     // 状态标识
	B         uint8     // buckets的数量的对数，哈希表中桶的数量为 2^B个
	noverflow uint16    // 溢出的bucket个数
	hash0     uint32    // hash seed，在哈希创建时随机生成，并在计算key的哈希的时候传入哈希函数，提高哈希函数的随机性

	buckets    unsafe.Pointer   // 指向buckets数组的指针
	oldbuckets unsafe.Pointer   // 如果发生扩容，该指针指向旧的buckets数组，旧的buckets数组时新的buckets数组大小的 1/2，非扩容状态下为nil
	nevacuate  uintptr          // 扩容进度，小于此地址的buckets表示已搬迁完成

	extra *mapextra // 用于扩容的指针
}
```
bmap：
```go
type bmap struct {
	tophash [bucketCnt]uint8
}
```