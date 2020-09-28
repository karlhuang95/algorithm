# 学习笔记


## 递归代码模版 python
```python3
def recursion(level, param1,param2,...):
    # recursion terminator
    if level > MAX_LEVEL:
        process_result
        return
    # process logic in current level
    process(level,data...)
    
    # drill down
    self.recursion(level+1,p1,...)
    
    # reverse the current level status if needed
```
## 递归代码模版 go
```go
func recursion(level,params1,param2,...) {
        // recursion terminator
        if level > MAX_LEVEL:
            process_result
            return
        // process logic in current level
        process(level,data...)
        
        // drill down
        recursion(level+1,p1,...)
        
        // reverse the current level status if needed
}
```


