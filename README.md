# `git-get`

Like `go get`, but for any `git` source.

```
$ git-get -h
usage: git-get REPO

Clone a Git repository, preserving remote structure under GITPATH.

  -version
    	print version and exit

Arguments:
  REPO     repository to clone
Environment variables:
  GITPATH  base of local tree of Git clones; defaults to $HOME/src
```

The default value for `GITPATH` is intended to compatible with a `GOPATH` of `$HOME`.

## Installation

- Download the [latest release](https://github.com/whilp/git-get/releases/latest).
- Copy `git-get` to a directory on your `$PATH`.

For example:

```console
curl -LO https://github.com/whilp/git-get/releases/download/v0.7/git-get-darwin-amd64
curl -LO https://github.com/whilp/git-get/releases/download/v0.7/git-get-darwin-amd64.sha256
shasum -a 256 -p -c git-get-darwin-amd64.sha256
export PATH=~/bin:$PATH
chmod a+x git-get-darwin-amd64
mkdir -p ~/bin
mv git-get-darwin-amd64 ~/bin/git-get
```

`git` will delegate the `get` subcommand to `git-get`, such that the following works:

```
git get github.com:whilp/git-get
```

## License

MIT; Copyright (c) 2014 Will Maier <wcmaier@m.aier.us>
