package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
    "os"
    "time"
)

var gsPercentage = 50
var fbPercentage = 50
var mu sync.Mutex
var clientCount int 

// TODO: Implement security tests
var forbiddenMessages = []string{".", "/", "'", ">", "<", "`"}

var dangerousMessageCounts = make(map[string]int)
var dangerousMessageMu sync.Mutex 


var clientBets = make(map[string]*Bet)
var clientBetsMu sync.Mutex

type Bet struct {
    Amount int
    Team   string
}

func writeLog(clientAddr string, action string, result string) {
    timeStamp := time.Now().Format("2006-01-02 15:04:05")
    logMessage := fmt.Sprintf("[%s] Client: %s, Action: %s, Result: %s\n", 
        timeStamp, clientAddr, action, result)
    
    file, err := os.OpenFile("../logs/server_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Log dosyası açılamadı:", err)
        return
    }
    defer file.Close()
    
    if _, err := file.WriteString(logMessage); err != nil {
        log.Println("Log yazılamadı:", err)
    }
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
    clientAddr := conn.RemoteAddr().String()
	clientCount++
	fmt.Printf("Yeni bir client bağlandı! Toplam client: %d\n", clientCount)
    writeLog(clientAddr, "CONNECT", fmt.Sprintf("Total clients: %d", clientCount))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			clientCount-- 
			log.Println("Bağlantı hatası:", err)
			fmt.Printf("Client ayrıldı! Mevcut client sayısı: %d\n", clientCount)
			
			dangerousMessageMu.Lock()
			delete(dangerousMessageCounts, clientAddr)
			dangerousMessageMu.Unlock()
			return
		}
		message := strings.TrimSpace(string(buf[:n]))
        
        response := ""
        dangerousFound := false
        for _, forbidden := range forbiddenMessages {
			if strings.Contains(strings.ToLower(message), forbidden) {
				dangerousFound = true
				fmt.Println("Uyarı! Tehlikeli bir mesaj algılandı:", message)
				dangerousMessageMu.Lock()
				dangerousMessageCounts[clientAddr]++
				count := dangerousMessageCounts[clientAddr]
				dangerousMessageMu.Unlock()
				
				writeLog(clientAddr, "DANGEROUS_MESSAGE", fmt.Sprintf("Message: %s, Count: %d", message, count))
				
				if count > 5 {
					writeLog(clientAddr, "KICKED", "Too many dangerous messages")
					clientCount--
					return
				}
				break
			}
		}

		if !dangerousFound {
			mu.Lock()
			parts := strings.Split(message, " ")
			
			switch parts[0] {
			case "GS", "FB", "status":
				if message == "GS" {
					gsPercentage += 5
					fbPercentage -= 5
					response = fmt.Sprintf("GS: %d%%, FB: %d%%", gsPercentage, fbPercentage)
					writeLog(clientAddr, "GS", response)
				} else if message == "FB" {
					fbPercentage += 5
					gsPercentage -= 5
					response = fmt.Sprintf("GS: %d%%, FB: %d%%", gsPercentage, fbPercentage)
					writeLog(clientAddr, "FB", response)
				} else if message == "status" {
					response = fmt.Sprintf("GS: %d%%, FB: %d%%", gsPercentage, fbPercentage)
					writeLog(clientAddr, "STATUS", response)
				}
			} 
			mu.Unlock() 
		}
		fmt.Println("Güncel oranlar:", response)
		conn.Write([]byte(response + "\n"))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":1337") 
	if err != nil {
		log.Fatal("Server başlatılamadı:", err)
	}
	defer ln.Close()

	fmt.Println("Server başlatıldı. Clientlar bekleniyor...")
    
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Bağlantı kabul edilemedi:", err)
			continue
		}
		go handleConnection(conn)
	}
}

