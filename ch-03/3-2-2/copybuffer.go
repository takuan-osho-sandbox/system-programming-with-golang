// 8KBのバッファを使う
buffer := make([]byte, 8*1024)
io.CopyBuffer(writer, reader, buffer)