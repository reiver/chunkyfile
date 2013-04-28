package main



import (
	"testing"
	"io/ioutil"
	"os"
)









type TestChunkyFileDir struct {
	dirPath string
}

func NewTestChunkyFileDir(t *testing.T) TestChunkyFileDir {
	me := TestChunkyFileDir{}

	var err error

	me.dirPath, err = ioutil.TempDir("", "chunkyfile_test_")
	if nil != err {
		t.Errorf("Received error when trying to create a temp dir: %v", err)
	}

	return me
}

func (me *TestChunkyFileDir) CleanUp(t *testing.T) {

	err := os.RemoveAll(me.dirPath)
	if nil != err {
		t.Errorf("Received error when trying to clean up temp dir: %v", err)
	}
}

func (me *TestChunkyFileDir) DirPath() string {
	return me.dirPath
}









func TestDirDoesNotExist(t *testing.T) {
	dirPath := "doesNotExist"

	cf, err := New(dirPath)
	if nil == err {
		t.Errorf("Should have received an error for a non-existant directory, but did not.")

		cf.Close()
	}
}

func TestDirDoesExist(t *testing.T) {
	tchd := NewTestChunkyFileDir(t)
	defer tchd.CleanUp(t)


	cf, err := New(tchd.DirPath())
	if nil != err {
		t.Errorf("Should not have received an error for an existant directory, but did.")
	}
	defer cf.Close()
}

func TestShortestWrite(t *testing.T) {
	tchd := NewTestChunkyFileDir(t)
	defer tchd.CleanUp(t)


	cf, err := New(tchd.DirPath())
	if nil != err {
		t.Errorf("Should not have received an error for an existant directory, but did.")
	}
	// Do not defer close, since want to force a close early.


	needle := "a"

	cf.Write( []byte(needle) )
	cf.Close()


	filePath := tchd.DirPath() + "/cf_0.chunk"
	_, err = os.Stat(filePath)
	if nil != err {
		t.Errorf("Expected chunk file not created.")
	}


	buff, err := ioutil.ReadFile(filePath)
	if nil != err {
		t.Errorf("Could not read in content of created chunk file.")
	}

	buffStr := string(buff)

	if needle != buffStr {
		t.Errorf("Content of created chunk file not what expected.")
	}
}

