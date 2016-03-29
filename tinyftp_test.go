package tinyftp

import (
	"net"
	"os"
	"testing"
	"time"
)

const (
	ftpHost1   = "prep.ai.mit.edu:21"
	ftpDir1    = "/gnu/chess"
	ftpHost2   = "ftpprd.ncep.noaa.gov:21"
	ftpDir2    = "/pub/data/nccf/com/rtofs/prod/rtofs."
	uploadHost = "localhost:21"
	uploadDir  = "/test_dir"
	uploadFile = "./tinyftp_test.go"
)

func testNameList(t *testing.T, ftpHost, ftpDir string) {
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

func testUpload(t *testing.T, ftpHost, ftpDir string) {
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
	code, msg, err = c.Type("I")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	dconn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer dconn.Close()
	content := []byte("Test!!! Test!!! Test!!!")
	t.Log(content)
	n, code, msg, err := c.Upload("tinyftp_test_file.txt", dconn, content)
	t.Logf("code %d, msg %v", code, msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("written %d, code %d, msg %v", n, code, msg)
	code, msg, err = c.Quit()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
}

func testUploadFrom(t *testing.T, ftpHost, ftpDir, file string) {
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
	code, msg, err = c.Type("I")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
	dconn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer dconn.Close()
	fconn, err := os.OpenFile(file, os.O_RDONLY, 0444)
	if err != nil {
		t.Fatal(err)
	}
	defer fconn.Close()
	n, code, msg, err := c.UploadFrom(fconn.Name(), dconn, fconn)
	t.Logf("code %d, msg %v", code, msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("written %d, code %d, msg %v", n, code, msg)
	code, msg, err = c.Quit()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code %d, msg %v", code, msg)
}

func TestTinyFTP1(t *testing.T) {
	testNameList(t, ftpHost1, ftpDir1)
}

func TestTinyFTP2(t *testing.T) {
	testNameList(t, ftpHost2, ftpDir2+time.Now().Format("20060102"))
}

func TestTinyFTP3(t *testing.T) {
	_, _, _, err := DialTimeout("tcp", "localhost:21", time.Nanosecond)
	if neterr, ok := err.(net.Error); ok && !neterr.Timeout() {
		t.Fatal("Expected a timeout to occur")
	}
}

func TestTinyFTP4(t *testing.T) {
	testUpload(t, uploadHost, uploadDir)
}

func TestTinyFTP5(t *testing.T) {
	testUploadFrom(t, uploadHost, uploadDir, uploadFile)
}
