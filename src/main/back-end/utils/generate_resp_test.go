package utils

import "testing"

/**
    utils
    @author: roccoshi
    @desc: //TODO
**/

func TestNewBasicResp(t *testing.T) {
	t.Logf("%#v", NewBasicResp())
}

func TestNewErrBasicResp(t *testing.T) {
	t.Logf("%#v", NewErrBasicResp())
}
