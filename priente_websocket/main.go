package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket bağlantısı için ayarlamalar.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// CLIENT (İstemci) Bağlantılarını ve İsteklerini Yakalar.
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var mutex = &sync.Mutex{}

// CLIENT (İstemci) Bağlantıları tarafından gönderilen mesajların diğer istemciler ile paylaşılma yapısı.
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// WebSocket Bağlantısını Yöneten Fonksiyon
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("WebSocket bağlantısı yükseltilemedi: %v", err)
	}
	defer ws.Close()

	// Yeni CLIENT (İstemci) Ekler.
	clients[ws] = true
	mutex.Unlock()

	for {
		var msg Message
		// Yeni Mesajı Okur.
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Hata okuma: %v", err)
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}
		// Mesajı broadcast Kanalına Gönderir.
		broadcast <- msg
	}
}

// Mesajları Yayınlayan Fonksiyon
func handleMessages() {
	for {
		// Yeni Mesajı Alır.
		msg := <-broadcast
		// Tüm İstemcilere Mesajı Gönderir.
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Hata: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("HTTP sunucusu :8080 portunda başlatıldı")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
