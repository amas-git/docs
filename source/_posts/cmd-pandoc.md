---
title: 使用pandoc迁移文档，从trac到github
date: 2018-03-19 15:39:04
tags:
---
```zsh
#!/bin/zsh
# simple script transform trac wiki format to github markdown

file=$1
[[ -f $file ]] || {
    print "PLEASE SPECIF FILE"
    exit 250
}

content=$(< $file)
local -a _chunks
_chunks=(${(@f)content})


content=($(< $file))

for ((i=1; i<=$#_chunks; ++i)); do
    _line=$_chunks[i]
    # process head
    if [[ $_line =~ '^([=])+\s(.*)' ]]; then
        head=${match[1]}
        head=${head:gs/=/#}
        print -- "$head $match[2]"
    # process block
    elif [[ $_line =~ '^}}}.*' ]];  then
        print "```"
    # process block
    elif [[ $_line =~ '^\{\{\{.*' ]];  then
        print "```"
    else
        print -- $_line
    fi
done
```
