package main

import (
	"github.com/otpservice/otp"
	"net/http"
	"os"
	"github.com/go-kit/kit/log"
)


func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger,"listen", "8081", "caller", log.DefaultCaller)
	// use redis DB 0 for development and DB 1 for testing
	svc := otp.NewService(0)
	r := otp.NewHttpServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", "8081")
    logger.Log("err", http.ListenAndServe(":8081", r))
}

