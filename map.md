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

```go
var m map[int]int
fmt.Println(m[1])
```

`map`传递给函数的代价很小：在32位机器上占`4`个字节，64位机器上占`8`个字节，无论实际上存储了多少数据。通过`key`在`map`中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢100倍；所以如果你很在乎性能的话还是建议用切片来解决问题。

不要使用 `new`，永远用`make`来构造`map`,如果你错误的使用`new ()`分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址

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