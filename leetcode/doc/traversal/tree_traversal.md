# 二叉树前中后序遍历方法模板

***

## 1、递归

```go
type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func traversal(root *TreeNode) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
	    if (node == nil) {
	    	return
        }
        // preorder
        dfs(node.left)
        // inorder
	    dfs(node.right)
        // postorder
    }
    dfs(root)
}
```

前中后序，是指根节点相对于左右子节点的记录顺序来定义的。

前序：在记录左右子节点前记录根节点。

中序：在记录左节点之后，右节点之前记录根节点。

后序：在记录左右节点之后记录根节点。

所以在递归版本中，只需要在对应的位置插入对应的代码即可。

## 2、迭代

二叉树的迭代有一个通用的流程：

1. 将当前节点的左边界子节点依次压入栈，直到左子节点为空
2. 如果栈不为空，弹出栈顶节点
3. 如果栈顶节点存在右子节点，则将其压入栈并重复第1步，否则重复第2步

按照以上流程，可以很容易的写出迭代遍历的代码模板

```go
func traversal(root *TreeNode) {
	stack, node := make([]*TreeNode, 0), root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		n := len(stack) - 1
		node = stack[n]
		stack = stack[:n]
		node = node.Right
	}
}
```

在上文中我们提到，前中后序遍历，在本质上只是记录子树根节点的时机不一样。在有了迭代遍历的模板后，我们现在只需要知道在前中后序在何处记录即可。

相对于前序和中序，它们二者的区别在于前序在节点入栈时记录，而中序在节点出栈时记录。

```go
func traversal(root *TreeNode) {
	stack, node := make([]*TreeNode, 0), root
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 前序，入栈时记录
			// fmt.Printf("%d", node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		n := len(stack) - 1
		node = stack[n]
		stack = stack[:n]
		// 中序，出栈时记录
		// fmt.Printf("%d", node.Val)
		node = node.Right
	}
}
```

后序相对于前两者来说有点不同，因为我们按照上面的开始的流程来遍历的话，那么**子树根节点总是在左子节点出栈后，右子节点入栈前出栈**。而此时右子节点及其子树均还未处理。对于后序遍历来说，根节点需要在右子节点之后记录。所以，我们需要在原有的流程上加上以下条件：

1. 新增一个`prev`指针，指向上一个被记录的节点，其作用是让我们从栈弹出`node`后，用来判断`node`的右子树遍历过没有
2. 在上述流程第2步中
   1. 节点`node`出栈后，如果`node`不存在右子节点，说明当前子树已经按照后序遍历完，记录`node`顺序，并将`prev`指向`node`，继续重复上述流程第2步
   2. 存在右子节点`right`，且`prev`指针不是指向该`right`节点，意味着该节点的右子树还未处理。则将`node`和`right`按序压入栈并重复上述流程第1步。
   3. 如果`node`存在右子节点`right`，且`prev`指针指向`right`节点，则意味着`node`的左子树和右子树已经遍历过了，则记录当前`node`，并将`prev` 指向`node`，然后继续重复上述流程第2

```go
func traversal(root *TreeNode) {
	stack, node, prev := make([]*TreeNode, 0), root, root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		n := len(stack) - 1
		node = stack[n]
		stack = stack[:n]
		if node.Right == nil || node.Right == prev {
			// 左右子节点已经遍历完成
			// fmt.Printf("%d", node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
}
```

## 3、Morris

Morris是一种二叉树遍历算法，其核心思想是利用树中的空闲指针，将遍历过程中的空间消耗降低到常数级别，算法实现如下：

```go
func traversal(root *TreeNode) {
	node := root
	for node != nil {

		if node.left != nil {
			// I、左子树不为空
			// I-a、找到mostRight
			mostRight := node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}

			if mostRight.right == nil {
				// I-b、找到了该层的mostRight，将其右指针指向node
				//		node继续往下层左子树移动
				mostRight.right = node
				node = node.left
			} else {
				// I-c、mostRight不为空，该节点之前已经被遍历过一次了
				// 		第二次来到这个节点时，将该节点的右指针指向空
				//		并使node往右走
				mostRight.right = nil
				node = node.right
			}
		} else {
			// II、左子树为空、子树往右走，即往上跳回上面找到mostRight时所处的node节点
			node = node.right
		}
	}
}
```

将该算法遍历所经过节点的顺序称为Morris序，在Morris序中，所有有左子节点的节点都会被遍历两次。并且Morris遍历总是按照这样一种顺序来进行遍历，假设当前节点为`node`，左子节点为`leftt`，右子节点为`right`，并且假设`left`和`right`均为叶子节点

```
node ---> left ---> node ---> right 
```

通过该规律，可以总结出：

- 前序：在每次访问到一个第一次访问的节点时，记录该节点
- 中序：如果一个节点只会访问一次（没有左子节点的节点），记录该节点，如果一个节点会被访问两次，则在第二次访问时记录该节点

于是可得前中序遍历的代码如下：

```go
func preorder(root *TreeNode) {
	node := root
	for node != nil {
		if node.left != nil {
			mostRight := node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				// 第一次访问node节点
				fmt.Printf("%d\t", node.val)
				mostRight.right = node
				node = node.left
			} else {
				mostRight.right = nil
				node = node.right
			}
		} else {
			// 当前处于左叶子节点或node的右子节点
			fmt.Printf("%d\t", node.val)
			node = node.right
		}
	}
}

func inorder(root *TreeNode) {
	node := root
	for node != nil {
		if node.left != nil {
			mostRight := node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = node
				node = node.left
				continue
			} 
			mostRight.right = nil
		}
		fmt.Printf("%d\t", node.val)
		node = node.right
	}
}
```

和普通的迭代方法一样，Morris实现后序遍历相对于前中序遍历来说仍然要麻烦一点。但通过下面这个图来理解就会轻松不少。

![morris_postorder](D:\workspace\profile\go\go-journey\leetcode\doc\traversal\image\morris_postorder.jpg)

以上面这棵树为例，写出它的Morris序：`1`、`2`、`4`、`2`、`5`、`1`、`3`、`6`、`3`、`7`

可以知道，对于非叶子节点：`1`、`2`、`3`来说，总会被遍历两次。来节点`2`来解释：当第一次到达节点`2`时，我们可以得到其左子节点`4`，遍历完`4`之后，第二次回到节点`2`时，根据morris序，又可以得到右子节点`5`。即`2`、`4`、`2`、`5`，但是这个顺序并无法满足后序遍历的顺序要求。

看回上图，还是以`2`节点为例，加入我们要按后序来遍历，那就需要先记录`4`，然后记录`5`再记录`2`。那有没有办法满足这个需求呢？显然是有的，在我们第二次访问`2`这个节点时，`4`已经访问完成，此时可以将`2`-->`5`看成一个单链表，想要逆序记录这两个节点，只需要翻转该单链表即可。

1. 虚箭头a：从节点`4`回到节点`2`时，逆序记录节点`2`的左子节点，即节点`4`，因为节点4没有右子节点，所以得到4
2. 然后会从节点`2`移动向右子节点`5`，节点`5`没有左子节点，继续向右子节点移动，即虚箭头b
3. 虚箭头b：从节点5回到节点1时，逆序记录节点1的左子节点，即节点2，节点2有一个右子节点5，所有逆序可以得到5、2
4. 然后从节点1移动向右子节点3，遍历左子节点6，并沿虚箭头c回到节点3
5. 虚箭头c：从节点6回到节点3时，逆序记录节点3的左子节点6，得到6
6. 然后节点3向右子节点移动到节点7，结束Morris遍历。

利用Morris遍历实现后序的做法是在每次沿虚箭头回到上级节点，逆序打印其左子节点。但最后需要注意的是，整棵树的右边界，也就是虚线④是不会被处理的，所以我们需要单独逆序一次右边界。

```go
func postorder(root *TreeNode) []int {
    node, mostRight, ans := root, (*TreeNode)(nil), make([]int, 0)
	for node != nil {
		if node.left != nil {
			mostRight = node.left
			for mostRight.right != nil && mostRight.right != node {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = node
				node = node.left
			} else {
				mostRight.right = nil
				ans = append(ans, reverseRightNode(node.left)...)
				node = node.right
			}
		} else {
			node = node.right
		}
	}
	ans = append(ans, reverseRightNode(root)...)
	return ans
}


func reverseRightNode(root *TreeNode) []int {
	if root.right == nil {
		return []int{root.val}
	}
	ans := make([]int, 0)
	node := root
	for node != nil {
		ans = append(ans, node.val)
		node = node.right
	}
    // 通过翻转切片达到和翻转链表一样的效果
	for left, right := 0, len(ans)-1; left < right; {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return ans
}
```

