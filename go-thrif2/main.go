package main

import (
    "context"
    "fmt"
    "hbase"
    "time"
    "github.com/apache/thrift/lib/go/thrift"
)

const HOST = "10.128.0.7"
const PORT = "9090"

func main() {
    table := "my-table"
    rowkey := "10"
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

    client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
    if err := transport.Open(); err != nil {
        panic(err)
    }
    defer transport.Close()

    tmpendTime := currentTimeMillis()
    fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))

    // 写入数据
    // fmt.Printf(logformatstr_, "调用Put方法写数据")
    // putData(client, table, rowkey, family)

    // 读取数据
    fmt.Printf(logformatstr_, "调用Get方法读数据")
    getData(client, table, rowkey, family)

    fmt.Printf(logformatstr, logformattitle)
}

func putData(client *hbase.THBaseServiceClient, table, rowkey, family string) {
    cvarr := []*hbase.TColumnValue{
        {
            Family:    []byte(family),
            Qualifier: []byte("idoall.org"),
            Value:     []byte("welcome idoall.org"),
        },
    }
    temptput := hbase.TPut{
        Row:           []byte(rowkey),
        ColumnValues:  cvarr,
    }

    err := client.Put(
        context.Background(),
        []byte(table),
        &temptput,
    )

    if err != nil {
        fmt.Printf("Put err:%s\n", err)
    } else {
        fmt.Println("Put done")
    }
}

func getData(client *hbase.THBaseServiceClient, table, rowkey, family string) {
    tget := &hbase.TGet{
        Row: []byte(rowkey),
    }

    result, err := client.Get(
        context.Background(),
        []byte(table),
        tget,
    )

    if err != nil {
        fmt.Printf("Get err:%s\n", err)
    } else {
        fmt.Println("Get done")
        for _, columnValue := range result.ColumnValues {
            fmt.Printf("Family: %s, Qualifier: %s, Value: %s\n",
                string(columnValue.Family),
                string(columnValue.Qualifier),
                string(columnValue.Value))
        }
    }
}

func currentTimeMillis() int64 {
    return time.Now().UnixNano() / 1000000
}
