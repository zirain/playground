-- if-else
local x = 10
if x > 0 then print("x is a positive number") end

x = -1
if x > 0 then
    print("x is a positive number")
else
    print("x is a non-positive number")
end

local score = 90
if score == 100 then
    print("Very good!Your score is 100")
elseif score >= 60 then
    print(
        "Congratulations, you have passed it,your score greater or equal to 60")
    -- 此处可以添加多个elseif
else
    print("Sorry, you do not pass the exam! ")
end

-- while

x = 1
local sum = 0

while x <= 5 do
    sum = sum + x
    x = x + 1
end
print(sum) -- >output 15

-- break
local t = {1, 3, 5, 8, 11, 18, 21}
for i, v in ipairs(t) do
    if 11 == v then
        print("index[" .. i .. "] have right value[11]")
        break
    end
end

-- repeat

-- x = 10
-- repeat
--     print(x)
-- until false
-- dead lock

-- for
for i = 1, 5 do print(i) end
-- output:
--  1
--  2
--  3
--  4
--  5

-- 打印数组a的所有值
local a = {"a", "b", "c", "d"}
for i, v in ipairs(a) do print("index:", i, " value:", v) end

-- output:
-- index:  1  value: a
-- index:  2  value: b
-- index:  3  value: c
-- index:  4  value: d

local days = {
    "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"
}

local revDays = {}
for k, v in pairs(days) do revDays[v] = k end

-- print value
for k, v in pairs(revDays) do print("k:", k, " v:", v) end



-- break
sum = 0
local i = 1
while true do
    sum = sum + i
    if sum > 100 then
        break
    end
    i = i + 1
end
print("The result is " .. i)  -->output:The result is 14


-- return
local function add(x, y)
    return x + y
    --print("add: I will return the result " .. (x + y))
    --因为前面有个return，若不注释该语句，则会报错
end

local function is_positive(x)
    if x > 0 then
        return x .. " is positive"
    else
        return x .. " is non-positive"
    end

    --由于return只出现在前面显式的语句块，所以此语句不注释也不会报错
    --，但是不会被执行，此处不会产生输出
    print("function end!")
end

sum = add(10, 20)
print("The sum is " .. sum)  -->output:The sum is 30
local answer = is_positive(-10)
print(answer)                -->output:-10 is non-positive


-- goto
for i=1, 3 do
    if i <= 2 then
        print(i, "yes continue")
        goto continue
    end

    print(i, " no continue")

    ::continue::
    print([[i'm end]])
end


local function process(input)
    print("the input is", input)
    if input < 2 then
        goto failed
    end
    -- 更多处理流程和 goto err

    print("processing...")
    do return end
    ::failed::
    print("handle error with input", input)
end

process(1)
process(3)