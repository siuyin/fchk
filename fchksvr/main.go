package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	sv "github.com/siuyin/fchk"
)

func main() {
	fc := new(sv.FChk)
	rpc.Register(fc)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+os.Getenv("FCHK_PORT"))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()
	http.Serve(l, nil)

}
