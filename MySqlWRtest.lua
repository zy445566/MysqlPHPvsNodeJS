mysql = require "luasql.mysql"
print("mysql connect....");
local env  = mysql.mysql()
local conn = env:connect('test','root','')
print("mysql clean table....");
conn:execute([[TRUNCATE TABLE `PHPvsNodeJS`]])
print("mysql clean table....");
print("start test...");
print("mysql write....");
local startWriteTime = os.clock()
for i=1,100000,1 do
    conn:execute("INSERT INTO `PHPvsNodeJS` SET `text` = 'luatext"..i.."'")
end
local endWriteTime = os.clock()

print("mysql read....");
local startReadTime = os.clock()
for i=1,100000,1 do
    conn:execute("SELECT * FROM `PHPvsNodeJS` WHERE `id` = "..i)
end
local endReadTime = os.clock()

print("--------------------------------------------------------------");
print("Lua Wirte Time:"..(endWriteTime-startWriteTime).."s");
print("Lua Read Time:"..(endReadTime-startReadTime).."s");
print("--------------------------------------------------------------");

conn:close()
env:close()
