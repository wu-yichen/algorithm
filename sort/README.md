# 排序

![](https://static001.geekbang.org/resource/image/fb/cd/fb8394a588b12ff6695cfd664afb17cd.jpg)

## 冒泡排序 bubble sort
> 冒泡排序只会操作相邻的两个数据，每次冒泡操作都会对相邻的两个元素进行比较
>![](https://static001.geekbang.org/resource/image/40/e9/4038f64f47975ab9f519e4f739e464e9.jpg)
```
// 冒泡排序，a表示数组，n表示数组大小
public void bubbleSort(int[] a, int n) {
  if (n <= 1) return;
 
 for (int i = 0; i < n; ++i) {
    // 提前退出冒泡循环的标志位
    boolean flag = false;
    for (int j = 0; j < n - i - 1; ++j) {
      if (a[j] > a[j+1]) { // 交换
        int tmp = a[j];
        a[j] = a[j+1];
        a[j+1] = tmp;
        flag = true;  // 表示有数据交换      
      }
    }
    if (!flag) break;  // 没有数据交换，提前退出
  }
}
```

* 冒泡排序是原地排序 稳定的排序算法 最好的时间复杂度是o(n) 最坏的情况是 o(n2)

## 插入排序 insertion sort
> 我们将数组中的数据分为两个区间，已排序区间和未排序区间。初始已排序区间只有一个元素，就是数组的第一个元素。插入算法的核心就是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入
>![](https://static001.geekbang.org/resource/image/b6/e1/b60f61ec487358ac037bf2b6974d2de1.jpg)
```
// 插入排序，a表示数组，n表示数组大小
public void insertionSort(int[] a, int n) {
  if (n <= 1) return;

  for (int i = 1; i < n; ++i) {
    int value = a[i];
    int j = i - 1;
    // 查找插入的位置
    for (; j >= 0; --j) {
      if (a[j] > value) {
        a[j+1] = a[j];  // 数据移动
      } else {
        break;
      }
    }
    a[j+1] = value; // 插入数据
  }
}
```

* 插入排序是原地排序 稳定的排序算法 最好的时间复杂度是o(n) 最坏的情况是 o(n2)

## 选择排序 selection sort
> 选择排序算法也分已排序区间和未排序区间。每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾
>![](https://static001.geekbang.org/resource/image/32/1d/32371475a0b08f0db9861d102474181d.jpg)

* 选择排序是原地排序 **不是稳定**的排序算法 最好的时间复杂度是o(n2) 最坏的情况是 o(n2)

![](https://static001.geekbang.org/resource/image/34/50/348604caaf0a1b1d7fee0512822f0e50.jpg)

## 归并排序 merge sort
> 先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起
>![](https://static001.geekbang.org/resource/image/db/2b/db7f892d3355ef74da9cd64aa926dc2b.jpg)
> 归并排序使用的是分治思想， 一般用递归来实现

* 归并排序是稳定的。 最好最坏平均情况时间复杂度都是o(nlogn) 但是归并排序不是原地的  空间复杂度是o(n)

## 快速排序 quick sort
> 如果要排序数组中下标从p到r之间的一组数据，我们选择p到r之间的任意一个数据作为pivot,我们遍历p到r之间的数据，将小于pivot的放到左边，将大于pivot的放到右边，将pivot放到中间
> ![](https://static001.geekbang.org/resource/image/4d/81/4d892c3a2e08a17f16097d07ea088a81.jpg)

* 快速排序是原地不稳定的排序算法。它的时间复杂度一般是o(nlogn) 但是每次分区得到的两个区间都是不均等的那么快排的复杂度就会退化成o(n2)
* 最理想的分区点是: 被分区点分开的两个分区中，数据的数量差不多


## 桶排序 bucket sort
![](https://static001.geekbang.org/resource/image/98/ae/987564607b864255f81686829503abae.jpg)
* 桶排序的时间复杂度是o(n)
* 桶排序比较适合用在外部排序中，数据存储在外部磁盘中，数据量比较大，内存有限，无法将数据全部加载到内存中

## 计数排序 counting sort
> 计数排序只能用在数据范围不大的场景中，如果数据范围k比要排序的数据n大很多，就不适合用计数排序了，而且计数排序只能给非负整数排序

![](https://static001.geekbang.org/resource/image/1f/fd/1f6ef7e0a5365d6e9d68f0ccc71755fd.jpg)



