#!/bin/bash

#git pull
go build main.go
mv main go-helios-da

netstat -anp | grep 9609

#cp go-helios-da ~/online/

#nohup ./go-helios-da > ./log/go-helios-da.log 2>&1 &