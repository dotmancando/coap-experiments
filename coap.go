package main

import (
	"log"
	"net"
	"time"

	coap "github.com/dustin/go-coap"
)

func pulse(conn *net.UDPConn, addr *net.UDPAddr, msg *coap.Message) *coap.Message {
	if msg.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: msg.MessageID,
			Token:     msg.Token,
			Payload:   []byte("ok"),
		}

		res.SetOption(coap.ContentFormat, coap.TextPlain)
		return res
	}

	return nil
}

func echo(conn *net.UDPConn, addr *net.UDPAddr, msg *coap.Message) *coap.Message {
	if msg.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: msg.MessageID,
			Token:     msg.Token,
			Payload:   msg.Payload,
		}

		res.SetOption(coap.ContentFormat, coap.TextPlain)
		return res
	}

	return nil
}

func loggingMiddleware(next coap.Handler) coap.Handler {
	return coap.FuncHandler(func(conn *net.UDPConn, addr *net.UDPAddr, msg *coap.Message) *coap.Message {
		startTime := time.Now()
		defer func() {
			log.Printf("[%s] /%s. Completed in %v", msg.Code, msg.PathString(), time.Now().Sub(startTime))
		}()

		resp := next.ServeCOAP(conn, addr, msg)
		return resp
	})
}

func main() {
	mux := coap.NewServeMux()
	mux.Handle("/pulse", loggingMiddleware(coap.FuncHandler(pulse)))
	mux.Handle("/echo", loggingMiddleware(coap.FuncHandler(echo)))

	port := ":5683"

	log.Printf("Starting COAP server listening on %s\n", port)
	coap.ListenAndServe("udp", port, mux)
}
