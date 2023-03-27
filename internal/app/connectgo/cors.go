package connectgo

import (
	"net/http"

	"github.com/rs/cors"
)

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{http.MethodPost},
		AllowedOrigins: []string{
			"http://localhost:4173",
			"http://localhost:5173",
			"https://neutral-diet.firebaseapp.com",
			"https://neutral-diet.web.app",
			"https://neutral-diet-*.web.app",
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
	})
}
