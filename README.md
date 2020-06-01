<p xmlns:dct="http://purl.org/dc/terms/" xmlns:vcard="http://www.w3.org/2001/vcard-rdf/3.0#">
  <a rel="license"
     href="http://creativecommons.org/publicdomain/zero/1.0/">
    <img src="http://i.creativecommons.org/p/zero/1.0/88x31.png" style="border-style: none;" alt="CC0" />
  </a>
  <br />
  To the extent possible under law,
  <a rel="dct:publisher"
     href="https://github.com/xaionaro/">
    <span property="dct:title">Dmitrii Okunev</span></a>
  has waived all copyright and related or neighboring rights to
  "<span property="dct:title">A package to disable UDP fragmentation</span>.
This work is published from:
<span property="vcard:Country" datatype="dct:ISO3166"
      content="IE" about="https://github.com/xaionaro-go/udpnofrag">
  Ireland</span>".
</p>

# Quick start

```go
package main

import (
    "log"
    "net"

    "github.com/xaionaro-go/udpnofrag"
)

func main() {
    conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
        IP:net.ParseIP("192.168.0.1"),
        Port:443,
    })
    if err != nil {
        log.Fatal(err)
    }

    err = udpnofrag.UDPSetNoFragment(conn)
    if err != nil {
        log.Fatal(err)
    }
    
    b := make([]byte, 4096)
    _, err = conn.Write(b)
    if err != nil {
        log.Fatal(err)
    }
}
```

```sh
$ go run ./example.go 
2020/06/01 11:43:04 write udp 192.168.0.129:60695->192.168.0.1:443: write: message too long
exit status 1
```