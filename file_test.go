package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var file = "./README.md"
var tmpDir = "./_THIS_IS_TEMP_DIR_"

func TestCheckFileExist(t *testing.T) {
	assert.Equal(t, true, CheckFileExist(file))
}

func TestCheckNotExist(t *testing.T) {
	assert.Equal(t, false, CheckFileNotExist(file))
}

func TestMakeDirIfNotExist(t *testing.T) {

	t.Log("-- make directory --")
	t.Logf("(before-c) dir %s Exist?:%v", tmpDir, CheckFileExist(tmpDir))
	t.Logf("(before-c) make directory if not exist")
	_ = MakeDirIfNotExist(tmpDir)
	t.Logf("(after -c) dir %s Exist?:%v", tmpDir, CheckFileExist(tmpDir))
	assert.Equal(t, true, CheckFileExist(tmpDir))

	t.Log("-- remove directory --")
	t.Logf("(before-r) remove directory if exist")
	_ = RemoveDirIfExist(tmpDir)
	t.Logf("(after -r) dir %s Exist?:%v", tmpDir, CheckFileExist(tmpDir))
	assert.Equal(t, true, CheckFileNotExist(tmpDir))

}
