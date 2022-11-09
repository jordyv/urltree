# urltree

## Description
This tool will output a tree of URLs from a list of URLs. The tree's root level are all the unique root domains found in the input.
All unique paths of each domain will be printed as second level and the unique query string combinations as third level.

An example of how you might use this during bug bounty:
```
$ echo example.com | subfinder > subs
$ cat subs | httpx > httpx.out
$ cat httpx.out | cut -d ' ' -f 1 > hosts
$ cat hosts | gau > gau.out
$ cat gau.out | urltree > urltree.out

# The urltree.out contains a sorted tree of all the unique URLs found by gau
```

## Installation

```
git clone https://github.com/jordyv/urltree.git
cd urltree
go build -o urltree cmd/urltree/main.go
```

_Go 1.18 or newer needed_.

## Usage

```
cat input| urltree
- https://a.foo.bar
  -- /a/b/coo
  -- /a/b/joo
  -- /a/c
  -- /b/a
    --- u=aa
    --- u=bb
- https://b.foo.bar
  -- /aa/bb
  -- /
```
