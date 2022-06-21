package ark

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (h Handler) findNewLog() (string, error) {
	fs, err := ioutil.ReadDir(h.dirpath)
	if err != nil {
		return "", errors.Wrap(err, "ReadDir")
	}

	var newestLog string
	var nLogNo int64
	for _, f := range fs {
		if !strings.HasPrefix(f.Name(), "ServerGame.") {
			continue
		}

		var fsep = strings.Split(f.Name(), ".")
		if len(fsep) < 2 {
			continue
		}

		t, err := time.Parse("2006.01.02_03.04.05.log", strings.Join(fsep[2:], ","))
		if err != nil {
			continue
		}

		if nLogNo < t.UnixMilli() {
			nLogNo = t.UnixMilli()
			newestLog = f.Name()
		}
	}

	if nLogNo == 0 {
		return "", errors.New("NotFound")
	}

	return newestLog, nil
}
