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

> cat /tmp/foo
2019-04-29 (Tue)
2019/04/29 (Wednesday)
2019/04/29 (木)
2019/04/29 (金曜日)
2019年04月29日 (土)
2019年04月29日 (土)
19年04月29日 (土)
04月29日 (土)
> go-weekday-check --fix /tmp/foo
/tmp/foo:1: invalid weekday 2019-04-29 (Tue)
/tmp/foo:2: invalid weekday 2019/04/29 (Wednesday)
/tmp/foo:3: invalid weekday 2019/04/29 (木)
/tmp/foo:4: invalid weekday 2019/04/29 (金曜日)
/tmp/foo:5: invalid weekday 2019年04月29日 (土)
/tmp/foo:6: invalid weekday 19年04月29日 (土)
/tmp/foo:7: invalid weekday 04月29日 (土)
> cat /tmp/foo
2019-04-29 (Mon)
2019/04/29 (Monday)
2019/04/29 (月)
2019/04/29 (月曜日)
2019年04月29日 (月)
19年04月29日 (月)
04月29日 (月)
```


