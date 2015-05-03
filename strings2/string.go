package strings2

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/cosiner/gohper/errors"
	"github.com/cosiner/gohper/index"
	"github.com/cosiner/gohper/unibyte"
)

const ErrQuoteNotMatch = errors.Err("Quote don't match")

// TrimQuote trim quote for string, return error if quote don't match
func TrimQuote(str string) (string, error) {
	str = strings.TrimSpace(str)
	if l := len(str); l > 0 {
		c := str[0]
		if c == '\'' || c == '"' || c == '`' {
			if str[l-1] == c {
				str = str[1 : l-1]
			} else {
				return "", ErrQuoteNotMatch
			}
		}
	}
	return str, nil
}

// TrimUpper return the trim and upper format of a string
func TrimUpper(str string) string {
	return strings.ToUpper(strings.TrimSpace(str))
}

// TrimLower return the trim and lower format of a string
func TrimLower(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// TrimSplit split string and return trim space string
func TrimSplit(s, sep string) []string {
	sp := strings.Split(s, sep)
	for i, n := 0, len(sp); i < n; i++ {
		sp[i] = strings.TrimSpace(sp[i])
	}
	return sp
}

// TrimAfter trim string and remove the section after delimiter and delimiter itself
func TrimAfter(s string, delimiter string) string {
	if idx := strings.Index(s, delimiter); idx >= 0 {
		s = s[:idx]
	}
	return strings.TrimSpace(s)
}

// SplitAtN find index of n-th sep string
func SplitAtN(str, sep string, n int) (index int) {
	index, idx, seplen := 0, -1, len(sep)
	for i := 0; i < n; i++ {
		if idx = strings.Index(str, sep); idx == -1 {
			break
		}
		str = str[idx+seplen:]
		index += idx
	}
	if idx == -1 {
		index = -1
	} else {
		index += (n - 1) * seplen
	}
	return
}

// SplitAtLastN find last index of n-th sep string
func SplitAtLastN(str, sep string, n int) (index int) {
	for i := 0; i < n; i++ {
		if index = strings.LastIndex(str, sep); index == -1 {
			break
		}
		str = str[:index]
	}
	return
}

// Seperate string by seperator, the seperator must in the middle of string,
// not first and last
func Seperate(s string, sep byte) (string, string) {
	if i := MidIndex(s, sep); i > 0 {
		return s[:i], s[i+1:]
	}
	return "", ""
}

func LastIndexByte(s string, b byte) int {
	for l := len(s) - 1; l >= 0; l-- {
		if s[l] == b {
			return l
		}
	}
	return -1
}

// IsAllCharsIn check whether all chars of string is in encoding string
func IsAllCharsIn(s, encoding string) bool {
	for i := 0; i < len(s); i++ {
		if index.CharIn(s[i], encoding) < 0 {
			return false
		}
	}
	return true
}

// MidIndex find middle seperator index of string, not first and last
func MidIndex(s string, sep byte) int {
	i := strings.IndexByte(s, sep)
	if i > 0 && i < len(s)-1 {
		return i
	}
	return -1
}

// RepeatJoin repeat s count times as a string slice, then join with sep
func RepeatJoin(s, sep string, count int) string {
	switch {
	case count <= 0:
		return ""
	case count == 1:
		return s
	case count == 2:
		return s + sep + s
	default:
		bs := make([]byte, 0, (len(s)+len(sep))*count-len(sep))
		buf := bytes.NewBuffer(bs)
		buf.WriteString(s)
		for i := 1; i < count; i++ {
			buf.WriteString(sep)
			buf.WriteString(s)
		}
		return buf.String()
	}
}

// SuffixJoin join string slice with suffix
func SuffixJoin(s []string, suffix, sep string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0] + suffix
	}
	n := len(sep) * (len(s) - 1)
	for i, sl := 0, len(suffix); i < len(s); i++ {
		n += len(s[i]) + sl
	}
	b := make([]byte, n)
	bp := copy(b, s[0])
	bp += copy(b[bp:], suffix)
	for _, s := range s[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
		bp += copy(b[bp:], suffix)
	}
	return string(b)
}

// JoinInt join int slice as string
func JoinInt(v []int, sep string) string {
	if len(v) > 0 {
		buf := bytes.NewBuffer([]byte{})
		buf.WriteString(strconv.Itoa(v[0]))
		for _, s := range v[1:] {
			buf.WriteString(sep)
			buf.WriteString(strconv.Itoa(s))
		}
		return buf.String()
	}
	return ""
}

// JoinInt join int slice as string
func JoinUint(v []uint, sep string) string {
	if len(v) > 0 {
		buf := bytes.NewBuffer([]byte{})
		buf.WriteString(strconv.FormatUint(uint64(v[0]), 10))
		for _, s := range v[1:] {
			buf.WriteString(sep)
			buf.WriteString(strconv.FormatUint(uint64(s), 10))
		}
		return buf.String()
	}
	return ""
}

// Compare compare two string, if equal, 0 was returned, if s1 > s2, 1 was returned,
// otherwise -1 was returned
func Compare(s1, s2 string) int {
	l1, l2 := len(s1), len(s2)
	for i := 0; i < l1 && i < l2; i++ {
		if s1[i] < s2[i] {
			return -1
		} else if s1[i] > s2[i] {
			return 1
		}
	}
	switch {
	case l1 < l2:
		return -1
	case l1 == l2:
		return 0
	default:
		return 1
	}
}

// RemoveSpace remove all space characters from string by unibyte.IsSpace
func RemoveSpace(s string) string {
	idx, end := 0, len(s)
	bs := make([]byte, end)
	for i := 0; i < end; i++ {
		if !unibyte.IsSpace(s[i]) {
			bs[idx] = s[i]
			idx++
		}
	}
	return string(bs[:idx])
}
