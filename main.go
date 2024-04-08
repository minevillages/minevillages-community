package main

import (
	"log"
	"minevillages/community/api"
	baseapi "minevillages/community/api/base"
	"minevillages/community/database"
	"net/http"
	"runtime"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	// "github.com/quic-go/quic-go/http3"
)

func main() {
	// 프로세스의 모든 CPU 코어를 활용하도록 설정합니다.
	runtime.GOMAXPROCS(runtime.NumCPU())

	handler := api.HostRouteHandler{}
	handler.RegisterHost("127.0.0.1:8080", baseapi.HTTPHandler{})

	database.Postgres()

	if err := http.ListenAndServe(
		":8080",
		// "ssl/certificate.crt",
		// "ssl/private.key",
		&handler,
	); err != nil {
		log.Fatalln("HTTP 관련 수신 소켓을 초기화하는 과정에서 예외가 발생하였습니다.", err.Error())
	}
}
