# 队列

* 数组实现的队列为 顺序队列
* 链表实现的队列为 链式队列
```
// 用数组实现的队列
public class ArrayQueue {
  // 数组：items，数组大小：n
  private String[] items;
  private int n = 0;
  // head表示队头下标，tail表示队尾下标
  private int head = 0;
  private int tail = 0;

  // 申请一个大小为capacity的数组
  public ArrayQueue(int capacity) {
    items = new String[capacity];
    n = capacity;
  }

   // 入队操作，将item放入队尾
  public boolean enqueue(String item) {
    // tail == n表示队列末尾没有空间了
    if (tail == n) {
      // tail ==n && head==0，表示整个队列都占满了
      if (head == 0) return false;
      // 数据搬移
      for (int i = head; i < tail; ++i) {
        items[i-head] = items[i];
      }
      // 搬移完之后重新更新head和tail
      tail -= head;
      head = 0;
    }
    
    items[tail] = item;
    ++tail;
    return true;
  }

  // 出队
  public String dequeue() {
    // 如果head == tail 表示队列为空
    if (head == tail) return null;
    // 为了让其他语言的同学看的更加明确，把--操作放到单独一行来写了
    String ret = items[head];
    ++head;
    return ret;
  }
}
```

## 阻塞队列
![](https://static001.geekbang.org/resource/image/5e/eb/5ef3326181907dea0964f612890185eb.jpg)
![](https://static001.geekbang.org/resource/image/9f/67/9f539cc0f1edc20e7fa6559193898067.jpg)
## 并发队列
> 线程安全的队列 叫做并发队列。 最简单的方法就是在enqueue dequeue上面加锁

基于链表的实现方式，可以实现一个支持无限排队的无界队列(unbounded queue) 但是可能会导致过多的请求排队等待。针对响应时间比较敏感的系统，基于链表实现的无限排队的线程池是不合适的
基于数组实现的有界队列，当排队的请求超过队列大小时，接下来的请求就会被拒绝。这种方式对相应时间敏感的系统来说，就更加合理
