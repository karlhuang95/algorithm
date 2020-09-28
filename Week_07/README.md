# 总结

**Trie 树代码模板**
*python3*
```python3
# Python 
class Trie(object):
  
	def __init__(self): 
		self.root = {} 
		self.end_of_word = "#" 
 
	def insert(self, word): 
		node = self.root 
		for char in word: 
			node = node.setdefault(char, {}) 
		node[self.end_of_word] = self.end_of_word 
 
	def search(self, word): 
		node = self.root 
		for char in word: 
			if char not in node: 
				return False 
			node = node[char] 
		return self.end_of_word in node 
 
	def startsWith(self, prefix): 
		node = self.root 
		for char in prefix: 
			if char not in node: 
				return False 
			node = node[char] 
		return True
```


*go*
```go
type runeTireNode struct {
    child map[rune]*runeTireNode
    Vaule interface{}
}
type RuneTire struct {
    root *runeTireNode
}

func NewRuneTire() *RuneTire  {
    return &RuneTire{root:&runeTireNode{child:make(map[rune]*runeTireNode)}}
}
func (r *RuneTire) Insert(key string,value interface{})  {
    node := r.root
    for _,c := range key{
        if n,ok :=node.child[c];!ok{
            newNode := &runeTireNode{child:make(map[rune]*runeTireNode)}
            node.child[c] = newNode
            node = newNode
        }else{
            node = n
        }
    }
    node.Vaule = value
}

func (r *RuneTire) Get(key string) interface{}  {
    node := r.root
    for _,c := range key{
        if n,ok := node.child[c];!ok{
            return nil
        }else{
            node = n
        }
    }
    return node.Vaule
}
func main() {
    r := NewRuneTire()
    r.Insert("河北","河北")
    r.Insert("湖南","湖南")
    r.Insert("湖北","湖北省")
    fmt.Println(r.Get("湖北"))
}
```
# 分析单词搜索 2 用 Tire 树方式实现的时间复杂度
```go
type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
    root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}

	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, root, &result)
		}
	}
	return result
}

func dfs(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {  //访问过了或者字典中没有
		return
	}
	node = node.children[c-'a']
    
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""  // 防止重复添加
	}

	board[i][j] = '#'
	dfs(i+1, j, board, node, result)
	dfs(i-1, j, board, node, result)
	dfs(i, j+1, board, node, result)
	dfs(i, j-1, board, node, result)
	board[i][j] = c
}
```

# 并查集代码模板
**java**
```java
// Java
class UnionFind { 
	private int count = 0; 
	private int[] parent; 
	public UnionFind(int n) { 
		count = n; 
		parent = new int[n]; 
		for (int i = 0; i < n; i++) { 
			parent[i] = i;
		}
	} 
	public int find(int p) { 
		while (p != parent[p]) { 
			parent[p] = parent[parent[p]]; 
			p = parent[p]; 
		}
		return p; 
	}
	public void union(int p, int q) { 
		int rootP = find(p); 
		int rootQ = find(q); 
		if (rootP == rootQ) return; 
		parent[rootP] = rootQ; 
		count--;
	}
}
```

**python3**
```python3
# Python 
def init(p): 
	# for i = 0 .. n: p[i] = i; 
	p = [i for i in range(n)] 
 
def union(self, p, i, j): 
	p1 = self.parent(p, i) 
	p2 = self.parent(p, j) 
	p[p1] = p2 
 
def parent(self, p, i): 
	root = i 
	while p[root] != root: 
		root = p[root] 
	while p[i] != i: # 路径压缩 ?
		x = i; i = p[i]; p[x] = root 
	return root
```

