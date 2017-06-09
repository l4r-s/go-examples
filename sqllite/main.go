package main

 import (
     "database/sql"
     "fmt"
     _"github.com/mattn/go-sqlite3"
 )

 func main() {
     db, err := sql.Open("sqlite3", "./foo.db")
     checkErr(err)

     // insert
     stmt, err := db.Prepare("INSERT INTO data(devid, temp, hum) values(?,?,?)")
     checkErr(err)

     res, err := stmt.Exec("dev01", "9.9", "89.3")
     checkErr(err)

     id, err := res.LastInsertId()
     checkErr(err)

     fmt.Println(id)
     // update
     //stmt, err = db.Prepare("update userinfo set username=? where uid=?")
     //checkErr(err)

     //res, err = stmt.Exec("astaxieupdate", id)
     //checkErr(err)

     //affect, err := res.RowsAffected()
     //checkErr(err)

     //fmt.Println(affect)

     // query
     rows, err := db.Query("SELECT * FROM data")
     checkErr(err)
     var dbid int
     var devid string
     var temp float64
     var hum float64

     for rows.Next() {
         err = rows.Scan(&dbid, &devid, &temp, &hum)
         checkErr(err)
         fmt.Println(dbid)
         fmt.Println(devid)
         fmt.Println(temp)
         fmt.Println(hum)
     }

     rows.Close() //good habit to close

     // delete
     //stmt, err = db.Prepare("delete from userinfo where uid=?")
     //checkErr(err)

     //res, err = stmt.Exec(id)
     //checkErr(err)

     //affect, err = res.RowsAffected()
     //checkErr(err)

     //fmt.Println(affect)

     //db.Close()

 }

 func checkErr(err error) {
     if err != nil {
         panic(err)
     }
 }
