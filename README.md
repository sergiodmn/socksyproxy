# socksyproxy
HTTP proxy to redirect requests to a SOCKS proxy. 

Useful to solve SSL/TLS issues with Burp

Client/Burp -> SockyProxy -> Socks -> Internet

~~~
Usage of ./socksyproxy:
  -addr string
        Proxy listen address (default "127.0.0.1:8080")
  -cert
        Use custom CA files (ca.pem and ca.pem.key) in the same folder
  -socks string
        Socks proxy to connect
  -v    Should every proxy request be logged to stdout
~~~

You can use Burp's CA by exporting the Certificate and Private Key in DER format and converting it to PEM.
