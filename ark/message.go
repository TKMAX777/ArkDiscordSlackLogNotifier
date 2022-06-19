package ark

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

type Message struct {
	RawLine string
	Time    time.Time
	Content interface{}
}

var illigalFormatError = errors.New("IlligalFormat")

var regexps = struct {
	JoinLeft *regexp.Regexp
	Other    *regexp.Regexp
}{
	// [2022.06.19-07.33.44:247][542]2022.06.19_07.33.44: tkmax777 joined this ARK!
	JoinLeft: regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\d+\]\d+\.\d+\.\d+_\d+\.\d+\.\d+:\s(\S+)\s(joined|left)\sthis\sARK`),
	Other:    regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\d+\](.*)`),
}

func NewMessageFromLine(line string) (*Message, error) {
	var message = Message{
		RawLine: line,
	}

	var parseContent = func(line string) (string, *time.Time, error) {
		var contents = regexps.Other.FindAllStringSubmatch(line, 1)[0]

		t, err := time.Parse("2006.01.02-15.04.05", contents[1])
		if err != nil {
			return "", nil, errors.Wrapf(err, "TimeParseError(%s)", contents[1])
		}

		return contents[2], &t, nil
	}

	switch {
	case regexps.JoinLeft.MatchString(line):
		var contents = regexps.JoinLeft.FindAllStringSubmatch(line, 1)[0]

		c, t, err := parseContent(line)
		if err != nil {
			return &message, err
		}

		message.Time = *t

		switch contents[3] {
		case "joined":
			var join = MessageTypeJoin{
				UserName: contents[2],
				Content:  c,
			}

			message.Content = join
		case "left":
			var left = MessageTypeLeave{
				UserName: contents[2],
				Content:  c,
			}

			message.Content = left
		}

		return &message, nil

	case regexps.Other.MatchString(line):
		c, t, err := parseContent(line)
		if err != nil {
			return &message, err
		}

		message.Content = MessageTypeOther{
			Content: c,
		}
		message.Time = *t

		return &message, nil
	default:
		return &message, illigalFormatError
	}
}
