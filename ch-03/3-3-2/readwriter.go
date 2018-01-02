import (
	"bufio"
	"io"
)

var readwriter io.ReadWriter = bufio.NewReadWriter(reader, writer)