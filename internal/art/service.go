package art

import (
	"fmt"
	"strconv"
	"strings"
)

type Mode string

const (
	ModeDecode Mode = "decode"
	ModeEncode Mode = "encode"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Execute(mode Mode, input string, multiline bool) (string, error) {
	if strings.Contains(input, "[") && strings.Count(input, "[") != strings.Count(input, "]") {
		return "", fmt.Errorf("unbalanced brackets")
	}

	if multiline {
		lines := strings.Split(input, "\\n")
		out := make([]string, 0, len(lines))
		for _, line := range lines {
			res, err := s.executeSingle(mode, line)
			if err != nil {
				return "", err
			}
			out = append(out, res)
		}
		return strings.Join(out, "\n"), nil
	}

	return s.executeSingle(mode, input)
}

func (s *Service) executeSingle(mode Mode, input string) (string, error) {
	switch mode {
	case ModeDecode:
		return decode(input)
	case ModeEncode:
		return encode(input), nil
	default:
		return "", fmt.Errorf("unsupported mode")
	}
}

func decode(input string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(input); {
		if input[i] != '[' {
			b.WriteByte(input[i])
			i++
			continue
		}

		end := strings.IndexByte(input[i:], ']')
		if end == -1 {
			return "", fmt.Errorf("unbalanced brackets")
		}
		end += i

		content := input[i+1 : end]
		sep := strings.IndexByte(content, ' ')
		if sep == -1 {
			return "", fmt.Errorf("missing separator")
		}

		countStr := content[:sep]
		repeat := content[sep+1:]
		if repeat == "" {
			return "", fmt.Errorf("empty repeat")
		}

		count, err := strconv.Atoi(countStr)
		if err != nil || count < 0 {
			return "", fmt.Errorf("invalid count")
		}

		for j := 0; j < count; j++ {
			b.WriteString(repeat)
		}
		i = end + 1
	}
	return b.String(), nil
}

func encode(input string) string {
	if input == "" {
		return ""
	}

	var b strings.Builder
	runes := []rune(input)
	for i := 0; i < len(runes); {
		j := i + 1
		for j < len(runes) && runes[j] == runes[i] {
			j++
		}

		runLen := j - i
		seg := string(runes[i])
		if runLen > 1 || seg == "[" || seg == "]" {
			b.WriteString("[")
			b.WriteString(strconv.Itoa(runLen))
			b.WriteString(" ")
			b.WriteString(seg)
			b.WriteString("]")
		} else {
			b.WriteString(seg)
		}
		i = j
	}

	return b.String()
}
