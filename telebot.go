package telebot

import (
    "net/http"
    "io/ioutil"
    "errors"
    "fmt"
)


func checkErr(err error){
    if err != nil{
        panic(err)
    }
}


type Bot struct{
    Api string
    Key string
}


func (b Bot)Getme()(string, error){
    url := fmt.Sprintf("%v/bot%v/getMe",b.Api,b.Key)
    resp, err := http.Get(url)
    if err != nil{
        return "", errors.New("get request fail," + err.Error())
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return "", errors.New("read get request data fail," + err.Error())
    }
    s := string(body)
    return s,nil
}
