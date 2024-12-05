# BetBall - Futbol Bahis Simülasyonu

BetBall, eğitim amaçlı geliştirilmiş basit bir futbol bahis simülasyon sistemidir. Go ile yazılmış bir server ve C ile yazılmış bir client uygulamasından oluşur. Kullanıcılar Galatasaray (GS) ve Fenerbahçe (FB) takımları üzerine sanal bahisler yapabilir ve sonuçları gerçek zamanlı olarak takip edebilirler.

## Özellikler

- Gerçek zamanlı bahis yapma ve sonuç takibi
- Takım bazlı oy/bahis sistemi
- Anlık oran güncelleme
- Güvenlik kontrolleri ve tehlikeli istek engelleme
- Loglama sistemi

## Başlangıç

### Gereksinimler

- Go 1.16 veya üzeri
- GCC veya uyumlu bir C derleyici
- Linux/Unix tabanlı işletim sistemi

### Kurulum

1. Projeyi klonlayın:

```bash
git clone [https://github.com/Arastaci/betball.git]
cd betball
```

2. Server'ı başlatın:

```bash
cd server
go run .
```

3. Client'ı derleyin ve çalıştırın:

```bash
cd client
gcc -o client client.c
./client
```

## Kullanım

Client uygulaması başlatıldığında aşağıdaki komutları kullanabilirsiniz:

- `GS` - Galatasaray'a oy/bahis ver
- `FB` - Fenerbahçe'ye oy/bahis ver
- `status` - Mevcut oy oranlarını görüntüle
- `help` - Komut listesini görüntüle
    