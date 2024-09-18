package main

import (
	"log"
	"net"
	"net/http"
)

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		//fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
*/

// used to auto detect the active local IP address - not used yet
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
