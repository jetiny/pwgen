package main

import (
	"bytes"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"github.com/atotto/clipboard"
	"runtime"
)

const (
	name     = "pwgen"
	version  = "0.0.3"
	revision = "HEAD"
)

var (
	numbers = []byte("0123456789")
	symbols = []byte("!\"#$%&'()=~|-^\\`{*}<>?_@[;:],./'")
)

func count(bs []byte, bp []byte) int {
	n := 0
	for _, b := range bs {
		if bytes.Index(bp, []byte{b}) != -1 {
			n++
		}
	}
	return n
}

func run(w io.Writer, o, n, nc, sc int, b bool) error {
	if o <= 0 || n < 1 {
		return flag.ErrHelp
	}
	nnc := 0
	if nc > 0 {
		nnc = nc
	}
	nsc := 0
	if sc > 0 {
		nsc = sc
	}
	if n < nnc+nsc {
		return errors.New("total length of numbers+symbols is greeter than maximum characters")
	}

	var buf bytes.Buffer
	if n > nnc+nsc {
		for r := 'a'; r < 'z'; r++ {
			buf.WriteRune(r)
		}
		for r := 'A'; r < 'Z'; r++ {
			buf.WriteRune(r)
		}
	}
	if nc != 0 && nsc != n {
		buf.Write(numbers)
	}
	if sc != 0 && nnc != n {
		buf.Write(symbols)
	}
	chars := buf.Bytes()
	no := 0
	nw := 0
	var pw bytes.Buffer
	dst := ""
	for {
		for {
			pw.Reset()
			for i := 0; i < n; i++ {
				r, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
				if err != nil {
					panic(err)
				}
				pw.WriteByte(chars[int(r.Int64())])
			}
			if nc > 0 && count(pw.Bytes(), numbers) < nc {
				continue
			}
			if sc > 0 && count(pw.Bytes(), symbols) < sc {
				continue
			}
			break
		}
		dst = dst + pw.String()
		//fmt.Fprint(w, pw.String())

		nw += n + 1
		if nw < 80-n {
			//fmt.Fprint(w, " ")
			dst = dst + " "
		} else {
			dst = dst + "\n"
			//fmt.Fprintln(w)
			nw = 0
		}
		no++
		if no == o {
			break
		}
	}
	//fmt.Fprintln(w)
	if b {
		fmt.Fprintln(w, dst)
		fmt.Fprintln(w, "password has been copied to the clipboard.")
		clipboard.WriteAll(dst)
	} else {
		fmt.Fprintln(w, dst)
	}
	return nil
}

func main() {
	var c int
	var n int
	var nc int
	var sc int
	var b bool
	var showVersion bool
	flag.IntVar(&c, "c", 1, "count of output, 密码个数")
	flag.IntVar(&n, "n", 8, "number of characters, 密码位数")
	flag.IntVar(&nc, "nc", -1, "minimum count of numbers (default: any), 数字个数")
	flag.IntVar(&sc, "sc", -1, "minimum count of symbols (default: any), 字符个数")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&b, "b", true, "copy to clipboard, 复制到剪切板")
	flag.Parse()

	if showVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	err := run(os.Stdout, c, n, nc, sc, b)
	if err != nil {
		if err == flag.ErrHelp {
			flag.Usage()
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(2)
	}
}
