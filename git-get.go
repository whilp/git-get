// git-get is like `go get` but for any Git source.

/*
 * Copyright (c) 2014 Will Maier <wcmaier@m.aier.us>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
	"syscall"
)

// usage prints a helpful usage message.
func usage() {
	self := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "usage: %s REPO\n\n", self)
	fmt.Fprint(os.Stderr, "Clone a Git repository, preserving remote structure under GITPATH.\n\n")
	fmt.Fprintln(os.Stderr, "Arguments:")
	fmt.Fprintln(os.Stderr, "  REPO     repository to clone")
	fmt.Fprintln(os.Stderr, "Environment variables:")
	fmt.Fprintln(os.Stderr, "  GITPATH  base of local tree of Git clones; defaults to $HOME/src")
	os.Exit(2)
}

// lsRemote calls `git ls-remote --get-url`, resolving a remote to a local path.
func lsRemote(remote string) (string, error) { 	
	cmd := exec.Command("git", "ls-remote", "--get-url", remote)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out[:])), nil
}

// stripUser removes an optional 'user@' portion from a Git remote.
func stripUser(remote string) (string) {
	fields := strings.SplitN(remote, "@", 2)
	switch len(fields) {
	case 1: remote = fields[0]
	case 2: remote = fields[1]
	}

	return remote
}

// importPath converts a Git remote path to a local path.
func importPath(remote string) (string) {
	return strings.Replace(stripUser(remote), ":", "/", 1)
}

// getGitpath finds a suitable value for GITPATH.
// If the GITPATH environment variable is not set, it defaults to `$HOME/src`.
func getGitpath() (string) {
	p := os.Getenv("GITPATH")
	if p == "" {
		var home string
		u, err := user.Current()
		if err != nil {
			home = os.Getenv("HOME")
		} else {
			home = u.HomeDir
		}
		p = path.Join(home, "src")
	}
	return p
}

// clone calls `git clone remote local`.
func clone(remote, local string) (error) {
	cmd := exec.Command("git", "clone", remote, local)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	remote := args[0]
	resolved, err := lsRemote(remote)
	if err != nil {
		log.Panic(err)
	}

	gitroot := getGitpath()
	local := path.Join(gitroot, importPath(resolved))

	exit := 0
	err = clone(remote, local)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus := exitError.Sys().(syscall.WaitStatus)
			exit = waitStatus.ExitStatus()
		} else {
			log.Panic(err)
		}
	}
	os.Exit(exit)
}
