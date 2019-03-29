package example_server

/**
 * 基于Web服务器监听
 */
import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
