package main

import (
 "context"
 "fmt"
 "hbase"
 "time"
 "github.com/apache/thrift/lib/go/thrift"
)

const HOST = "10.128.0.23"
const PORT = "9090"

func main() {
 table := "my-table"
 rowkey := "1"
 family := "cf1"
 startTime := currentTimeMillis()
 logformatstr_ := "----%s\n"
 logformatstr := "----%s 用时:%d-%d=%d毫秒\n\n"
 logformattitle := "建立连接" 
 protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
 transport, err := thrift.NewTSocket(HOST + ":" + PORT)
 if err != nil {
   panic(err)
 }

 client:=hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
 if err := transport.Open(); err != nil {
   panic(err)
 }
 tmpendTime := currentTimeMillis()
 fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))


 logformattitle = "调用Put方法写数据"
 fmt.Printf(logformatstr_, logformattitle)

 cvarr := []*hbase.TColumnValue{
 {
 Family: []byte(family),
 Qualifier: []byte("idoall.org"),
 Value: []byte("welcome idoall.org"),
 },
 }
 temptput := hbase.TPut{Row: []byte(rowkey), ColumnValues: cvarr}
 err = client.Put(context.Background(),[]byte(table), &temptput)
 if err != nil {
 fmt.Printf("Put err:%s\n", err)
 } else {
 fmt.Println("Put done")
 }
 fmt.Printf(logformatstr, logformattitle)


}

func currentTimeMillis() int64 {
 return time.Now().UnixNano() / 1000000
}

