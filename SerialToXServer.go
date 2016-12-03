package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/tarm/serial"
    "github.com/go-vgo/robotgo"
)

const arduinoPort = "/dev/ttyACM0"

func openSerial() *serial.Port {
    c := &serial.Config{Name: arduinoPort, Baud: 115200}
    s, err := serial.OpenPort(c)
    if err != nil {
        log.Fatal(err)
    }
    return s
}

func getByteFromPort(port *serial.Port) byte {
    buf := make([]byte, 1)
    _, err := port.Read(buf)
    if err != nil {
            log.Fatal(err)
    }
    return buf[0]
}

func getKeyFromByte(b byte) string {
    switch b {
    case 7: //left
        return "left"
    case 6: //right
        return "right"
    case 5: //up
        return "up"
    case 4: //down
        return "down"
    case 3: //down
        return "enter"
    default: //wtf
        return ""
    }
}

func waitForInterrupt(c chan os.Signal, port *serial.Port) {
    <-c
    err := port.Close()
    if (err == nil) {
        log.Println("!!!")
        log.Println("Successfully closed port")
    } else {
        log.Fatal(err)   
    }

    os.Exit(1)
}

func main() {
    s := openSerial()
    
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go waitForInterrupt(c, s)

    var lastKey string = "" 
    for true {
        b := getByteFromPort(s)
        key := getKeyFromByte(b)
        if (key == "" && lastKey != "") {
            robotgo.KeyToggle(lastKey, "up")
            lastKey = ""
        } else if (key != lastKey) {
            if (lastKey != "") {
                robotgo.KeyToggle(lastKey, "up")
            }
            robotgo.KeyToggle(key, "down")
            lastKey = key
        } 
    }
}
