链表一般操作时，一种常用的技巧是添加一个哑节点（dummy node），它的next 指针指向链表的头节点。这样一来，我们就不需要对头节点进行特殊的判断了。


牛客 5 6 7
# [206_reverseList 反转链表](https://leetcode.cn/problems/UHnkqh/)

将链表进行反转，也就是遍历已经提供链表，对其进行头插法，实现链表的反转
**头插法**
输入的数据次序生成的链表节点次序相反

**尾插法**
输入的数据次序生成的链表节点次序相同

# [92_reverseBetween 区间反转链表](https://leetcode.cn/problems/reverse-linked-list-ii/)
获取区间的前的节点以及区间的最后一个节点；
截断区间进行区间节点进行反转（不用进行返回，底层链表已经反转）
```go
pre.Next = nil
rightNode.Next = nil
reverse(leftNode)
```

# [25_reverseKGroup k组区间反转](https://leetcode.cn/problems/reverse-nodes-in-k-group/)
**递归**——程序调用自身,也就是函数自己调用自己。递归通常从顶部将问题分解，通过解决掉所有分解出来的小问题，来解决整个问题；
**迭代**——利用变量的原值推算出变量的一个新值。递归中一定有迭代,但是迭代中不一定有递归,大部分可以相互转换；
**动态规划**——通常与递归相反，其从底部开始解决问题。将所有小问题解决掉，进而解决的整个问题。

1. 递归

k组区间的遍历，通过递归的方式

- 将待反转的k个节点（如果是小于k，则不需要反转，直接返回头结点）
- 对k个结点进行反转，并返回头节点（左闭右开，所以本轮操作的尾结点其实就是下一轮操作的头结点）
- 对下一轮k节点进行反转
- 将上一节点的尾节点指向下一轮反转后的头节点，即将每一轮翻转的k节点连接起来

2. 模拟

  循环遍历，以k为单位

- 按照k个一个分组
- k组进行反转
- 采用两个临时指针节点（pre tail）
- 将子链表的头部和上一个的尾部连接，实现链表的反转

# [2413_getKthFromEnd  倒数最后k个节点](https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)
1. 顺序遍历
获取链表的总长度，然后顺序遍历链表的第n-k个节点返回
2. 双指针
将fast指针指向k+1个节点，第二个指针slow始终和fast相隔k，最终返回slow节点

# [2487_removeNthFromEnd 删除倒数第n个节点](https://leetcode.cn/problems/SLwz0R/)
1. 顺序遍历
获取链表长度，遍历len-n删倒数第n个元素
2. 栈
元素依次进栈，弹出来的第n个节点，就是需要删除的节点
3. 双指针