# pwgen

安装

```
go get github.com/jetiny/pwgen
```

简单的密码生成, 生成后自动进入剪切板

```
pwgen
```

```
/;vRXH>&
password has been copied to the clipboard.
```

```
pwgen --help
Usage of pwgen:
  -b	copy to clipboard, 复制到剪切板 (default true)
  -c int
    	count of output, 密码个数 (default 1)
  -n int
    	number of characters, 密码位数 (default 8)
  -nc int
    	minimum count of numbers (default: any), 数字个数 (default -1)
  -sc int
    	minimum count of symbols (default: any), 字符个数 (default -1)
  -v	show version
```
