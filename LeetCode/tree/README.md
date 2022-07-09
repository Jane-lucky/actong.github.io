# 堆
堆也是实现优先队列的常用方式
## 小顶堆
一颗完全二叉树，非叶子结点的值不大于左孩子和和右孩子
## 原理
堆中元素的类型需要**实现**heap.Interface这个接口
```go
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
```
sort.Interface
其中Less()提供按照什么样的方式排序
```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```
提供了container/heap来实现堆的操作

实现：

```go
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimim:%d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
```

初始化Init()
分别构建左子树和右子树
```go
for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
    func down(h Interface, i0, n int) bool {
	i := i0
	for {
        //循环遍历生成：j1设置为左子树
		j1 := 2*i + 1
        //生产条件：在数组中
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
```

## 面试题
1. 如何将一个存有10亿个数字的文档中获取最大的10个数，计算机内存只有1M
- 让计算机io取读取文件
- 将读取的文件构成一个仅包含10个元素的小顶堆
- 构建完成后，每一次从文件读取出来的一个数组和堆顶比较
如果比堆顶小，跳过，如果比堆顶大，就可以替代，调整小顶堆
# 败者树

## 面试题
1. 对5亿个数进行排序，内容中只能容纳5千万数据：可采用外部排序、多路归并排序
- 对内容进行切片，以上述为例，可分为10段
- 依次放入内容进行排序，再写回硬盘
- 取每一段的头，取最小值，写入最终文件，然后补充数据（时间复杂度为O(n)）
**采取败者树**，也就是归并