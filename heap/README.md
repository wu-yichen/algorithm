# 堆

1. 堆是一个完全二叉树
2. 每个节点的值都小于等于 （**小顶堆**）或者大于等于（**大顶堆**）子节点的值

## 堆化 heapify

1. 往堆中插入元素
> 把节点插入到最后，从下往上堆化
![](https://static001.geekbang.org/resource/image/e3/0e/e3744661e038e4ae570316bc862b2c0e.jpg)
```
public class Heap {
  private int[] a; // 数组，从下标1开始存储数据
  private int n;  // 堆可以存储的最大数据个数
  private int count; // 堆中已经存储的数据个数

  public Heap(int capacity) {
    a = new int[capacity + 1];
    n = capacity;
    count = 0;
  }

  public void insert(int data) {
    if (count >= n) return; // 堆满了
    ++count;
    a[count] = data;
    int i = count;
    while (i/2 > 0 && a[i] > a[i/2]) { // 自下往上堆化
      swap(a, i, i/2); // swap()函数作用：交换下标为i和i/2的两个元素
      i = i/2;
    }
  }
 }
```
2. 删除堆顶元素
> 把最后一个元素放到堆顶，然后从上往下堆化
![](https://static001.geekbang.org/resource/image/11/60/110d6f442e718f86d2a1d16095513260.jpg)
```
public void removeMax() {
  if (count == 0) return -1; // 堆中没有数据
  a[1] = a[count];
  --count;
  heapify(a, count, 1);
}

private void heapify(int[] a, int n, int i) { // 自上往下堆化
  while (true) {
    int maxPos = i;
    if (i*2 <= n && a[i] < a[i*2]) maxPos = i*2;
    if (i*2+1 <= n && a[maxPos] < a[i*2+1]) maxPos = i*2+1;
    if (maxPos == i) break;
    swap(a, i, maxPos);
    i = maxPos;
  }
}
```

###  插入和删除的时间复杂度是 o(logn)

## 堆排序
1. 建堆 - 时间复杂度是 o(n)
2. 排序 - 时间复杂度是 o(nlogn) 原地不稳定排序算法

* 如果堆顶是从下标 1 开始 - 左子节点就是 2*i 右子节点是2*i+1 父节点就是i/2
* 如果堆顶是从下标 0 开始 - 左子节点就是 2*i+1 右子节点是 2*i+2 父节点就是（i-1）/2

## 堆应用

1. 优先级队列
2. 利用堆求TOP k
> 查找前K大元素， 维护一个大小为K的小顶堆，一次与堆顶比较 如果比堆顶大，就删除堆顶然后插入这个元素
3. 利用堆求中位数  
> 维护两个堆，一个大顶堆，一个小顶堆，大顶堆的堆顶就是中位数


