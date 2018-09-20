#!/bin/bash

hexo g && {
    git add -A
    git commit -am "Commit From IDEA: $(date '+%Y-%m-%d %H:%M')"
    git pull
    git push
}
