package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Data struct {
    ID        string   `json:"id,omitempty"`
    DeviceID  string   `json:"devid,omitempty"`
    Temp      float64      `json:"temp,omitempty"`
    Hum       float64      `json:"hum,omitempty"`
}

var data []Data

func GetDataEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range data {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Data{})
}


func GetDatasEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(data)
}



func DeleteDataEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range data {
        if item.ID == params["id"] {
            data = append(data[:index], data[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(data)
}

func CreateDataEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var datanew Data
    _ = json.NewDecoder(req.Body).Decode(&datanew)
    datanew.ID = params["id"]
    data = append(data, datanew)
    json.NewEncoder(w).Encode(data)
}

func main() {
    router := mux.NewRouter()
    data = append(data, Data{ID: "1", DeviceID: "dev01", Temp: 27.4, Hum: 80 })
    data = append(data, Data{ID: "2", DeviceID: "dev03", Temp: 22.568, Hum: 66 })
    router.HandleFunc("/data", GetDatasEndpoint).Methods("GET")
    router.HandleFunc("/data/{id}", GetDataEndpoint).Methods("GET")
    router.HandleFunc("/data/{id}", CreateDataEndpoint).Methods("POST")
    router.HandleFunc("/data/{id}", DeleteDataEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
}
