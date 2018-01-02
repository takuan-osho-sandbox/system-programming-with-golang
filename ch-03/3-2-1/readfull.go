buffer := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)