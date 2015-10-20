# Save Urban Dictionary JSON API data to Couchbase 4

## Overview

Very simple command line program that you type in keyword(s) that are looked up on [urbandictionary.com](http://urbandictionary.com) and the JSON response saved to Couchbase 4.

## Usage

```
❯ go run main.go                                                                                                      ⏎
==> Enter keyword(s):
hipster
Inserted document CAS is `c8201d8e40`
==> Enter keyword(s):
hip hop
Inserted document CAS is `c8203d08d0`
==> Enter keyword(s):
```

## Install

In your GOPATH:

- `git clone https://github.com/couchbase/gocb.git`
- `git clone https://github.com/rudijs/urban.git`
