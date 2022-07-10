面向对象，关注的是对象。
为了描述某个事物在整个解决问题的步骤中的行为
面向过程，关注的是过程。
分析解决问题的步骤，然后采用函数实现每一个步骤

## 优势
- 容易起名字
- 代码管理方便
- 代码冗余量较小

## go面向对象
没有Java强调的继承和多态。
没有class类——**但提供了结构体，方法可以在结构体上添加。**

## 结构体
单一的数据类型无法满足现实开发需求
### 结构体的定义和初始化
定义只是一种内存布局的描述，仅有被实例化的时候才会真正分配内存。

使用内置函数new()对结构体实例化的时候，形成的是指针类型的结构体。


### 结构体是值传递类型
如果赋值一份传递到函数中，如果在函数中改变参数，不会影响到实际的结构体

### 深拷贝和浅拷贝【demo3】
值类型是深拷贝，引用类型是浅拷贝

### 结构体作为函数的参数及返回值【demo4】
### 匿名结构体和匿名字段【demo5】

创建匿名对象的时候，同时创建对象。

### 结构体的嵌套【demo6】
将一个结构体作为另一个结构体字段的属性
**可以模拟面向对象编程中的以下两种关系
- 聚合关系：一个类作为另一个类的属性
- 继承关系：一个类作为来不过一个类的子类。

## 接口【demo7】
接口十一组方法的签名，接口指定了类型应该具有什么方法，类型决定了如何实现这些方法