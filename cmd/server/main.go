package main

import (
	"log"      //Logger kütüphanesi yani conloglar için bu kütüpü import ederiz.
	"net/http" //web server kütüp

	apphttp "history-hub/internal/http"
)

func main() {
	router := apphttp.HistoryHubRouter()
	log.Println("server started on :8080")
	//Aslında bütün programın basladıgı yer burası index.js gibi.
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("api is running")) // burda []byte stringi byte çeviriyor neden?
	//burda w.Write string kabul etmez sadece byte slice kabul ediyormus ondan dolayı
	//http byte bazlı çalışıyormus ve performans açısından daha iyi olduğu için w.Write kullanıyoruz.
	//productionda bunu kullanıyoruz.
	//string = read-only
	//[]byte = read-write
	//})

	//log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil { //servere baslatıyoruz, burda default router kullanmısız
		log.Fatal(err) //eğer server baslamazsa hatayı döner ve programı kapatır.
	}
}
