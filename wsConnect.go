package leialib

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func WsConnect(url string, headers *http.Header, rootCApool *x509.CertPool) (ws *websocket.Conn, err error) {
	wsDialer := websocket.Dialer{
		HandshakeTimeout:  5 * time.Second,
		EnableCompression: true,
		TLSClientConfig: &tls.Config{
			RootCAs:            rootCApool,
			InsecureSkipVerify: false,
		},
	}
	ws, _, err = wsDialer.Dial(url, *headers)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
