package wordfilter

import (
	"testing"
)

var T *Trie

func init() {
	T = NewTrie()
	importWords(T, "alarm_sc.txt")
}

func TestCheck(t *testing.T) {
	tests := []struct {
		txt    string
		result bool
	}{
		{txt: "汇款", result: true},
		{txt: "你好", result: false},
	}
	for _, tt := range tests {
		v, _ := T.Check(tt.txt)
		t.Logf("Replace() txt = %v, filterTxt %v", tt.txt, v)
	}
}
