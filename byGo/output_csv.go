package main

import(
  "encoding/csv"
  iconv "github.com/djimenez/iconv-go"
  "log"
  "os"
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

  writer := csv.NewWriter(converter)
  writer.Write([]string{"Kazuki", "20"})
  writer.Write([]string{"Yuki", "21"})
  writer.Write([]string{"Taro", "21"})
  writer.Flush()
}
