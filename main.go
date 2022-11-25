package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	jww "github.com/spf13/jwalterweatherman"
	"golang.org/x/net/nettest"
	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	session.Must(session.NewSession())
	jww.ERROR.Println()

	nettest.SupportsIPv4()
	&git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
	}
}
