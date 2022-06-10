package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"os"
	"time"
)

var myLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

const url = "https://github.com/freebsd/freebsd-src.git"

const directoryByBranch = "/tmp/repo-go-git-branch"
const directoryByTag = "/tmp/repo-go-git-tag"
const directoryFullClone = "/tmp/repo-go-git-full"

const branch = "releng/10.0"
const tag = "release/10.0.0"

func main() {
	goGitShallowCloneByBranch()
	goGitShallowCloneByTag()
	goGitFullClone()
}

func cloneMe(reference plumbing.ReferenceName, url, directory string) {
	os.RemoveAll(directory)

	myLogger.Printf("Cloning repo %v into local directory %v with reference %v ...", url, directory, reference)
	start := time.Now()
	var cloneErr error
	if "" == reference {
		_, cloneErr = git.PlainClone(directory, false, &git.CloneOptions{
			URL: url,
		})
	} else {
		_, cloneErr = git.PlainClone(directory, false, &git.CloneOptions{
			URL:           url,
			ReferenceName: reference,
			SingleBranch:  true,
			Depth:         1,
		})
	}

	elapseTime := time.Since(start)
	AssertTruef(nil == cloneErr, myLogger, "Expect err is nil. But got: %v", cloneErr)
	myLogger.Printf("Cloning by %v is done! It took %v", reference, elapseTime)
}

func AssertTruef(b bool, logger *log.Logger, format string, v ...any) {
	if !b {
		myLogger.Fatalf(format, v...)
	}
}

func goGitShallowCloneByBranch() {
	refByBranch := plumbing.NewBranchReferenceName(branch)
	myLogger.Printf("Go Git shallow clone by branch")
	cloneMe(refByBranch, url, directoryByBranch)
}

func goGitShallowCloneByTag() {
	refByTag := plumbing.NewTagReferenceName(tag)
	myLogger.Printf("Go Git shallow clone by tag")
	cloneMe(refByTag, url, directoryByTag)
}

func goGitFullClone() {
	myLogger.Printf("Go Git shallow clone by tag")
	cloneMe("", url, directoryFullClone)
}
