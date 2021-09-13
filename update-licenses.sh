#!/usr/bin/env bash

GO111MODULE=off go get github.com/google/go-licenses

rm -f ./licenses.md
rm -rf licenses/

go-licenses save ./... --save_path="./licenses"

touch ./licenses.md

echo "# Licenses" >> licenses.md

function traverse() {
for file in "$1"/*
do
    if [ ! -d "${file}" ] ; then
        echo "### $(echo ${file} | sed -r 's@licenses/(.+)/LICENSE@\1@g')" >> licenses.md
        echo "\`\`\`" >> licenses.md
        cat "${file}" >> licenses.md
        echo "" >> licenses.md
        echo "\`\`\`" >> licenses.md
        echo "" >> licenses.md
    else
        traverse "${file}"
    fi
done
}

function main() {
    traverse "$1"
    rm -rf licenses/
}

main "licenses"
