# `git-get`

Like `go get`, but for any `git` source.

```
 $ ./git-get -h
usage: git-get REPO

Clone a Git repository, preserving remote structure under GITPATH.

Arguments:
  REPO     repository to clone
Environment variables:
  GITPATH  base of local tree of Git clones; defaults to $HOME/src
```

The default value for `GITPATH` is intended to compatible with a `GOPATH` of `$HOME`.

## Installation

```
make test build
```

Copy `git-get` to a directory on your `$PATH`. `git` will delegate the `get` subcommand to `git-get`, such that the following works:

```
git get github.com:whilp/git-get
```
