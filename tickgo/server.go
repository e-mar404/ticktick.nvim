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
	rpc.Register(NewAuthService())

	log.Printf("listening on %s\n", RPC_PORT)
	listener, err := net.Listen("tcp", RPC_PORT)
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

func startCallbackServer(callback chan TickTickOAuthRes, state string) {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:        CALLBACK_PORT,
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

		callbackState := r.URL.Query().Get("state")
		if callbackState != state || callbackState == "" {
			log.Printf("incorrect state. Expected: %s, got: %s\n", state, callbackState)

			callback <- TickTickOAuthRes{
				Err: ErrStateMismatch,
			}

			w.WriteHeader(http.StatusBadRequest)
		}

		log.Printf("code: %s, state: %s\n", code, callbackState)

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

	log.Printf("callback server started on %s\n", CALLBACK_PORT)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Printf("error on callback server: %v\n", err)
		}
	}
}
