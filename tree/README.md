# 树

![](https://static001.geekbang.org/resource/image/40/1e/4094a733986073fedb6b9d03f877d71e.jpg)

## 二叉树

* 每个节点最多有两个叉， 也就是两个子节点， 分别是左子节点和右子节点。 二叉树并不要求每个节点都有两个子节点。
* **Full Binary Tree** 编号为2的二叉树，叶子节点全都在最底层，除了叶子节点之外，每个节点都有左右两个子节点，这种二叉树叫作满二叉树
* **complete binary tree** 编号为3的二叉树中，叶子节点都在最底下两层，最后一层的叶子节点都靠左排列，并且除了最后一层，其他层的节点个数都要达到最大
![](https://static001.geekbang.org/resource/image/09/2b/09c2972d56eb0cf67e727deda0e9412b.jpg)
**complete binary tree**
![](https://static001.geekbang.org/resource/image/18/60/18413c6597c2850b75367393b401ad60.jpg)

### 二叉树存储

1. 链式存储 - 基于指针或者引用
![](https://static001.geekbang.org/resource/image/12/8e/12cd11b2432ed7c4dfc9a2053cb70b8e.jpg)
2. 顺序存储法 - 基于数组
> 我们把根结点存储在下标i=1的位置，左子节点存储在下标2*i=2的位置， 右子节点存储在2*i+1=3的位置
![](https://static001.geekbang.org/resource/image/14/30/14eaa820cb89a17a7303e8847a412330.jpg)
* 如果是完全二叉树，用数组存储最节省内存
* 堆就是一种完全二叉树，存储方式就是数组

### 二叉树的遍历
* 前序遍历
> middle-left-right
> 递推公式 **preOrder(r)= print r-> preOrder(r ->left) -> preOrder(r->right)**
* 中序遍历
> left-middle-right
> 递推公式 **inOrder(r)= inOrder(r ->left) -> print r -> inOrder(r->right)**
* 后序遍历
> left-right-middle
> 递推公式 **postOrder(r)= postOrder(r ->left) -> postOrder(r->right) -> print r**

![](https://static001.geekbang.org/resource/image/ab/16/ab103822e75b5b15c615b68560cb2416.jpg)
---
```
void preOrder(Node* root) {
  if (root == null) return;
  print root // 此处为伪代码，表示打印root节点
  preOrder(root->left);
  preOrder(root->right);
}

void inOrder(Node* root) {
  if (root == null) return;
  inOrder(root->left);
  print root // 此处为伪代码，表示打印root节点
  inOrder(root->right);
}

void postOrder(Node* root) {
  if (root == null) return;
  postOrder(root->left);
  postOrder(root->right);
  print root // 此处为伪代码，表示打印root节点
}
```

#### 二叉树遍历的时间复杂度为 o(n)

## binary search tree 二叉查找树 -  支持动态数据集合的快速插入，删除，查找操作
> 在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值
> ![](https://static001.geekbang.org/resource/image/f3/ae/f3bb11b6d4a18f95aa19e11f22b99bae.jpg)
### 二叉查找树查找
> 1. 如果根节点等于查找的值就返回
> 2. 如果小于就在左子树中递归查找
> 3. 如果要是比根节点大就在右子树中递归查找
```
public class BinarySearchTree {
  private Node tree;

  public Node find(int data) {
    Node p = tree;
    while (p != null) {
      if (data < p.data) p = p.left;
      else if (data > p.data) p = p.right;
      else return p;
    }
    return null;
  }

  public static class Node {
    private int data;
    private Node left;
    private Node right;

    public Node(int data) {
      this.data = data;
    }
  }
}
```
### 二叉查找树插入
> 1. 如果插入的数据比节点的数据大，并且节点的右子树为空，就将数据直接插到右子节点的位置
> 2. 如果不为空，就再递归遍历右子树，查找插入位置
> 3. 如果插入的数据比节点数值小，并且左子树为空，就插入到左子节点
> 4. 如果不为空，就再递归左子树，查找插入位置
```
public void insert(int data) {
  if (tree == null) {
    tree = new Node(data);
    return;
  }

  Node p = tree;
  while (p != null) {
    if (data > p.data) {
      if (p.right == null) {
        p.right = new Node(data);
        return;
      }
      p = p.right;
    } else { // data < p.data
      if (p.left == null) {
        p.left = new Node(data);
        return;
      }
      p = p.left;
    }
  }
}
```
### 二叉查找树删除
> 1. 如果要删除的节点没有子节点，我们只需要直接将父节点中，指向要删除节点的指针设置为null
> 2. 如果要删除的节点只有一个子节点，我们只需要更新父节点中，指向要删除节点的子节点就可以了
> 3. 如果要删除的节点有两个子节点，我们需要找到这个节点的右子树中的最小节点，把它替换到要删除的节点上
> 4. 还可以直接标记为删除，并不真正从树中将这个节点去掉

#### 中序遍历二叉查找树，可以输出有序的数据序列 时间复杂度o(n)
#### 二叉查找树的时间复杂度
1. 最坏的情况：根结点的左右子树极度不平衡，退化成链表，查找的时间复杂度变为o(n)
2. 最理想情况: 二叉查找树是一棵完全二叉树（或满二叉树）o(height)
3. 平衡二叉查找树 插入删除查找时间复杂度o(logn)

#### 散列表和二叉查找树比较
* 如果输出有序的数据： 散列表中的数据是无序存储的，对于二叉查找树只需要中序遍历
* 散列表扩容耗时很多，遇到散列冲突性能不稳定

## 平衡二叉树
> 二叉树中任意一个节点的左右子树的高度相差不能大于1。 完全二叉树，满二叉树都是平衡二叉树

> A balanced binary search tree is a tree that automatically keeps its height small (guaranteed to be logarithmic) for a sequence of insertions and deletions.
![](https://static001.geekbang.org/resource/image/dd/9b/dd9f5a4525f5029a8339c89ad1c8159b.jpg)

* 最先被发明的平衡二叉查找树是 [AVL树](https://zh.wikipedia.org/wiki/AVL%E6%A0%91)
### 发明平衡二叉查找树的初衷是解决普通二叉查找树在频繁的插入删除等动态更新的情况下，出现的时间复杂度退化的问题
> 所以，如果我们现在设计一个新的平衡二叉查找树，只要树的高度不比 log2n 大很多就可以

### red-black tree (R-B Tree)
* 根节点是黑色的
* 每个叶子节点都是黑色的空节点(nil), 也就是说，叶子节点不存数据
* 任何相邻的节点都不能同时为红色，红色节点是被黑色节点隔开的
* 每个节点，从该节点到达其可达叶子节点的所有路径，都包含相同数目的黑色节点
![](https://static001.geekbang.org/resource/image/90/9a/903ee0dcb62bce2f5b47819541f9069a.jpg)

- 红黑树是一种性能非常稳定的二叉查找树，插入，删除，查找操作的时间复杂度都是o(logn)
- 但是实现难度很高，更倾向用跳表来代替
