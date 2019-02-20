#!/bin/bash

for gitdir in `find ./ -maxdepth 2 -type d -name .git`; 
    do 
        workdir=${gitdir%/*}; 
        echo; 
        echo $workdir; 
        git --git-dir=$gitdir --work-tree=$workdir status -b -s; 
    done
