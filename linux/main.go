package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "os/exec"
    "time"
)
type Config struct {
    List []string
    Interval int
}
var config = Config{}
func randomPaper(args []string) {
    index := rand.Int31n(int32(len(config.List)))
    args[3] = config.List[index]
}
func init() {
    data,err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Fatal("read config file err:",err)
    }
    err = json.Unmarshal(data,&config)
    if err != nil{
        log.Fatal("unmarshal err:",err)
    }
    if len(config.List) <= 0 {
        log.Fatal("list no exist")
    }
    if config.Interval < 3 {
        log.Fatal("interval is too fast")
    }
}
func main() {
    args := []string{
        "set","org.gnome.desktop.background","picture-uri","",
    }
    for {
        randomPaper(args)
        cmd := exec.Command("gsettings",args...)
        cmd.Run()
        i := time.Second * (time.Duration(config.Interval))
        time.Sleep(i)
    }
}