# go-randpass

生成随机密码，用 `crypto/rand`，零依赖。

## 用法

```bash
go run .                  # 默认 16 位，含大小写字母和数字
go run . -l 24 -n 5       # 5 个 24 位密码
go run . -s               # 加上符号
go run . --no-upper -l 10 # 只要小写和数字，10 位
```

默认不含符号，要符号加 `-s`。字符池全关掉会直接报错，不会产出空密码。
