#!/bin/bash
set -x
pwd=$(cd $(dirname $0); pwd)
basedir=$(dirname $pwd)
tag=$(cat $(dirname $basedir)/version)

remotePrefix=ericyami
imageName=callchain-a

docker build -t $remotePrefix/$imageName:$tag -f $pwd/Dockerfile $basedir
docker tag $remotePrefix/$imageName:$tag $remotePrefix/$imageName:latest
docker push $remotePrefix/$imageName:$tag
docker push $remotePrefix/$imageName:latest
