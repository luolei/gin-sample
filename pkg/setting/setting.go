package setting

import (
    "fmt"
    "log"
    "time"
    "github.com/go-ini/ini"
)

var (
    Cfg *ini.File

    RunMode string
    HttpPort int
    ReadTimeout time.Duration
    WriteTimeout time.Duration

)

func init() {
    var err error
    Cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err);
    } else {
        //log.Fatalf("load 'conf/app.ini'");
    }

    LoadBase()
    LoadServer()
    //LoadApp()
}

func LoadBase() {
    RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
    fmt.Printf("run mode: %s", RunMode)
}

func LoadServer() {
    sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("Fail to get section 'server': %v", err)
    }

    RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
    HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
    ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
    WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

    fmt.Printf("load server mode: %s, port: %d, read_timeout: %d, write_timeout: %d", RunMode, HttpPort, ReadTimeout, WriteTimeout)
}

/*
func main(){
    fmt.Println("Hello, World!")
}
*/

