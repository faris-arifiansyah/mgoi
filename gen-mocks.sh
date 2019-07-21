#!/bin/bash
CURDIR=$(pwd)
for file in *.go
do
    echo "mockgen -source ${CURDIR}/${file} -destination=${CURDIR}/mgoimock/${file} -package mgoimock -imports \".=github.com/faris-arifiansyah/mgoi\""
    mockgen -source ${CURDIR}/${file} -destination=${CURDIR}/mgoimock/${file} -package mgoimock -imports ".=github.com/faris-arifiansyah/mgoi"
done
