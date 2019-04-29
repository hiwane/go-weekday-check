# go-weekday-check

[![Build Status](https://travis-ci.org/hiwane/go-weekday-check.svg?branch=master)](https://travis-ci.org/hiwane/go-weekday-check)
![GitHub](https://img.shields.io/github/license/hiwane/go-weekday-check.svg)

## install

```sh
> git clone https://github.com/hiwane/go-weekday-check.git
> cd go-weekday-check
> go install
```

## Usage

```sh
> LANG=c date '+%F %A'
2019-04-29 Monday
> cat /tmp/boo
2019-04-29 (Sun)
> go-weekday-check /tmp/boo
/tmp/boo:1: invalid weekday 2019-04-29 (Sun)
> cat /tmp/boo
2019-04-29 (Sun)
> go-weekday-check --fix /tmp/boo
/tmp/boo:1: invalid weekday 2019-04-29 (Sun)
> cat /tmp/boo
2019-04-29 (Mon)
```

```
YYYY-MM-DD (WWW)
YYYY/MM/DD (WWW)
YYYY年MM月DD日 (WWW)
```


