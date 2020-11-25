/***********************************************************************
# File Name: cita_method.go
# Author: TiDao
# mail: tidao2049@gmail.com
# Created Time: 2020-09-11 13:19:39
*********************************************************************/
package cita


import(
    "unsafe"
)


type errString struct{
    prefix string
    err  error
}

func (Err *errString)Error() string{
    return Err.prefix+Err.err.Error()
}

func byteToString(b []byte) string{
    return *(*string)(unsafe.Pointer(&b))
}

func byteToBool(b []byte) bool{
    return *(*bool)(unsafe.Pointer(&b))
}

func responseUnmarshal(req *Request,url *string) (*Response,error){

    Err := &errString{}
    reqjson,err := req.MarshalJSON()
    if err != nil{
        Err.prefix = "marshal request json error "
        Err.err = err
        return nil,Err
    }

    respjson,err := Post(*url,reqjson)
    if err != nil{
        Err.prefix = "post messages error "
        Err.err = err
        return nil,Err
    }

    var response = &Response{}
    err = response.UnmarshalJSON(respjson)
    if err != nil{
        Err.prefix = "unmarshal response  error "
        Err.err = err
        return nil,Err
    }

    return response,nil
}

func PeerCount(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func Call(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}
func BlockNumber(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func PeersInfo(req *Request,Result *ResultPeerInfo,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    errResult := Result.UnmarshalJSON(response.Result)
    if errResult != nil{
        return response.Error,errResult
    }
    return response.Error,nil
}


func GetVersion(req *Request,Result *ResultVersion,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    errResult := Result.UnmarshalJSON(response.Result)
    if errResult != nil{
        return response.Error,errResult
    }
    return response.Error,nil
}

func GetBlock(req *Request,Result *ResultBlock,url string) (Error,error) {
    
    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    errResult := Result.UnmarshalJSON(response.Result)
    if errResult != nil{
        return response.Error,errResult
    }
    return response.Error,nil
}

func GetLogs(req *Request,Result *[]ResultLogs,url string) (Error,error){

    Err := &errString{}
    reqJson,err := req.MarshalJSON()
    if err != nil{
        Err.prefix = "marshal request json error "
        Err.err = err
        return Error{},Err
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        Err.prefix = "Post request error "
        Err.err = err
        return Error{},Err
    }

    var response = &Logs{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        Err.prefix = "response UnmarshalJson error: "
        Err.err = err
        return response.Error,Err
    }

    return response.Error,nil
}

func GetTransaction(req *Request,Result *ResultTransaction,url string) (Error,error){

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    errResult := Result.UnmarshalJSON(response.Result)
    if errResult != nil{
        return response.Error,errResult
    }
    return response.Error,nil
}

func GetBalance(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func NewFilter(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func NewBlockFilter(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func UninstallFilter(req *Request,Result *bool,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToBool(response.Result)

    return response.Error,nil
}

func GetFilterChanges(req *Request,Result *[]string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }
    //*Result = byteToString(response.Result)
    for i,results := range response.Result {
        (*Result)[i] = *(*string)(unsafe.Pointer(&results))
    }
    return response.Error,nil
}

func GetFilterLogs(req *Request,Result *[]ResultLogs,url string) (Error,error) {

    Err := &errString{}
    reqJson,err := req.MarshalJSON()
    if err != nil{
        Err.prefix = "marshal request json error "
        Err.err = err
        return Error{},Err
    }

    respJson,err := Post(url,reqJson)
    if err != nil{
        Err.prefix = "Post request error "
        Err.err = err
        return Error{},Err
    }

    var response = &Logs{}
    err = response.UnmarshalJSON(respJson)
    if err != nil{
        Err.prefix = "Unmarshal response  error "
        Err.err = err
        return response.Error,Err
    }
    return response.Error,nil
}

func GetTransactionProof(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func GetBlockHeader(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func GetStateProof(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func GetStorageAt(req *Request,Result *string,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    *Result = byteToString(response.Result)

    return response.Error,nil
}

func GetMetaData(req *Request,Result *ResultMetaData,url string) (Error,error) {

    response,err := responseUnmarshal(req,&url)
    if err != nil{
        return response.Error,err
    }

    errResult := Result.UnmarshalJSON(response.Result)
    if errResult != nil{
        return response.Error,errResult
    }
    return response.Error,nil

}
