package main

import (
	"context"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/ugorji/go/codec"
)

var (
	mh          codec.MsgpackHandle
	rpcServices = make(map[string]any)
)

func startRPCServer() {
	server := rpc.NewServer()

	for name, service := range rpcServices {
		l.Debug("registered service on rpc server", "name", name)
		server.Register(service)
	}

	l.Info("starting server")
	listener, err := net.Listen("tcp", RPC_PORT)
	if err != nil {
		l.Fatalf("could not listen to port: %v\n", err)
	}
	defer listener.Close()

	l.Info("open to connections", "port", RPC_PORT)
	for {
		conn, err := listener.Accept()

		if err != nil {
			l.Fatalf("Error accepting connection: %v\n", err)
		}

		rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, &mh)
		go server.ServeCodec(rpcCodec)
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
		l.Debug("handling /callback request")

		code := r.URL.Query().Get("code")
		if code == "" {
			l.Error("%v\n", ErrNoCodeProvided)

			callback <- TickTickOAuthRes{
				Err: ErrNoCodeProvided,
			}

			w.WriteHeader(http.StatusBadRequest)
		}

		callbackState := r.URL.Query().Get("state")
		if callbackState != state || callbackState == "" {
			l.Errorf("incorrect state. Expected: %s, got: %s\n", state, callbackState)

			callback <- TickTickOAuthRes{
				Err: ErrStateMismatch,
			}

			w.WriteHeader(http.StatusBadRequest)
		}

		l.Debug("callback params found", "code", code, "state", state)

		callback <- TickTickOAuthRes{
			Code: code,
		}

		res := []byte(`<html>
	<p>you can close this tab</p>
</html>`)

		w.Write(res)

		go func() {
			l.Info("shutting down callback server")
			srv.Shutdown(context.Background())
		}()
	})

	l.Info("callback server started", "port", CALLBACK_PORT)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			l.Printf("error on callback server: %v\n", err)
		}
	}
}
