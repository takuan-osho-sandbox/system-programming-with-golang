import "bytes"

// 空のバッファ
var buffer1 bytes.Buffer

// バイト列で初期化
buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})

// 文字列で初期化
buffer3 := bytes.NewBufferString("初期文字列")