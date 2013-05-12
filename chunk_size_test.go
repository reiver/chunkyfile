package main



import (
	"testing"
)



func TestChunkSize1(t *testing.T) {
	tchd := NewTestChunkyFileDir(t)
	defer tchd.CleanUp(t)


	var chunkSize uint64 = 1

	cf, err := NewWithSize(tchd.DirPath(), chunkSize)
	if nil != err {
		t.Errorf("Should not have received an error for an existant directory, but did.")
	}
	defer cf.Close()

	if chunkSize != cf.ChunkSize() {
		t.Errorf("The chunk size is not what is expected")
	}

}

func TestChunkSize2(t *testing.T) {
	tchd := NewTestChunkyFileDir(t)
	defer tchd.CleanUp(t)


	var chunkSize uint64 = 2

	cf, err := NewWithSize(tchd.DirPath(), chunkSize)
	if nil != err {
		t.Errorf("Should not have received an error for an existant directory, but did.")
	}
	defer cf.Close()

	if chunkSize != cf.ChunkSize() {
		t.Errorf("The chunk size is not what is expected")
	}

}

func TestChunkSize13(t *testing.T) {
	tchd := NewTestChunkyFileDir(t)
	defer tchd.CleanUp(t)


	var chunkSize uint64 = 13

	cf, err := NewWithSize(tchd.DirPath(), chunkSize)
	if nil != err {
		t.Errorf("Should not have received an error for an existant directory, but did.")
	}
	defer cf.Close()

	if chunkSize != cf.ChunkSize() {
		t.Errorf("The chunk size is not what is expected")
	}

}

func TestChunkSize1234567(t *testing.T) {
	tchd := NewTestChunkyFileDir(t)
	defer tchd.CleanUp(t)


	var chunkSize uint64 = 1234567

	cf, err := NewWithSize(tchd.DirPath(), chunkSize)
	if nil != err {
		t.Errorf("Should not have received an error for an existant directory, but did.")
	}
	defer cf.Close()

	if chunkSize != cf.ChunkSize() {
		t.Errorf("The chunk size is not what is expected")
	}

}



