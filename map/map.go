package main
import (
	"fmt"
	"flag"
)

type ServerConfig struct {
	env string
	port int
} 

func main() {
	var serverConfig ServerConfig
	flag.IntVar(&serverConfig.port, "p", 8000, "[SERVER PORT]")
	flag.StringVar(&serverConfig.env, "e", "localhost", "Env")
	flag.Parse()


	fmt.Printf("%+v\n", serverConfig)
}
