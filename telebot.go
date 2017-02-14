package telebot

import (
    "net/http"
    "io/ioutil"
    "errors"
    "encoding/json"
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
    Maxid float64
    mess chan string
}

type Mess struct{
    Update_id float64
    Mess_id float64
    Chat interface{}
    Date float64
    Text string
    Entity interface{}
}


func (b *Bot)GetChan(){
    if b.mess == nil{
        b.mess = make(chan string)
    }

}

func (b *Bot)Getupdate()(string,error){
    url := fmt.Sprintf("%v/bot%v/getUpdates",b.Api,b.Key)
    s,err := send(url)
    if err != nil{
        return "",err
    }
    res := make(map[string]interface{})
    err = json.Unmarshal([]byte(s), &res)
    if err != nil{
        return "",err
    }
    if res["ok"] != true{
        return "",errors.New("resp err; " + s)
    }

    for _,v := range(res["result"].([]interface{})){
        dic_mess := v.(map[string]interface{})
        up_id := dic_mess["update_id"].(float64)
        if up_id > b.Maxid{
            mess_con := dic_mess["message"].(map[string]interface{})
            one_mess := Mess{Update_id:up_id,Mess_id:mess_con["message_id"].(float64),Chat:mess_con["chat"],Date:mess_con["date"].(float64),Text:mess_con["text"].(string),Entity:mess_con["entities"]}
            fmt.Println(one_mess)
            b.Maxid = up_id
        }
    }
    return s,nil
}

func (b Bot)Getme()(string, error){
    url := fmt.Sprintf("%v/bot%v/getMe",b.Api,b.Key)
    return send(url)
}

func send(url string)(string, error){
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
