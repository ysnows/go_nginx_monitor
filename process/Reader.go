package process

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"time"
)

type Read interface {
	Read(Rc chan string)
}

type Reader struct {
	LogPath string
}

func (reader *Reader) Read(Rc chan []byte) {

	file, e := os.Open("./access.log")

	if e != nil {
		panic(fmt.Sprintf("hello%s", e.Error()))
	}

	file.Seek(0, 2)

	newReader := bufio.NewReader(file)

	for true {
		readBytes, i := newReader.ReadBytes('\n')

		if i == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if i != nil {
			panic(fmt.Sprintf("Helo"))
		}

		Rc <- readBytes
	}

}
