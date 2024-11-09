#!/bin/bash

platforms=("windows/amd64" "linux/amd64")

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  goos=${platform_split[0]}
  goarch=${platform_split[1]}
  output_name=build/gofit
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	
  env GOOS=$goos GOARCH=$goarch go build -o $output_name cmd/gofit/main.go
done
