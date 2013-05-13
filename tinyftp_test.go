package tinyftp

import (
	"net"
	"testing"
)

const (
	ftpHost = "prep.ai.mit.edu:21"
	ftpDir  = "/gnu/chess"
)

func TestTinyFTP(t *testing.T) {
	c, code, msg, err := Dial("tcp", ftpHost)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	code, msg, err = c.Login("", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	code, msg, err = c.Type("A")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	code, msg, err = c.Cwd(ftpDir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	addr, code, msg, err := c.Passive()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("addr %s, code %d, msg %v", addr, code, msg)
	dconn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer dconn.Close()
	dir, code, msg, err := c.NameList("", dconn)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	t.Logf("dir %#v", dir)
	code, msg, err = c.Type("I")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	for _, name := range dir {
		t.Logf("doing %v", name)
		size, code, msg, err := c.Size(name)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("code %d, msg %v", code, msg)
		t.Logf("file %v, size %v", name, size)
	}
	code, msg, err = c.Quit()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
}
