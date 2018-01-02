// 全てコピー
writeSize, err := io.Copy(writer, reader)
// 指定したサイズだけコピー
writeSize, err := io.Copy(writer, reader, size)