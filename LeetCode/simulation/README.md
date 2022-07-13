# [BM97 旋转数组](https://www.nowcoder.com/practice/e19927a8fd5d477794dac67096862042?tpId=295&tqId=1024689&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj)

数组A中存有 n 个整数，在不允许使用另外数组的前提下，将每个整数循环向右移 M（ M >=0）个位置
**不允许使用另外数组：**将右移的元素追加到原数组的后面，最后返回截取的长度的后半部分
1. 因为m可能是大于n的，因此对于数组而言，可以对n取余，每次对于n右移数组无变化

2. 也可以采用数组反转的方式求解
- 可以先将所有元素翻转，这样尾部的 m mod n 个元素就被移至数组头部，
- 然后再翻转 [0,m mod n−1] 区间的元素
- 最后翻转[m mod n,n−1] 区间的元素即能得到最后的答案。

# [54_0_spiralOrder 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/)
1. 模拟矩阵

模拟螺旋矩阵。初始位置是矩阵左上角，初始方向左，当路径超过或者进入之前访问过的位置，顺时针旋转进入下一个方向。

2. 按圈遍历（按层遍历）
- 从左到右遍历——》从上到下——》从右到左——》从下到上

- 之后四个角限制的范围分别缩减一圈

# [BM99_rotateMatrix 顺时针旋转矩阵](https://www.nowcoder.com/practice/2e95333fbdd4451395066957e24909cc?tpId=295&tqId=25283&ru=%2Fpractice%2F2e95333fbdd4451395066957e24909cc&qru=%2Fta%2Fformat-top101%2Fquestion-ranking&sourceUrl=%2Fexam%2Foj)
1. 使用辅助数组

顺时针旋转矩阵，也就意味着将数组的每一列转换成行，倒叙遍历每一列，放置在新的数组中；
O(N*2) O(N*2)

2. **[原地旋转](https://leetcode.cn/problems/rotate-matrix-lcci/)**：
TODO