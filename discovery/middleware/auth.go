package middleware

import (
	"log"
	"net/http"
	"twirp/define"
)

func BaseAuth(base http.Handler, token string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("auth")
		if auth != token {
			log.Printf("BaseAuth 401 Addr: %s \n", request.RemoteAddr)
			writer.WriteHeader(401)
			writer.Write([]byte(define.AuthError.Error()))
			return
		}
		base.ServeHTTP(writer, request)
	})
}
