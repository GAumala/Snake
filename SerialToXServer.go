package main

import (
    "log"
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
    case 0:
        return "up"
    case 1:
        return "right"
    case 2:
        return "down"
    case 3:
        return "left"
    default:
        return "k"
    }
}


func main() {
    s := openSerial()
    for true {
        b := getByteFromPort(s)
        key := getKeyFromByte(b)
        robotgo.KeyTap(key)
    }
}
