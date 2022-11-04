package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/api/v1/test", func(http.ResponseWriter, *http.Request){
		log.Println("Done");
	});
	http.ListenAndServe(":8080", nil);
}