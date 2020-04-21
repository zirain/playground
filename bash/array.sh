#!/bin/sh

arr1=(1 2 3 4)
arr2=(
    1
    2
    3
    4
)

arr3[0]=1
arr3[1]=2
arr3[2]=3
arr3[3]=4
echo 'arr1内容:' ${arr1[@]}
echo 'arr2内容:' ${arr2[@]}
echo 'arr3内容:' ${arr3[@]}

echo 'length:' ${#arr1[@]}