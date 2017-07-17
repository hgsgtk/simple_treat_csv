package main

import(
  "database/sql"
  "encoding/csv"
  iconv "github.com/djimenez/iconv-go"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "os"
  "fmt"
  "strconv"
)

func failOnError(err error){
  if err != nil{
    log.Fatal("Error: ", err)
  }
}

func main(){

  file, err := os.OpenFile("./output.csv", os.O_WRONLY|os.O_CREATE, 0600)
  failOnError(err)
  defer file.Close()

  err = file.Truncate(0)
  failOnError(err)

  converter, err := iconv.NewWriter(file, "utf-8", "sjis")
  failOnError(err)

  // DB接続
  dbconf := "root:verysecret@tcp(mysqld:13306)/customer"
  db, err := sql.Open("mysql", dbconf)
  if err != nil{
    log.Fatal("open error: ", err)
  }
  fmt.Println("connected!")
  defer db.Close()

  // DB READ
  rows, qerr := db.Query("SELECT * FROM customers")
  defer rows.Close()
  if qerr != nil{
    log.Fatal("query error: ", qerr)
  }

  writer := csv.NewWriter(converter)

  for rows.Next(){
    var id int
    var name string
    var sex int
    var tel string

    var idStr string
    var sexStr string

    if berr := rows.Scan(&id, &name, &sex, &tel); berr != nil{
      log.Fatal("scan erro: ", berr)
    }

    idStr = strconv.Itoa(id)

    if(sex == 0){
      sexStr = "男"
    }else{
      sexStr = "女"
    }

    writer.Write([]string{idStr, name, sexStr, tel})
  }
  writer.Flush()
}
