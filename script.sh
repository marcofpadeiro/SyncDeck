#!/bin/bash

for file in `find . -name '*.go' -or -name '*.mod'`; do \
     cat $file | sed s/ZipFolder/Compress/g > tmp
     mv tmp $file
done
