package term

import (
	"bytes"
	"errors"

	"github.com/wangyuntao/terminfo"
)

type Key rune

const (
	KeyCtrlA Key = 1 + iota
	KeyCtrlB
	KeyCtrlC
	KeyCtrlD
	KeyCtrlE
	KeyCtrlF
	KeyCtrlG
	KeyCtrlH
	KeyCtrlI
	KeyCtrlJ
	KeyCtrlK
	KeyCtrlL
	KeyCtrlM
	KeyCtrlN
	KeyCtrlO
	KeyCtrlP
	KeyCtrlQ
	KeyCtrlR
	KeyCtrlS
	KeyCtrlT
	KeyCtrlU
	KeyCtrlV
	KeyCtrlW
	KeyCtrlX
	KeyCtrlY
	KeyCtrlZ
)

const (
	KeyF1 Key = -1 - iota
	KeyF2
	// TODO ...
)

var (
	keys = make([][]byte, 0)
)

func initKey() error {
	ti, err := terminfo.LoadEnv()
	if err != nil {
		return err
	}

	err = addKey(KeyF1, terminfo.KeyF1, ti)
	if err != nil {
		return err
	}

	err = addKey(KeyF2, terminfo.KeyF2, ti)
	if err != nil {
		return err
	}

	return nil
}

func addKey(k Key, tik int, ti *terminfo.Terminfo) error {
	s, _ := ti.GetStr(tik)
	if len(s) == 0 || s[0] != '\x1b' { // TODO What if the key is not supported by some terminal?
		return errors.New("key: illegal escape sequence")
	}
	keys = append(keys, s)
	return nil
}

func decodeKey(bf []byte) (Key, int, bool) {
	b := bf[0]
	if b == '\x1b' {
		for i, key := range keys {
			if bytes.HasPrefix(bf, key) {
				return Key(-i - 1), len(key), true
			}
		}
		return 0, 0, false
	}
	if k := Key(b); k >= KeyCtrlA && k <= KeyCtrlZ {
		return k, 1, true
	}
	return 0, 0, false
}
