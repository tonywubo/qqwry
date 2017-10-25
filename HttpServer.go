package main

import (
	"encoding/json"
	"net/http"
)

func handlerIP(w http.ResponseWriter, req *http.Request) {
	k := req.FormValue("ip")
	rq := g_qqwry.Find(k)

	data, _ := json.Marshal(rq)
	w.Write(data)

}
func main() {
	g_qqwry = NewQQwry("./qqwry.dat")
	if g_qqwry != nil {
		http.HandleFunc("/", handlerIP)
		http.ListenAndServe("127.0.0.1:80", nil)
	}
}
