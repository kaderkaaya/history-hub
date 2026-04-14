package main

import (
	"context"
	"fmt"
	"log"      //Logger kütüphanesi yani conloglar için bu kütüpü import ederiz.
	"net/http" //web server kütüp
	"time"

	"history-hub/internal/cache"
	config "history-hub/internal/config"
	apphttp "history-hub/internal/http"
	handlers "history-hub/internal/http/handlers"
	provider "history-hub/internal/provider/wikimedia"
	service "history-hub/internal/service"
)

func main() {
	cfg := config.Load()
	redisAddr := cfg.RedisHost + ":" + cfg.RedisPort
	redisClient := cache.NewRedisClient(redisAddr, cfg.RedisPassword, 0)
	if err := redisClient.Ping(context.Background()); err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}
	fmt.Printf("cfg.WikimediaBaseURL", cfg.WikimediaBaseURL)
	client := provider.NewClient(cfg.WikimediaBaseURL, 10*time.Second, "history-hub")
	eventService := service.NewEventsService(client, redisClient, cfg.CacheTTLTodayH, cfg.CacheTTLPastH)
	eventHandler := handlers.NewEventsHandler(eventService)
	router := apphttp.HistoryHubRouter(eventHandler)
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
