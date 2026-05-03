package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/ugorji/go/codec"
)

var mh codec.MsgpackHandle

func startRPCServer() {
	auth := new(Auth)
	rpc.Register(auth)

	// TODO get port from env var, but needs to be coordinated with the lua implementation
	log.Println("listening on :8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error listening: ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln("Error accepting connection: ", err)
		}

		rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, &mh)
		go rpc.ServeCodec(rpcCodec)
	}
}

func startCallbackServer(callback chan TickTickOAuthRes) {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:        ":9090",
		Handler:     mux,
		IdleTimeout: 5 * time.Minute,
		ReadTimeout: 5 * time.Minute,
	}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			log.Printf("no code found on the redirect_uri\n")

			callback <- TickTickOAuthRes{
				Err: ErrNoCodeProvided,
			}

			w.WriteHeader(http.StatusBadRequest)
		}

		state := r.URL.Query().Get("state")
		if state == "" {
			log.Printf("no state found on the redirect_uri\n")

			callback <- TickTickOAuthRes{
				Err: ErrStateMismatch,
			}

			w.WriteHeader(http.StatusBadRequest)
		}

		log.Printf("code: %s, state: %s\n", code, state)

		callback <- TickTickOAuthRes{
			Code: code,
		}

		res := []byte(`<html>
	<p>you can close this tab</p>
</html>`)

		w.Write(res)

		go func() {
			srv.Shutdown(context.Background())
		}()
	})

	log.Printf("callback server started on :9090\n")
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("error on callback server: %v\n", err)
	}
}
