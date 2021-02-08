#! /bin/bash

git submodule update --init --recursive
git submodule update --remote --merge
yarn --verbose --cwd ui install
yarn --verbose --cwd ui build
mkdir -p build/ui
mv ui/build/* build/ui
go build -v -o build/turrium