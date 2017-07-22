package main

import(
  // "fmt"
  "encoding/csv"
  "io"
  "os"
  "strings"
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)

func main(){
  /**
  * CSVファイル読み込み
  **/
  var fp *os.File
  if len(os.Args) < 2 {
    fp = os.Stdin
  } else {
    var err error
    fp, err = os.Open(os.Args[1])
    if err != nil {
            panic(err)
    }
    defer fp.Close()
  }

  /**
  * DB接続
  **/
  dbconf := "root:verysecret@tcp(mysqld:13306)/customer"
  db, err := sql.Open("mysql", dbconf)
  if err != nil{
    log.Fatal("open error: ", err)
  }
  defer db.Close()

  /**
  * DB書き込み
  **/

  reader := csv.NewReader(fp)
  reader.Comma = '\t'
  reader.LazyQuotes = true

  stmtIns, err := db.Prepare("INSERT INTO customers(name, sex, tel) values(?,?,?)")

  for{
    record, err := reader.Read()
    if err == io.EOF{
      break
    }else if err != nil{
      panic(err)
    }
    arr := strings.Split(record[0], " ")

    // DB書き込み
    name := arr[1]
    var sex int
    if arr[2] == "男"{
      sex = 0
    }else if arr[2] == "女"{
      sex = 1
    }
    tel := arr[3]

    _, err = stmtIns.Exec(name, sex, tel)
    if err != nil{
      log.Fatal("insert error: ", err)
    }
  }
}
