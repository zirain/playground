#!/bin/sh

your_name='runoob'

str='Hello, I know you are \"$your_name\"! \n'
echo $str


str="Hello, I know you are \"$your_name\"! \n"
echo $str


# 拼接字符串

# 使用双引号拼接
greeting="hello, "$your_name" !"
greeting_1="hello, ${your_name} !"
echo $greeting  $greeting_1
# 使用单引号拼接
greeting_2='hello, '$your_name' !'
greeting_3='hello, ${your_name} !'
echo $greeting_2  $greeting_3

# 获取字符串长度
string="abcd"
echo ${#string} #输出 4

# 提取子字符串
string="runoob is a great site"
echo ${string:1:4} # 输出 unoo

# 查找子字符串
string="runoob is a great site"
echo `expr index "$string" io`  # 输出 4