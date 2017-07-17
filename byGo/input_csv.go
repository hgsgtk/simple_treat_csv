package main

import(
  "fmt"
  "encoding/csv"
  "io"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main(){
  var fp *os.File
  var err error

  if len(os.Args) < 2{
    fp = os.Stdin
  }else{
    fmt.Printf(">> read file: %s\n", os.Args[1])
    fp, err = os.Open(os.Args[1])
    if err != nil{
      panic(err)
    }
    defer fp.Close()
  }

  reader := csv.NewReader(fp)
  reader.Comma = '\t'
  reader.LazyQuotes = true
  for{
    record, err := reader.Read()
    if err == io.EOF{
      break
    }else if err != nil{
      panic(err)
    }
    fmt.Println(record)
  }
}
