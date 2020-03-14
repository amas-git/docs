---
title: Zsh Subshell
tags:
---

## Current Shell          
运行在当前shell中的语句包括:
 1. ![[...]]
 2. [wiki:Zsh/ComplexCommands if|while|repeat|select|for]
 3. [wiki:Zsh/Functions functions]
 4. source | .
 5. {...} 
 6. ... | `...` (管道右面的部分, 只在Zsh中有可靠的保障,其他shell也许有所不同)
 7. !`...` | $(...) | =(...) | <(...) | >(...)

## Subshell
 1. 所有外部命令
 2. `...` | ... (管道左面的部分)
 3. (...)
 4. [wiki:Zsh/UserGuide/Substitutions 代换过程中执行的命令行]
 5. 任何以&结尾的命令行
 6. 任何被挂起(suspend)命令. 比如你正在执行一系列的命令,然后^Z挂起当前执行的命令. Zsh为了能够与你交互,就会将正在执行的命令转入subshell.



subshell:
```zsh
#!/bin/zsh
(
local x=1
 (
    local x=2
    ( local x=3 ; echo $x )
    echo $x
 )
echo $x
)
```
```zsh
$ subshell
3
2
1
```
