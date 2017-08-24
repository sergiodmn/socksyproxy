package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"fmt"

	"github.com/elazarl/goproxy"
)

var caCert = cert()
var caKey = key()

func setCA(caCert, caKey []byte) error {
	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	return nil
}

func cert() []byte {
    c, err := ioutil.ReadFile("ca.pem")
    if err != nil {
        fmt.Print(err)
    }
    return (c)
}

func key() []byte {
    k, err := ioutil.ReadFile("ca.key.pem")
    if err != nil {
        fmt.Print(err)
    }
    return []byte(k)
}

