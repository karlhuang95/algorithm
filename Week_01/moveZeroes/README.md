#  移动零

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。


**示例**
```
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
```
**说明**
```
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
```

**思路**
我们首先要找到有0的索引，我们利用切片去做处理，将0追加到最后


发现python这种方法也比较简洁，就写下来

```python
def moveZeroes(self, nums):
    zero = 0  # records the position of "0"
    for i in xrange(len(nums)):
        if nums[i] != 0:
            nums[i], nums[zero] = nums[zero], nums[i]
            zero += 1
```


https://leetcode-cn.com/problems/move-zeroes/submissions/