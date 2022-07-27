# 广度遍历代码模板

## 1、以二叉树为例

```go
func BFS(node *structure.Node) {
    queue := []*structure.Node{node}
    for index := 0; index < len(queue); index++ {
        n := queue[index]
        // fmt.Println(n.Val)
        if n.Left != nil{
            queue = append(queue, n.Left)
        }
        if n.Right != nil {
            queue = append(queue, n.Right)
        }
    }
}
```

## 2、图
```go
func BFSEdges(n int, edges [][]int) {
    adj := make([][]int, n)
    
    // 构造邻接矩阵
    for _, v := range edges {
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    
    queue := []int{0}
    node := 0
    visited := make([]bool, n)
    for len(queue) > 0 {
        // 取得与当前节点连通的所有节点
        node,queue = queue[0], queue[1:]
        // 若该节点已访问则跳过
        if visited[node] {
            continue
        }
        for _, childrenV := range adj[node] {
            // 将所有与该节点连用的节点加入队列
            queue = append(queue, childrenV)
        }
        // fmt.Println(node)
        visited[node] = true
    }
}
```