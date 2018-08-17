#!/bin/bash

rm xhCrm -f
rm web.tar.gz -f
go build src/xhCrm.go
tar zcvf web.tar.gz web
echo 'sshpass -p hihShSe4K2wZKwk6 scp xhCrm root@47.97.49.163:/root'
sshpass -p hihShSe4K2wZKwk6 scp xhCrm root@47.97.49.163:/root
sshpass -p hihShSe4K2wZKwk6 scp web.tar.gz root@47.97.49.163:/root

