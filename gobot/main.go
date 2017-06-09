package main

import (
        "time"
        "fmt"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/aio"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)

func main() {
        firmataAdaptor := firmata.NewTCPAdaptor("192.168.15.226:3030")
        led := gpio.NewLedDriver(firmataAdaptor, "2")
        dht22 := aio.NewGroveTemperatureSensorDriver(firmataAdaptor, "4")

        work := func() {
                gobot.Every(1*time.Second, func() {
                        temp, _ := dht22.Read()
                        fmt.Println(temp)
                })
        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led},
                work,
        )

        robot.Start()
}
