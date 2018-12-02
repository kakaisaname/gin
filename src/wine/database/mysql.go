package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
//golang中的连接来自内部实现的连接池，连接的建立是惰性的，当你需要连接的时候，连接池会自动帮你创建。通常你不需要操作连接池。一切都有go来帮你完成
var Db *sql.DB  //Db必须要为大写才能被外部访问
func init() {

	var err error
	//db, err := sql.Open("mysql","user:password@tcp(127.0.0.1:3306)/hello")
	Db, err = sql.Open("mysql","root:123456@tcp(118.24.61.194:3306)/yang?parseTime=true")
	if err != nil {
		//Fatal表示程序遇到了致命的错误，需要退出，这时候使用Fatal记录日志后，然后程序退出，也就是说Fatal相当于先调用Print打印日志，然后再调用os.Exit(1)退出程序。
		log.Fatal(err.Error())
	}
	//err = Db.Ping() //如果想立即验证连接，需要用Ping()方法
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	//rows, err := Db.Query("SELECT email FROM user")   //查询测试
	//if err!= nil {
	//	log.Fatalln(err)
	//}
	//for rows.Next(){
	//	var s string
	//	err = rows.Scan(&s)
	//	if err!= nil {
	//		log.Fatalln(err)
	//	}
	//	log.Printf("found row containing %q",s)
	//}
	//rows.Close()
}
