#!/bin/bash

# 使用变量
your_name="qinjx"
echo $your_name
echo ${your_name}

for skill in Ada Coffe Action Java; do
    echo "I am good at ${skill}Script"
done

your_name="tom"
echo $your_name
your_name="alibaba"
echo $your_name

# 只读变量
myUrl="http://www.google.com"
readonly myUrl
# myUrl="http://www.runoob.com" myUrl: readonly variable


# 删除变量

myUrl2="http://www.runoob.com"
echo $myUrl2
unset myUrl2
echo $myUrl2