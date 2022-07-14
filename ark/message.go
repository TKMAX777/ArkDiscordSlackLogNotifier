package ark

import (
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var illigalFormatError = errors.New("IlligalFormat")

var regexps = struct {
	JoinLeft *regexp.Regexp
	Tamed    *regexp.Regexp
	Killed   *regexp.Regexp
	Other    *regexp.Regexp
}{
	JoinLeft: regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\s*\d+\]\d+\.\d+\.\d+_\d+\.\d+\.\d+:\s(\S+)\s(joined|left)\sthis\sARK`),
	Tamed:    regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\s*\d+\].*: (\S+) of Tribe (.*) Tamed (a|an) (.+) - Lvl (\d+) \((.+)\)!`),
	Killed:   regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\s*\d+\].*: (\S+) - Lvl (\d+) \((.*)\) was killed by (a|an) (.+) - Lvl (\d+) \((.*)\)`),
	Other:    regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\s*\d+\](.*:\s*)?(.*)`),
}

func NewMessageFromLine(line string) (*Message, error) {
	var parseContent = func(line string) (*Message, error) {
		var contents = regexps.Other.FindAllStringSubmatch(line, 1)[0]

		t, err := time.Parse("2006.01.02-15.04.05", contents[1])
		if err != nil {
			return nil, errors.Wrapf(err, "TimeParseError(%s)", contents[1])
		}

		var message = Message{
			RawLine: line,
			Content: contents[3],
			Time:    t,
		}

		return &message, nil
	}

	switch {
	case regexps.JoinLeft.MatchString(line):
		var contents = regexps.JoinLeft.FindAllStringSubmatch(line, 1)[0]

		message, err := parseContent(line)
		if err != nil {
			return message, err
		}

		switch contents[3] {
		case "joined":
			message.Event = MessageTypeJoin{
				User: User{
					Name: contents[2],
				},
			}
		case "left":
			message.Event = MessageTypeLeave{
				User: User{
					Name: contents[2],
				},
			}
		}

		return message, nil
	case regexps.Tamed.MatchString(line):
		var contents = regexps.Tamed.FindAllStringSubmatch(line, 1)[0]
		message, err := parseContent(line)
		if err != nil {
			return message, err
		}

		krl, _ := strconv.Atoi(contents[5])

		message.Event = MessageTypeTamed{
			User: User{
				Name:  contents[2],
				Tribe: contents[3],
			},
			Tamed: Resident{
				Name:  contents[7],
				Level: krl,
			},
		}
		return message, nil
	case regexps.Killed.MatchString(line):
		var contents = regexps.Killed.FindAllStringSubmatch(line, 1)[0]
		message, err := parseContent(line)
		if err != nil {
			return message, err
		}

		kul, _ := strconv.Atoi(contents[3])
		krl, _ := strconv.Atoi(contents[7])

		message.Event = MessageTypeKilled{
			KilledUser: User{
				Name:  contents[2],
				Level: kul,
				Tribe: contents[4],
			},
			KilledBy: Resident{
				Name:  contents[6],
				Level: krl,
				Tribe: contents[8],
			},
		}
		return message, nil
	case regexps.Other.MatchString(line):
		message, err := parseContent(line)
		if err != nil {
			return message, err
		}

		message.Event = MessageTypeOther{}

		return message, nil
	default:
		return &Message{RawLine: line, Event: MessageTypeIlligalFormat{}}, illigalFormatError
	}
}
