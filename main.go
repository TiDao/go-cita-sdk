/***********************************************************************
# File Name: main.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2020-09-09 15:09:28
*********************************************************************/
package main

import(
    "encoding/json"
    "./cita"
    "fmt"
//    "os"
    "bytes"
    "log"
)

func main(){
    url := "http://192.168.1.65:1337"
    request := &cita.Request{
        Jsonrpc: "2.0",
        Method: "getMetaData",
        Params: []interface{}{"latest"},
        Id: 1,
    }

    requestJson,err := json.Marshal(request)
    if err != nil{
        fmt.Println(err)
    }
    responseJson := &[]byte{}

    *responseJson,err = cita.Post(url,requestJson)
    if err != nil{
        log.Fatal(err)
    }

    var response = &cita.Response{}

    err = json.Unmarshal(*responseJson,response)
    if err != nil{
        fmt.Println(err)
    }

    fmt.Println(string(*responseJson))

    var out bytes.Buffer

    err = json.Indent(&out,*responseJson,"","\t")
    fmt.Println(out.String())
    //out.WriteTo(os.Stdout)
}