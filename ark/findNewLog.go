package ark

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func (h Handler) findNewLog() (string, error) {
	fs, err := ioutil.ReadDir(h.dirpath)
	if err != nil {
		return "", errors.Wrap(err, "ReadDir")
	}

	var newestLog string
	var nLogNo int
	for _, f := range fs {
		if !strings.HasPrefix(f.Name(), "ServerGame.") {
			continue
		}

		logNo, err := strconv.Atoi(strings.Split(strings.TrimSuffix(f.Name(), "ServerGame."), ".")[0])
		if err != nil {
			continue
		}

		if logNo > nLogNo {
			nLogNo = logNo
			newestLog = f.Name()
		}
	}

	if nLogNo == 0 {
		return "", errors.New("NotFound")
	}

	return newestLog, nil
}
