package main

import (
    "log"
    "net/http"
    "fmt"
  	"github.com/stianeikeland/go-rpio"
  	"os"
    "github.com/gorilla/mux"
)

var (
  //gpio pin 17 (physical)
	pin = rpio.Pin(17)
)

func LedControll(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    if err := rpio.Open(); err != nil {
  		fmt.Println(err)
  		os.Exit(1)
  	}

  	// Unmap gpio memory when done
  	defer rpio.Close()

  	// Set pin to output mode
  	pin.Output()
    if "on" == params["id"] {
      //led on
      pin.PullUp()
      w.WriteHeader(http.StatusOK)
      w.Write([]byte("LED On!\n"))
    } else if "off" == params["id"] {
      // led off
      pin.PullDown()
      w.WriteHeader(http.StatusOK)
      w.Write([]byte("LED Off!\n"))
    } else {
      w.Write([]byte("Enter /on or /off\n"))
    }
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/{id}", LedControll).Methods("GET")
  log.Fatal(http.ListenAndServe(":12345", router))
}
