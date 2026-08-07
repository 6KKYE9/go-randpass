package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const (
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	symbols = "!@#$%^&*()-_=+[]{}<>?"
)

// 按开关拼出可选字符池，避免某些场景不要符号。
func buildCharset(useLower, useUpper, useDigit, useSymbol bool) string {
	var pool string
	if useLower {
		pool += lower
	}
	if useUpper {
		pool += upper
	}
	if useDigit {
		pool += digits
	}
	if useSymbol {
		pool += symbols
	}
	return pool
}

// 生成 n 个长度为 length 的密码。
// 每个字符都走 crypto/rand，比 math/rand 更适合干这个。
func genPasswords(n, length int, pool string) ([]string, error) {
	if len(pool) == 0 {
		return nil, fmt.Errorf("字符池为空，至少打开一类字符")
	}
	if length <= 0 {
		return nil, fmt.Errorf("长度得是正整数")
	}
	out := make([]string, 0, n)
	max := big.NewInt(int64(len(pool)))
	for i := 0; i < n; i++ {
		b := make([]byte, length)
		for j := range b {
			idx, err := rand.Int(rand.Reader, max)
			if err != nil {
				return nil, err
			}
			b[j] = pool[idx.Int64()]
		}
		out = append(out, string(b))
	}
	return out, nil
}

func main() {
	length := 16
	count := 1
	useLower, useUpper, useDigit, useSymbol := true, true, true, false

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-l", "--length":
			if i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &length)
				i++
			}
		case "-n", "--count":
			if i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &count)
				i++
			}
		case "-s", "--symbols":
			useSymbol = true
		case "--no-lower":
			useLower = false
		case "--no-upper":
			useUpper = false
		case "--no-digit":
			useDigit = false
		case "-h", "--help":
			fmt.Println("go-randpass 生成随机密码")
			fmt.Println("用法: go-randpass [-l 长度] [-n 个数] [-s]")
			fmt.Println("  -l  密码长度，默认 16")
			fmt.Println("  -n  生成几个，默认 1")
			fmt.Println("  -s  包含符号（默认不含）")
			return
		}
	}

	pool := buildCharset(useLower, useUpper, useDigit, useSymbol)
	pwds, err := genPasswords(count, length, pool)
	if err != nil {
		fmt.Fprintln(os.Stderr, "出错:", err)
		os.Exit(1)
	}
	for _, p := range pwds {
		fmt.Println(p)
	}
}
