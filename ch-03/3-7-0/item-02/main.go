package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)
	// けどバッファを残ってる
	fmt.Println(buffer.String())
}
