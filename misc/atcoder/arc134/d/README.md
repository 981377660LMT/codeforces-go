设输入的数组为 $A$。对于**字典序最小子序列问题**，通常用单调栈解决：从左往右遍历 $A$，把 $A[i]$ 入栈，入栈前，把栈中 $\ge A[i]$ 的数都弹出。这样最后可以得到一个字典序最小且递增的序列。

求出 $A$ 的前 $n$ 个数的最小字典序，记作数组 $a$。此时后 $n$ 个数中的剩余元素，记作数组 $b$。

那么当前剩余元素组成的序列就是 $a+b$。接下来，看看能否得到比 $a+b$ 更小的字典序。

下面的讨论，下标从 $0$ 开始。

设 $\textit{mn}$ 为满足 $a[i]=a[0]$ 的 $b[i]$ 的最小值，分类讨论：

- 如果 $\textit{mn}\le a[0]$，例如 $a=[2,2,3,3],\ b=[3,1,2,4]$，我们只需要保留 $a[1]$ 和 $b[1]$ 即可，组成最小字典序 $[2,1]$。
- 如果 $\textit{mn}> a[0]$，继续分类讨论：
    - 例如 $a=[2,2,2,3,3,4,5],\ b=[3,1,\texttt{x},\texttt{x},\texttt{x},\texttt{x}]$，无论 $\texttt{x}$ 是什么数，我们都可以把 $a$ 中 $3$ 以及后面的数字去掉，从而得到更小的字典序。
    - 例如 $a=[2,2,2,3,3,4,5],\ b=[3,5,\texttt{x},\texttt{x},\texttt{x},\texttt{x}]$，无论 $\texttt{x}$ 是什么数，我们都可以把 $a$ 中 $4$ 以及后面的数字去掉，从而得到更小的字典序。注意不能去掉 $a$ 中 $3$ 以及后面的数字，那样会得到更大的字典序。
    - 为避免复杂的分类讨论，我们可以直接算出 $a$ 中去掉 $\ge b[0]$ 和 $> b[0]$ 后剩余的序列，哪个字典序更小，就是答案。

```py
from bisect import bisect_left, bisect_right

n = int(input())
A = list(map(int, input().split()))

a, b = [], []
for x, y in zip(A[:n], A[n:]):
    while a and x < a[-1]:
        a.pop()
        b.pop()
    a.append(x)
    b.append(y)

i = bisect_right(a, a[0])
mn = min(b[:i])
if mn <= a[0]:
    print(a[0], mn)
    exit()

l = bisect_left(a, b[0])
r = bisect_right(a, b[0])
print(*min(a[:l] + b[:l], a[:r] + b[:r]))
```

时间复杂度：$\mathcal{O}(n)$。
