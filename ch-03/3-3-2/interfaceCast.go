import (
	"io"
	"io/ioutil"
	"strings"
)

var reader io.Reader = strings.NewReader("テストデータ")
var readerClose io.ReadCloser = ioutil.NopCloser(reader)