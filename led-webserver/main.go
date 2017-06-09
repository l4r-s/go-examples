package main

import (
    "log"
    "net/http"
    "fmt"
    "github.com/stianeikeland/go-rpio"
    "github.com/gorilla/mux"
    "os"
)

var (
	// PIN19 GPIO10
	pin = rpio.Pin(10)
)

func LedOn(w http.ResponseWriter, req *http.Request) {
  if err := rpio.Open(); err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  if pin.Read() == 1 {
    w.WriteHeader(http.StatusOK)
	  fmt.Fprintf(w, "LED already enabled \n")
    fmt.Println("LED already enabled")
  } else {
    pin.Output()
    pin.Toggle()
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "LED enabled \n")
    fmt.Println("LED enabled")
  }
}

func LedOff(w http.ResponseWriter, req *http.Request) {
  if err := rpio.Open(); err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  if pin.Read() == 0 {
    w.WriteHeader(http.StatusOK)
	  fmt.Fprintf(w, "LED already disabled \n")
    fmt.Println("LED already disabled")
  } else {
    pin.Output()
    pin.Toggle()
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "LED disabled \n")
    fmt.Println("LED disabled")
  }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/on", LedOn).Methods("GET")
    router.HandleFunc("/off", LedOff).Methods("GET")

    log.Fatal(http.ListenAndServe(":12345", router))
}
