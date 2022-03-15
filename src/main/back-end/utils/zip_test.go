package utils

import (
	"path/filepath"
	"testing"
)

/**
    utils
    @author: roccoshi
    @desc: test for zip
**/

func TestCreate(t *testing.T) {
	dir := "./testdata"
	fileName := filepath.Join(dir, "test.zip")
	zipFile, err := Zip.Create(fileName)
	if nil != err {
		t.Error(err)

		return
	}

	err = zipFile.AddDirectoryN(".", "../keystore_storage")
	if err != nil {
		return
	}

	err = zipFile.Close()
	if nil != err {
		t.Error(err)

		return
	}
}

func TestUnzip(t *testing.T) {
	dir := "./testdata"
	err := Zip.Unzip(dir+".zip", dir)
	if nil != err {
		t.Error(err)

		return
	}
}
