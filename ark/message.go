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
	Killed   *regexp.Regexp
	Other    *regexp.Regexp
}{
	// [2022.06.19-07.33.44:247][542]2022.06.19_07.33.44: tkmax777 joined this ARK!
	JoinLeft: regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\d+\]\d+\.\d+\.\d+_\d+\.\d+\.\d+:\s(\S+)\s(joined|left)\sthis\sARK`),
	// tkmax777 - Lvl 5 () was killed by a Dilophosaur - Lvl 4 ()!
	Killed: regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\d+\].*: (\S+) - Lvl (\d+) \((.*)\) was killed by (a|an) (.+) - Lvl (\d+) \((.*)\)`),
	Other:  regexp.MustCompile(`\[(\d+\.\d+\.\d+-\d+\.\d+\.\d+):\d+\]\[\d+\](.*:\s*)?(.*)`),
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
	case regexps.Killed.MatchString(line):
		var contents = regexps.Killed.FindAllStringSubmatch(line, 1)[0]
		message, err := parseContent(line)
		if err != nil {
			return message, err
		}

		kul, _ := strconv.Atoi(contents[3])
		kel, _ := strconv.Atoi(contents[7])

		message.Event = MessageTypeKilled{
			KilledUser: User{
				Name:  contents[2],
				Level: kul,
				Tribe: contents[4],
			},
			KilledBy: Enemy{
				Name:  contents[6],
				Level: kel,
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
