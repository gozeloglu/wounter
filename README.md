# wounter

Wounter is a CLI tool for counting words in a given string. You need to pass the string in quotes or double quotes.

## Install

```shell
go install github.com/gozeloglu/wounter@latest
```

## Run

```shell
wounter "Example string for counting." 
```

As another option, you can pass file path to count words.

```shell
wounter --path example.txt
```

Output will be as follows:

```shell
============================
Word count: 4
============================
```

## LICENSE

[MIT](LICENSE)
