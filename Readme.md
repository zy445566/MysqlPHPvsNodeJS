#Mysql
* need import test.sql

#PHP
* php need open extionsion 
* php_pdo_mysql.dll or php_pdo_mysql.so

#Nodejs
* ModeJS need open module
* npm install mysql

#Golang
* need install git,and need set $GoPath
* go get github.com/go-sql-driver/mysql


#result

```sh
>php MySqlWRtest.php
    mysql connect....
    mysql clean table....
    start test....
    mysql write....
    mysql read....
    
    --------------------------------------------------------------

    PHP Wirte Time:7.5462689399719s
    PHP Read Time:9.7443377971649s

    --------------------------------------------------------------
```
```sh
>node MySqlWRtest.js
    mysql connect....
    mysql clean table....
    start test....
    mysql write....
    mysql read....

    --------------------------------------------------------------

    NodeJS Wirte Time:1.049s
    NodeJS Read Time:1.05s

    --------------------------------------------------------------
```
```sh
>lua MySqlWRtest.lua
    mysql connect....
    mysql clean table....
    mysql clean table....
    start test...
    mysql write....
    mysql read....
    --------------------------------------------------------------
    Lua Wirte Time:9.921s
    Lua Read Time:11.806s
    --------------------------------------------------------------
```
```sh
>go run MySqlWRtest.go
    mysql connect....
    mysql clean table....
    start test....
    mysql write....
    mysql read....
    --------------------------------------------------------------
    GoLang Wirte Time: 8.763000011444092
    GoLang Wirte Time: 9.901000022888184
    --------------------------------------------------------------
```
