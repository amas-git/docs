---
title: 几个基础算法 
tags:
---
<!-- toc -->
# 算法
## 杨辉三角
###  zsh
```sh
function pascalTriangle() {
    local -a xs 
    xs=(1)
    function next() {
        local -a rx
        for (( i=1; i<=$#xs; i++ )); do
            [[ -n $xs[i+1] ]] && rx+=($(($xs[i]+$xs[i+1])))
        done
        xs=(1 $rx 1)
        print $xs
    }
    print $xs
    [[ $1 == 1 || -z $1 ]] && return
    for x in {1..$((${1:=2}-1))}; do
        next
    done
}
```
```sh
$ pascalTriangle 3
1
1 1
1 2 1
$ pascalTriangle 6
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1
1 5 10 10 5 1
```

## Bubble Sort
冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
如果两个相等的元素没有相邻，那么即使通过前面的两两交换把两个相邻起来，这时候也不会交换，所以相同元素的前后顺序并没有改变，所以冒泡排序是一种稳定排序算法。
附张图，对冒泡的表现很形象

### haskell
```hs
bsort xs = bsort' xs (length xs)
  where
    bsort' xs 1 = xs
    bsort' xs n = bsort' (bubble xs) (n-1)
    
    bubble [] = []
    bubble (x:[]) = [x]
    bubble (x:y:[])
      | x < y = [x,y]
      | otherwise = [y,x]
    bubble (x:y:xs)
      | x < y = x:(bubble (y:xs))
      | otherwise = y:(bubble (x:xs))
```

### zsh
这个程序用来演示冒泡排序的整个过程
```sh
#!/bin/zsh
xs=($argv)
for (( i=$#xs; i>0; --i )); do
    for (( j=1; j<i; ++j )); do
        if (( xs[j] > xs[j+1] )); then
            x=$xs[j]
            xs[j]=$xs[j+1]
            xs[j+1]=$x
        fi
    done
    echo $xs
done
```
```sh
$ ./bsort 5 4 3 2 1
4 3 2 1 5
3 2 1 4 5
2 1 3 4 5
1 2 3 4 5
1 2 3 4 5
```

## Quick Sort
快速排序是由东尼·霍尔所发展的一种排序算法。在平均状况下，排序 n 个项目要Ο(n log n)次比较。在最坏状况下则需要Ο(n2)次比较，但这种状况并不常见。事实上，快速排序通常明显比其他Ο(n log n) 算法更快，因为它的内部循环（inner loop）可以在大部分的架构上很有效率地被实现出来。
快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
步骤为：
 1. 从数列中任选一元素，称为 "基准"（pivot），
 2. 以基准将数列划分为2部分，其中一部分为小于pivot值的元素集合，另一部分为大于pivot值的元素集合
 3. 递归处理由2产生的非空子集

### zsh
```sh
function qsort() {
    local pivot=$1
    local -a lt gt
    [[ -z $pivot ]] && return
        
    for x in $argv[2,-1]; do
        if (( x > pivot )); then
            gt+=$x
        else
            lt+=$x
        fi
    done
    qsort $lt
    echo $pivot
    qsort $gt
}
```
```sh
$ qsort 3 2 1
1
2
3
# 递归版本实际上有一些使用限制,因为在zsh中函数嵌套的深度是有限制的
$ qsort {1..1000}
1
...
maximum nested function level reached
```

###  haskell
```hs
{- qsort.hs -}
qsort :: (Ord a) => [a] -> [a]
qsort [] = []
qsort (x:xs) = (qsort [l| l<-xs, l<x]) ++ [x] ++ (qsort [g|g<-xs, g>=x])
```
```hs
$ ghci 
> :load qsort.hs
> qsort [1,3,7,8,2,0,4]
[0,1,2,3,4,7,8]
```

## Selection Sort
### python
```python
#! /usr/bin/python
# -*- coding: utf-8 -*-
def selection_sort(un_sorted):
	for i in range(len(un_sorted), 1, -1):
		max_pos = 0
		for j in range(0, i):
			if un_sorted[j] > un_sorted[max_pos]:
				max_pos = j
			
		un_sorted[i - 1], un_sorted[max_pos] = un_sorted[max_pos], un_sorted[i - 1]
```

### haskell
```hs
-- sort.hs
selectionSort :: (Ord a) => [a] -> [a]
selectionSort [] = []
selectionSort xs = [x|x<-xs, x == min] ++ (selectionSort [x|x<-xs, x>min])
  where min = minimum xs
```
```hs
$ ghci 
> :load sort.hs
> selectionSort []
[]
> selectionSort [9,5,2,1,8.0]
[0,1,2,5,8,9]
```

## 参考
 * http://www.keithschwarz.com/interesting/
 * http://coolshell.cn/articles/4671.html
 * http://www.brpreiss.com/books/opus4/html/page494.html
 * http://learnyousomeerlang.com/recursion
