# binary search 二分查找

## 二分查找的时间复杂度是o(logn)
```
public int bsearch(int[] a, int n, int value) {
  int low = 0;
  int high = n - 1;

  while (low <= high) {
    int mid = (low + high) / 2;
    if (a[mid] == value) {
      return mid;
    } else if (a[mid] < value) {
      low = mid + 1;
    } else {
      high = mid - 1;
    }
  }

  return -1;
}
```
* 容易出错的三个地方
    1. 循环退出条件 low <= high 而不是 low < high
    2. mid 的取值 （low+high）/2 有可能溢出 应该是 mid= low+(high-low)/2 更进一步 low + ((high-low)>>1)
    3. low和high的更新 low=mid+1 high=mid-1

## 用递归实现
```
// 二分查找的递归实现
public int bsearch(int[] a, int n, int val) {
  return bsearchInternally(a, 0, n - 1, val);
}

private int bsearchInternally(int[] a, int low, int high, int value) {
  if (low > high) return -1;

  int mid =  low + ((high - low) >> 1);
  if (a[mid] == value) {
    return mid;
  } else if (a[mid] < value) {
    return bsearchInternally(a, mid+1, high, value);
  } else {
    return bsearchInternally(a, low, mid-1, value);
  }
}
```

## 二分查找的局限性
* 二分查找依赖的是顺序表结构，简单说就是数组
    - 因为二分查找需要按照下标随机访问，所以不可以是链表，因为链表的随机访问时间复杂度是o(n)
* 其次二分查找针对的是有序数据
    - 二分查找只能用在插入删除操作不频繁，一次排序多次查找的场景中
* 数据量太小不适合二分查找
* 数据量太大不适合二分查找

![](https://static001.geekbang.org/resource/image/42/36/4221d02a2e88e9053085920f13f9ce36.jpg)

* 查找第一个值等于给定值的元素
```
public int bsearch(int[] a, int n, int value) {
  int low = 0;
  int high = n - 1;
  while (low <= high) {
    int mid =  low + ((high - low) >> 1);
    if (a[mid] > value) {
      high = mid - 1;
    } else if (a[mid] < value) {
      low = mid + 1;
    } else {
      if ((mid == 0) || (a[mid - 1] != value)) return mid;
      else high = mid - 1;
    }
  }
  return -1;
}
```
* 查找最后一个值等于给定值的元素
```
public int bsearch(int[] a, int n, int value) {
  int low = 0;
  int high = n - 1;
  while (low <= high) {
    int mid =  low + ((high - low) >> 1);
    if (a[mid] > value) {
      high = mid - 1;
    } else if (a[mid] < value) {
      low = mid + 1;
    } else {
      if ((mid == n - 1) || (a[mid + 1] != value)) return mid;
      else low = mid + 1;
    }
  }
  return -1;
}
```
* 查找第一个大于等于给定值的元素
```
public int bsearch(int[] a, int n, int value) {
  int low = 0;
  int high = n - 1;
  while (low <= high) {
    int mid =  low + ((high - low) >> 1);
    if (a[mid] >= value) {
      if ((mid == 0) || (a[mid - 1] < value)) return mid;
      else high = mid - 1;
    } else {
      low = mid + 1;
    }
  }
  return -1;
}
```
* 查找最后一个小于等于给定值的元素
```
public int bsearch7(int[] a, int n, int value) {
  int low = 0;
  int high = n - 1;
  while (low <= high) {
    int mid =  low + ((high - low) >> 1);
    if (a[mid] > value) {
      high = mid - 1;
    } else {
      if ((mid == n - 1) || (a[mid + 1] > value)) return mid;
      else low = mid + 1;
    }
  }
  return -1;
}
```