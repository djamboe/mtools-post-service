package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net"
	"net/http"
	"os"

	"github.com/djamboe/mtools-post-service/pkg/cmd"
)

func main() {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
