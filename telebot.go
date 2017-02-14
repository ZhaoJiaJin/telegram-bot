package telebot

import (
    "net/http"
    "io/ioutil"
    "errors"
    "time"
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
    receive bool
}

type Mess struct{
    Update_id float64
    Mess_id float64
    Chat interface{}
    Date float64
    Text string
    Entity interface{}
}

func (b *Bot)Stop_receive(){
    b.receive = false
}

func (b *Bot)Start_receive(mess_chan chan Mess){
    b.receive = true
    go b.Receive(mess_chan)
}

func (b *Bot)Receive(mess_chan chan Mess){
    for ;b.receive; {
        b.Getupdate(mess_chan)
        time.Sleep(10 * time.Millisecond)
    }

}

func (b *Bot)Getupdate(mess_chan chan Mess){
    url := fmt.Sprintf("%v/bot%v/getUpdates?offset=%v&timeout=30",b.Api,b.Key,b.Maxid+1)
    s,err := send(url)
    if err != nil{
        fmt.Println(err)
    }
    res := make(map[string]interface{})
    err = json.Unmarshal([]byte(s), &res)
    if err != nil{
        fmt.Println(err)
    }
    if res["ok"] != true{
        fmt.Println("resp err:"+s)
    }

    for _,v := range(res["result"].([]interface{})){
        dic_mess := v.(map[string]interface{})
        up_id := dic_mess["update_id"].(float64)
        if up_id > b.Maxid{
            mess_con := dic_mess["message"].(map[string]interface{})
            one_mess := Mess{Update_id:up_id,Mess_id:mess_con["message_id"].(float64),Chat:mess_con["chat"],Date:mess_con["date"].(float64),Text:mess_con["text"].(string),Entity:mess_con["entities"]}
            //fmt.Println(one_mess)
            fmt.Println("put message",one_mess.Mess_id)
            mess_chan <- one_mess
            b.Maxid = up_id
        }
    }
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
