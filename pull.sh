#!/bin/bash
#git --git-dir=/Users/andy/Desktop/metacloud/test/.git --work-tree=/Users/andy/Desktop/metacloud/test pull

#gitlab
ssh -o "StrictHostKeyChecking no" andy@dev "git --git-dir=/home/app/metacloud/.git --work-tree=/home/app/metacloud
pull"

#github
ssh -o "StrictHostKeyChecking no" andy@dev "docker run --rm -v /home/app:/go -w /go/metacloud bowwow/goapi go get -
d -v"
ssh -o "StrictHostKeyChecking no" andy@dev "docker restart api"
ssh -o "StrictHostKeyChecking no" andy@dev "date +\"restart ok ! %Y/%m/%d %H:%M:%S\""
