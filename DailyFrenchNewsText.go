package main

import (
    "fmt"
    "strings"
    "net/http"
    "net/url"
    "math/rand"
    "time"
    "encoding/json"
    "io/ioutil"
    "github.com/jasonlvhit/gocron"
)

func main() {
    s := gocron.NewScheduler()
    s.Every(1).Day().Do(task)
    <- s.Start()
}

func task() {

    // urrl here b/c url used for twilio portion, need to be different

    urrl := "https://newsapi.org/v2/top-headlines?apiKey=APIKEY&language=fr"

    reqq, _ := http.NewRequest("GET", urrl, nil)

    reqq.Header.Add("cache-control", "no-cache")
    reqq.Header.Add("Postman-Token", "dcc8010d-6bc7-4112-93bb-98eae13fcdb2")

    res, _ := http.DefaultClient.Do(reqq)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    fmt.Println(res)
    // fmt.Println(string(body))

    var x string
    x = (string(body))

    var y string
    y = x[56:1600]

    //100 - 1650 works, 1550 (maybe 1600) appears to be the character limit for Twilio texts
    

    fmt.Println(y)
    // https://www.twilio.com/blog/2017/09/send-text-messages-golang.html

        accountSid := "ACCOUNTSID"
        authToken := "AUTHTOKEN"
        urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

         rand.Seed(time.Now().Unix())
         varieble := url.Values{}
         varieble.Set("To","+1234567890")
         varieble.Set("From","+1234567890")
         varieble.Set("Body",(string(y)))
         msgDataReader := *strings.NewReader(varieble.Encode())

         client := &http.Client{}
req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
req.SetBasicAuth(accountSid, authToken)
req.Header.Add("Accept", "application/json")
req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

         resp, _ := client.Do(req)
if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
  var data map[string]interface{}
  decoder := json.NewDecoder(resp.Body)
  err := decoder.Decode(&data)
  if (err == nil) {
    fmt.Println(data["sid"])
  }
} else {
  fmt.Println(resp.Status);
}

}


