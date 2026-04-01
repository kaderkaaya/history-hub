package wikimedia

import (
	"history-hub/internal/config"
	"net/http"
	"time"
)

type Client struct {
	WikimediaBaseURL string
	HTTPClient       *http.Client
	UserAgent        string
}

func NewClient(baseURL string, timeout time.Duration, userAgent string) *Client {
	return &Client{
		WikimediaBaseURL: config.Load().WikimediaBaseURL,
		UserAgent:        "history-hub", //bazi apiler isteğin kimden geldiğine bkar bundan dolayı bunu ekliyoruz.
		//eğer bir saldırı olursa engeller ama bunu eklediğimizde bizden geldiği belli olur.
		HTTPClient: &http.Client{Timeout: timeout},
		//burda request attığımızda api cevap vermezse sürekli bekler ama
		//timeout verdiğimizde timeout ile sistemi durdurur.
		//böylece server kilitlenmez ve kullanıcı beklemez.
		//tekrar kullanılır
		//tek yerden yönetilir
		// tüm requestler aynı ayarla çalışır
		//http.Client içindeki Timeout alanına fonksiyondan gelen timeout değerini ata
	}
}
