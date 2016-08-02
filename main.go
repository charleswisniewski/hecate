package main

import ("fmt"
        "net/http"
        "bytes"
        "io/ioutil"
        "os")

func main() {
    url := os.Getenv("hygieia_base")
    if url == "" {
        panic("hygieia_base environment variable not set!")
    } else {
        url = url + "/api/deploy"
    }
    fmt.Println("URL:>", url )
    var jsonString = os.Args[1]
    body := bytes.NewBufferString(jsonString)
    r, err := http.Post(url,  "application/json", body)

    if err != nil {
        panic(err)
    }
    response, _ := ioutil.ReadAll(r.Body)
    fmt.Println("response Body:", string(response))
}