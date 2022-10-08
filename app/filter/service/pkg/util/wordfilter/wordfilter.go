package wordfilter

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	level    int
	End      bool
}

type SenWord struct {
	Word  string
	Level int32
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Root = NewTrieNode()
	return t
}

func NewTrieNode() *TrieNode {
	n := &TrieNode{}
	n.Children = make(map[rune]*TrieNode)
	n.End = false

	return n
}

func (this *Trie) Add(txt string, level int) {
	if len(txt) < 1 {
		return
	}
	chars := []rune(txt)
	slen := len(chars)
	node := this.Root
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; !exists {
			node.Children[chars[i]] = NewTrieNode()
		}
		node = node.Children[chars[i]]
	}
	node.level = level
	node.End = true
}

//替换，不带返回敏感词
func (this *Trie) Replace(txt string) (string, bool) {
	chars := []rune(txt)
	result := []rune(txt)
	slen := len(chars)
	hasFind := false
	node := this.Root
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; exists {
			node = node.Children[chars[i]]
			for j := i + 1; j < slen; j++ {
				if _, exists := node.Children[chars[j]]; !exists {
					break
				}
				node = node.Children[chars[j]]
				if node.End == true {
					for t := i; t <= j; t++ {
						c, _ := utf8.DecodeRuneInString("*")
						result[t] = c
					}
					i = j
					hasFind = true
					node = this.Root
					break
				}
			}
			node = this.Root
		}
	}

	return string(result), hasFind
}

//检查是否有敏感词。返回
func (this *Trie) Check(txt string) (bool, []string) {
	if len(txt) <= 1 {
		return false, []string{}
	}
	chars := []rune(txt)
	slen := len(chars)
	node := this.Root
	foundMsg := false
	rtMsg := make([]string, 0)
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; exists {
			node = node.Children[chars[i]]
			for j := i + 1; j < slen; j++ {
				if _, exists := node.Children[chars[j]]; !exists {
					break
				}
				node = node.Children[chars[j]]
				if node.End == true {
					msg := chars[i : j+1]
					rtMsg = append(rtMsg, string(msg))
					i = j
					foundMsg = true
					node = this.Root
					break
				}
			}
			node = this.Root
		}
	}
	return foundMsg, rtMsg
}

//带返回敏感词
func (this *Trie) ReplaceEx(txt string) (string, bool, []*SenWord) {
	chars := []rune(txt)
	result := []rune(txt)
	slen := len(chars)
	hasFind := false
	node := this.Root
	//rtMsg := make([]string,0)
	//rtMsgLevel := make([]int,0)
	rtMsg := make([]*SenWord, 0)
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; exists {
			node = node.Children[chars[i]]
			for j := i + 1; j < slen; j++ {
				if _, exists := node.Children[chars[j]]; !exists {
					break
				}
				node = node.Children[chars[j]]
				if node.End == true {
					for t := i; t <= j; t++ {
						c, _ := utf8.DecodeRuneInString("*")
						result[t] = c
					}
					msg := chars[i : j+1]
					word := SenWord{}
					word.Word = string(msg)
					word.Level = int32(node.level)
					rtMsg = append(rtMsg, &word)
					i = j
					hasFind = true
					node = this.Root
					break
				}
			}
			node = this.Root
		}
	}

	return string(result), hasFind, rtMsg
}

func NewTrieWithFile(file string) *Trie {
	var r Trie
	r.Root = NewTrieNode()
	importWords(&r, file)
	return &r
}

//导入过滤词库
func importWords(t *Trie, file string) error {
	rd, err := os.Open(file)
	if err != nil {
		return err
	}

	r := bufio.NewReader(rd)
	for {
		line, isPrefix, e := r.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			continue
		}
		if word := strings.TrimSpace(string(line)); word != "" {
			s := strings.Split(word, " ")
			if len(s) < 2 {
				t.Add(word, 1)
			} else {
				t.AddTwo(s[1], s[0], 1)
			}
		}
	}
	rd.Close()
	return nil
}

//繁体 简体添加
func (this *Trie) AddTwo(sc, tc string, level int) {
	s := addTwo(sc, tc)
	for _, item := range s {
		this.Add(string(item), level)
	}
}

func _addTwo(s [][]rune, a, b []rune) [][]rune {
	if len(a) <= 0 {
		return s
	}
	var r [][]rune
	for _, t := range s {
		r = append(r, append(t, a[0]))
		r = append(r, append(t, b[0]))
	}
	return _addTwo(r, a[1:], b[1:])
}

func addTwo(a, b string) [][]rune {
	if len(a) != len(b) || len(a) == 1 {
		return [][]rune{[]rune(a), []rune(b)}
	}

	if a == b {
		return [][]rune{[]rune(a)}
	}
	chars := []rune(a)
	result := []rune(b)
	s1 := [][]rune{{chars[0]}}
	s2 := [][]rune{{result[0]}}

	s1 = _addTwo(s1, chars[1:], result[1:])
	s2 = _addTwo(s2, result[1:], chars[1:])

	return append(s1, s2...)
}
