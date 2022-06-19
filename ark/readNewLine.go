package ark

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

func (h *Handler) readNewLine(lpath string) (line string, err error) {
	const bufSize = 500

	file, err := os.Open(lpath)
	if err != nil {
		return "", errors.Wrap(err, "OpenAuthFile")
	}
	defer file.Close()

	file.Seek(-1*(bufSize+1), 2)

	var scanner = bufio.NewScanner(file)

	// get last line
	for scanner.Scan() {
		line = scanner.Text()
	}

	return line, nil
}
