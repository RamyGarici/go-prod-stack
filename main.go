package main

import ("net/http" 
        "log"
	)




func main() {
   s := http.NewServeMux()
   s.HandleFunc("GET /health", func(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}"status":"ok"}`))
	log.Println("ok")
   }) 
   log.Println("Listening...")
   if err:=http.ListenAndServe("127.0.0.1:8080",s); err!=nil{
	log.Fatal(err)
   }

}