package main

import (
	"testing"

	"golang.org/x/crypto/ssh"
)

func FuzzSshParseAuthorizedKey(f *testing.F) {
	f.Add("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIE+qJO8tDOXE2zmtG3/SsryMJDBxxY05mfiMVCAOgwEG")

	f.Fuzz(func(t *testing.T, key string) {
		ssh.ParseAuthorizedKey([]byte(key))
	})
}
