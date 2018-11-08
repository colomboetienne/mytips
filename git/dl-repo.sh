#!/bin/bash

repos=( "mytips2" "mytips" "test2" )

for repo in ${repos[@]}
do
        echo "Download ${repo} ..."
        if ! git clone git@github.com:colomboetienne/${repo}.git -q
        then
            echo "-------------------------------------"
            echo Failed to download repo $repo
            echo "-------------------------------------"
        fi
        echo ""
done
