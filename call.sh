#!/bin/bash

iter=$3
for i in `seq 1 $iter`
do
    curl http://$1/$2
done
