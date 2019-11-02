# 散列表

## 散列函数基本要求
1. 散列函数计算得到的散列植是一个非负整数
2. 如果key1 = key2 那 hash(hey1) == hash(key2)
3. 如果key1 != key2 那 hash(key1) != hash(key2)

## 哈希算法 

- [MD5](https://zh.wikipedia.org/wiki/MD5)

- [SHA](https://zh.wikipedia.org/wiki/SHA%E5%AE%B6%E6%97%8F)
 
- [CRC](https://zh.wikipedia.org/wiki/SHA%E5%AE%B6%E6%97%8F)

## 散列冲突

1. open addressing 开放寻址法
> 如果出现了散列冲突， 我们就重新探测一个空闲位置，将其插入。
>> 如何探测新位置 ？ 
>>> linear probing 线性探测
>>>  1. 当我们往散列表中插入数据时， 如果某个数据经过散列之后， 存储位置已经被占用了， 我们就从当前位置开始， 依次往后查找， 看是否有空闲位置，直到找到为止
>>>  ![](https://static001.geekbang.org/resource/image/5c/d5/5c31a3127cbc00f0c63409bbe1fbd0d5.jpg)
>>>  2. 对于使用线性探测法解决冲突的散列表， 我们不能单纯的把要删除的元素设置为空。 而是要标记为deleted。 当遇到deleted的空间并不是停下来 而是继续探测
>>>  ![](https://static001.geekbang.org/resource/image/fe/1d/fe7482ba09670cbe05a9dfe4dd49bd1d.jpg)
>>>  3. 线性探测法最坏的情况下的时间复杂度为o(n)
>>> Quadratic probing (二次探测)
>>>  1. 跟线性探测相似， 探测的步长是原来的二次方
>>> double hashing （双重散列）
>>>  1. 使用一组散列函数。 当第一个散列函数 计算得到的位置被占用， 再用第二个散列函数
>> 为了尽可能保证散列表的操作效率， 我们用装载因子 load factor 来表示空位的大小
    1. 装载因子=元素个数/散列表的长度  装载因子越大说明空闲越少， 冲突越多， 散列表的性能会下降

> open addressing 优点： 
> 1. 不需要链表，直接存储在数组中，可以有效利用CPU 缓存加快查询
> 2. 序列化比较简单（链表法包含指针， 序列化没那么容易）
> 缺点:
> 1. 删除数据的时候比较麻烦，需要特殊标记(deleted)
> 2. 所有的数据在一个数组中，冲突的代价更高。 所以装载因子上限不能太大。所以比链表更浪费内存
> 3. 当数据量比较小， 装载因子小的时候， 适合open addressing
2. chaining 链表法
> 在散列表中， 每个bucket或者slot会对应一条链表， 所有散列植相同的元素我们都放到相同bucket对应的链表中。 当插入的时候， 我们只需要通过散列汉书计算出对应的bucket， 将其插入到对应链表中即可。 插入的时间复杂度为o(1)
> ![](https://static001.geekbang.org/resource/image/a4/7f/a4b77d593e4cb76acb2b0689294ec17f.jpg)
> 散列表优点:
> 1. 对内存的利用率更高
> 2. 对装载因子的容忍度更高
> 3. 可以适当改造， 将链表改成跳表，红黑树
> 散列表适合存储大对象，大数据量的散列表。 更加灵活。

## 如何设计散列函数

1. 散列函数设计不能太复杂
2. 散列函数生成的值要尽可能随机并且均匀分布

## 装载因子过大

1. 当装载因子过大时， 我们可以进行动态扩容， 重新申请一个更大的散列表， 将数据搬移到这个散列表中
2. 有时候不需要一次性将老数据全部搬移。 随着新数据的插入将一部分老数据插入。 查询的时候先从新散列表中查
![](https://static001.geekbang.org/resource/image/6d/cb/6d6736f986ec4b75dabc5472965fb9cb.jpg)

## 优化的LRU
![](https://static001.geekbang.org/resource/image/ea/6e/eaefd5f4028cc7d4cfbb56b24ce8ae6e.jpg)
