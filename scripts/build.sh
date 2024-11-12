#!/bin/bash

version=0.1.0
platforms=("linux/amd64" "darwin/amd64" "windows/amd64")

rm build/*

for platform in "${platforms[@]}"
do
  platform_parts=(${platform//\// })
  goos=${platform_parts[0]}
  goarch=${platform_parts[1]}
  output_name=build/gofixt-v$version.$goos-$goarch
	if [ $goos = "windows" ]; then
		output_name+='.exe'
	fi
  env GOOS=$goos GOARCH=$goarch go build -o $output_name cmd/gofixt/main.go
done
