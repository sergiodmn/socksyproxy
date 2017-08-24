package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"os"
	netproxy "golang.org/x/net/proxy"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "Should every proxy request be logged to stdout")
	addr := flag.String("addr", "127.0.0.1:8080", "Proxy listen address")
	customcert := flag.Bool("cert", false, "Use custom CA files (ca.pem and ca.pem.key) in the same folder")
	socks := flag.String("socks", "", "Socks proxy to connect")
	flag.Parse()
        if *socks == ""{
                fmt.Fprintln(os.Stdout, "SOCKS Proxy not set")
                os.Exit(1)
        }
	if *customcert {
		setCA(caCert, caKey)
	}
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(
    		func(req *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			dialer, err := netproxy.SOCKS5("tcp", *socks, nil, netproxy.Direct)
			if err != nil {
                                fmt.Fprintln(os.Stdout, "Error connencting SOCKS proxy::", err)
			}

			transport := &http.Transport{Dial: dialer.Dial}
			client := &http.Client{Transport: transport}
			req.RequestURI = ""
			resp, err := client.Do(req)
			if err != nil {
				fmt.Fprintln(os.Stdout, "Can't GET page:", err)
			}
	        	return nil, resp
	    })
	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
