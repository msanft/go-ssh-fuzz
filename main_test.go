package main

import (
	"testing"

	"golang.org/x/crypto/ssh"
)

func FuzzSshParseAuthorizedKey(f *testing.F) {
	f.Add("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDHG4Gb4iiSrqLqrARn/2gmJvyBQ9I+qUN+TwdJDIJUZMrBA7cLNbX/qsyEgG9hMd7GqXVQW7USMu4Dl7S1h+3RUWVOM0QwAA35qjTUpulf4nSJvgIdCjp3DZ8yntHpfL7CeOpcQCi/d4TVHwXSvoPU1rMs566PLXMkJIlBc/Q86v6b+mBlsLRcecBoBWBXeIcABN3oUvBQ+O6A5HEBd+Ec2A2BLykRkBAeBpUSFEBF/5fDgnUsNgpa6lMMtSOCBjjko/PmGT7KuhMIdpdMtTSc3wc4QFKAFecddO5MgkPY0gAUVzVtEe1xlYen37+2sKzbfsQ2ZvW/tKxCVGc71SPTdtz+bhC6hIBLrT63RPUyXUmt1QiwveJHv8/tHYZmHDabAuVPqPnRw2mN37fx+oeWc4wE+B5GBQxdVDPM8M5IBbI6FB4IHA1SzGsBTOIaO1B+XSblxvk/8z1/BGM7ya0E2ilkGpD9C+4AH0yejT9ztFWqNf/JEvIedIXPqkptvvYKJTuhw+hsbsN2hAkF4yiEXko2cIV97btJJl3E9/M0dFSuh3Q3igejSIoDBZTl5Z73zJ4znKx/6FGUHk+Cq8JMU6X82Y3ULQaVPpaiCtXx+OJLlUeQ9K1WIn5wI3cdXj3A3gPXjks+lVxh4gUnzzv4jVDxP736TYJAXliR/d6Aaw== msanft@github/86771043 # ssh-import-id gh:msanft")

	f.Fuzz(func(t *testing.T, key string) {
		ssh.ParseAuthorizedKey([]byte(key))
	})
}

func FuzzSshParsePublicKey(f *testing.F) {
	f.Add("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIE+qJO8tDOXE2zmtG3/SsryMJDBxxY05mfiMVCAOgwEG")

	f.Fuzz(func(t *testing.T, key string) {
		ssh.ParsePublicKey([]byte(key))
	})
}

func FuzzSshParseKnownHosts(f *testing.F) {
	f.Add("52.156.216.95 ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDw1REpiwOue3JRGO16K/hKWo1Ni1cO7DW9PUbkbAVBnZ5lS9smrm7UBcRy08CtN3X/z8Slf9yyqmhellbagVYGQPWQ2OVHCqLJq/L6T5gEw7lIcZR/8Z+VjEFOIlkwGTrmBx6zznsRgi33muGFVK4J0fabjh7H4hCsbo6G2rS8c5g4IBtwRMfIXz/SfR+2BYosB9zbq0od+LNOncuhALMa7JeE5QMNzTzXD6qnLNSEIWm95OONRlboso5DyGQ1FSKeZz0ZzKF1Nyl/pVlIcXokQ4j7w/51wdkNZiD/CC+nj5hrKxXiTlSlI/ptguoJIz749/oLMnPMX0+ZxAkMVfjMZDBG3Pi5chJRAXkW9Jyn3z1aM6aAQhMqHIS+WJxeEZhWmQyTckyymoIRfgu//256VYsWZ819o4uV0GAmr2WdvCYuhuvxATWfv3wY8ocIODYCx3oEamcQtlpf7tLiy+n1QvEeWORcfAoc8xtpLzMPmQiTmPOJxSX7geyKIjA66TTxC/z4hN6qe4fgB77xseeCVJqGZAsb5totKHi9TqA/+3xsOoQafLkeryjcCdfAAQJNDXQcLDCwdU7W4cViCJQD7SkGKJejWFxvdbrGacg2zF7+7Zjyy3LUSdCWsCNE/NeMBX22afV0W91JlnnFaDa+HKCwgmyLYxZwCPa/prJROw==")

	f.Fuzz(func(t *testing.T, key string) {
		ssh.ParseKnownHosts([]byte(key))
	})
}
