#!/bin/bash
set -x

case "$1" in
"amd64" | "arm64") ARCH="$1" ;;
"") ARCH="amd64" ;;
*) echo "Unknown 1st option, ARCH:$1" && exit 1 ;;
esac

pwd=$(cd "$(dirname "$0")" && pwd)
basedir=$(dirname "$pwd")
tag=$(cat "$(dirname "$basedir")/version")

remotePrefix=ericyami
imageName=callchain-c

docker build --platform linux/"$ARCH" -f "$pwd"/Dockerfile -t $remotePrefix/$imageName:"$tag" "$basedir"
docker tag $remotePrefix/$imageName:"$tag" $remotePrefix/$imageName:latest
docker push $remotePrefix/$imageName:"$tag"
docker push $remotePrefix/$imageName:latest
