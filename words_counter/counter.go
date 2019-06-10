package main

import (
	"fmt"
	io "io/ioutil"
	"runtime"

	//"runtime"
	str "strings"
)

var workers = runtime.NumCPU()

//var workers = 2

type ocorrency struct {
	word string
	qnt  int
}

//https://gist.github.com/xlab/6e204ef96b4433a697b3
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func worker(file string, init byte) {
	bk, err := io.ReadFile(file)
	check(err)
	z := split(bk, 1000)

	for _, x := range z {
		fmt.Println(string(x))
	}
}
func compare2(x byte) bool {
	if x >= 65 && x <= 122 {
		return true
	}
	return false
}

///AO INVES DE CRIAR UMA SLICE, INSERE DIRETAMENTE NO ARQUIVO .TXT
func thisTam(file string) int {
	bk, err := io.ReadFile(file)
	check(err)
	a := 0
	for _, x := range bk {
		if x >= 65 && x <= 122 {
			a++
		}
	}
	return a
}

//PRÉ PROC
func clear(file string) /*[]string*/ {
	bk, err := io.ReadFile(file)
	check(err)
	a2 := thisTam(file)
	d := make([]string, a2)
	for i := 0; i < a2; i++ {

		if compare2(bk[i]) {
			s := str.ToLower(string(bk[i]))
			d[i] = s
		}

	}
	fmt.Println(d)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func low(words []string) {

}

func show(ocorr []ocorrency) {
	for _, oc := range ocorr {
		word := oc.word
		qnt := oc.qnt
		fmt.Println("[ Name: ", word, ", quantity:", qnt, "]")
	}
}

func words(file string) []string {
	bk, err := io.ReadFile(file)
	check(err)
	ws := str.Split(string(bk), " ")
	return ws
}

func verify(w string, oc []ocorrency) (bool, int) {
	r := false
	index := 0
	for i, s := range oc {
		if str.Compare(w, s.word) == 0 {
			r = true
			index = i
		}
	}
	return r, index
}

func result(bk []string) []ocorrency {
	var ocorrencys []ocorrency
	for _, w := range bk {
		b, i := verify(w, ocorrencys)
		if !b {
			ocorrencys = append(ocorrencys, ocorrency{w, 1})
		} else {
			ocorrencys[i].qnt++
		}
	}
	return ocorrencys
}

func main() {
	//ws := words("text2.txt")
	//rs := result(ws)
	//show(rs)
	//a := clear("text2.txt")
	clear("text2.txt")
	//fmt.Println(list)
}
