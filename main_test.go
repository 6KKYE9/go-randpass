package main

import (
	"strings"
	"testing"
)

func TestGenPasswordsLength(t *testing.T) {
	p, err := genPasswords(5, 12, lower+upper+digits)
	if err != nil {
		t.Fatal(err)
	}
	if len(p) != 5 {
		t.Fatalf("期望 5 个, 得到 %d", len(p))
	}
	for _, s := range p {
		if len(s) != 12 {
			t.Errorf("长度不对: %q", s)
		}
	}
}

func TestGenPasswordsCharset(t *testing.T) {
	// 只开数字，生成的串里不该出现字母
	p, err := genPasswords(3, 8, digits)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range p {
		for _, r := range s {
			if r < '0' || r > '9' {
				t.Errorf("字符池只含数字却出现了 %q", r)
			}
		}
	}
}

func TestEmptyPool(t *testing.T) {
	if _, err := genPasswords(1, 4, ""); err == nil {
		t.Error("空字符池应该报错")
	}
}

func TestUniqueish(t *testing.T) {
	// 不是严格唯一保证，但 100 个里大概率不全撞，顺手验证下分布
	p, _ := genPasswords(100, 10, lower)
	seen := map[string]bool{}
	for _, s := range p {
		seen[s] = true
	}
	if len(seen) < 90 {
		t.Errorf("重复率偏高: 仅 %d/100 不重复", len(seen))
	}
	_ = strings.Contains
}
