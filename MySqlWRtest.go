package main
  
import (
  "database/sql"
  "fmt"
  "time"
  "strconv"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("mysql connect....")
	db, _ := sql.Open("mysql", "root:@/test?charset=utf8")
	fmt.Println("mysql clean table....")
	db.Query("TRUNCATE TABLE `PHPvsNodeJS`")
	fmt.Println("start test....")
	fmt.Println("mysql write....")
	startWriteTime := makeTimestamp()/1000.0
	stmt1, _ := db.Prepare("INSERT INTO `PHPvsNodeJS` SET `text` = ?")
	for i := 1; i <= 100000; i++ {
    	stmt1.Exec("goLtext"+strconv.Itoa(i))
    }
    endWriteTime := makeTimestamp()/1000.0

    fmt.Println("mysql read....")
    startReadTime := makeTimestamp()/1000.0
	stmt2, _ := db.Prepare("SELECT * FROM `PHPvsNodeJS` WHERE `id` = ?")
	for i := 1; i <= 100000; i++ {
    	stmt2.Exec(i)
    }
    endReadTime := makeTimestamp()/1000.0
	// for rows.Next() {
	// 	var id int
	// 	err = rows.Scan(&id)
	// 	fmt.Println(id)
	// }
	db.Close()
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("GoLang Wirte Time:", endWriteTime-startWriteTime)
	fmt.Println("GoLang Wirte Time:", endReadTime-startReadTime)
	fmt.Println("--------------------------------------------------------------")

}

func makeTimestamp() float64 {
    return float64 (time.Now().UnixNano() / int64(time.Millisecond))
}
