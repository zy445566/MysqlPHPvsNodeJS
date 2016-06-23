
var mysql      = require('mysql');
var connection = mysql.createConnection({
  host     : 'localhost',
  user     : 'root',
  password : '',
  database : 'test'
});
console.log("mysql connect....");
connection.connect();
console.log("mysql clean table....");
connection.query("TRUNCATE TABLE `PHPvsNodeJS`");
console.log("start test....");
console.log("mysql write....");
var startWriteTime = Date.now();
for (var i = 1; i <= 100000; i++) {
	connection.query("INSERT INTO `PHPvsNodeJS` SET `text` = 'njstext"+i+"'");
};
var endWriteTime = Date.now();

console.log("mysql read....");
var startReadTime = Date.now();
for (var i = 1; i <= 100000; i++) {
	connection.query("SELECT * FROM `PHPvsNodeJS` WHERE `id` = "+i, function(err, rows, fields) {
		// console.log(rows);
	});
};
var endReadTime = Date.now();
console.log("--------------------------------------------------------------");
console.log("NodeJS Wirte Time:"+((endWriteTime-startWriteTime)/1000)+"s");
console.log("NodeJS Read Time:"+((endReadTime-startReadTime)/1000)+"s");
console.log("--------------------------------------------------------------");

connection.end();