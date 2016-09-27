package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	sv "github.com/siuyin/fchk"
)

func main() {
	client, err := rpc.DialHTTP("tcp", os.Getenv("FCHK_URL"))
	if err != nil {
		log.Fatal("dial:", err)
	}

	args := &sv.Args{os.Args[1], os.Getenv("FCHK_PATHS")}
	var reply []sv.FChk
	err = client.Call("FChk.Check", args, &reply)
	if err != nil {
		log.Fatal("fchk error:", err)
	}

	for i := range reply {
		fmt.Printf("%v: %v\n", reply[i].Found, reply[i].Path)
	}

}
