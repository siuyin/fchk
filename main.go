package fchk

import (
	"errors"
	"os"
	"path"
	"strings"
)

type Args struct {
	F, Paths string
}
type FChk struct {
	Path  string
	Found bool
}

func (f *FChk) Check(args *Args, reply *[]FChk) error {
	if args.F == "" || args.Paths == "" {
		return errors.New("filename and/or paths must not be empty")
	}

	ps := strings.Split(args.Paths, ",")
	for i := range ps {
		fc := FChk{}
		fc.Path = path.Join(ps[i], args.F)
		f, err := os.Open(fc.Path)
		if err == nil {
			f.Close()
			fc.Found = true
		} else {
			fc.Found = false
		}
		*reply = append(*reply, fc)
	}

	return nil
}
