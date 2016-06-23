<?php
echo "mysql connect....\n";
$dbh = new PDO('mysql:host=localhost;dbname=test', 'root', '');
echo "mysql clean table....\n";
$sql = "TRUNCATE TABLE `PHPvsNodeJS`";  
$stmt = $dbh->prepare($sql);  
$stmt->execute();
echo "start test....\n";

echo "mysql write....\n";
$startWriteTime=microtime(true);
for ($i=1; $i <= 100000; $i++) { 
	$sql = "INSERT INTO `PHPvsNodeJS` SET `text` = :text";  
	$stmt = $dbh->prepare($sql);  
	$stmt->execute(array(':text' => 'phptext'.$i));
}
$endWriteTime=microtime(true);

echo "mysql read....\n";
$startReadTime=microtime(true);
for ($i=1; $i <= 100000; $i++) {
	$sql = "SELECT * FROM `PHPvsNodeJS` WHERE `id` = :id";  
	$stmt = $dbh->prepare($sql);  
	$stmt->execute(array(':id' => $i));
	// var_dump( $stmt->fetchAll(PDO::FETCH_ASSOC)); 
}
$endReadTime=microtime(true);

echo "--------------------------------------------------------------\n";
echo "PHP Wirte Time:".($endWriteTime-$startWriteTime)."s\n";
echo "PHP Read Time:".($endReadTime-$startReadTime)."s\n";
echo "--------------------------------------------------------------\n";