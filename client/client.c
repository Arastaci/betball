#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>

#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 1337 

int main() {
    int sock;
    struct sockaddr_in server;
    char message[2000], server_reply[2000];

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1) {
        printf("Soket oluşturulamadı\n");
        return 1;
    }
    printf("Soket oluşturuldu\n");

    server.sin_addr.s_addr = inet_addr(SERVER_IP);
    server.sin_family = AF_INET;
    server.sin_port = htons(SERVER_PORT);

    if (connect(sock, (struct sockaddr *)&server, sizeof(server)) < 0) {
        perror("Bağlantı hatası");
        return 1;
    }
    printf("Sunucuya bağlandı\n");

    while (1) {
        printf("Bahis yapacağınız takımı yazın (GS veya FB) [Yardım için 'help' yazın]: ");
        scanf("%s", message);

        if (strcmp(message, "help") == 0) {
            printf("Mevcut komutlar:\n");
            printf("GS: Galatasaray'a oy ver\n");
            printf("FB: Fenerbahçe'ye oy ver\n"); 
            printf("status: Mevcut oy oranlarını göster\n");
            continue;
        }
        if (send(sock, message, strlen(message), 0) < 0) {
            printf("Mesaj gönderme hatası\n");
            return 1;
        }
        if (recv(sock, server_reply, 2000, 0) < 0) {
            printf("Sunucudan yanıt alınamadı\n");
            return 1;
        }
        server_reply[strcspn(server_reply, "\n")] = '\0'; 
        printf("Güncel durum: %s\n", server_reply);
    }

    close(sock);
    return 0;
}
