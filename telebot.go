package telebot

import "net/http"

func Get(url string)(resp *http.Response){
    resp, _ = http.Get(url)
    return resp
}
