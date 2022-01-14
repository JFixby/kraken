package input

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"testing"
)

func TestInput(t *testing.T) {
	home := fileops.Abs("")
	list := fileops.ListFiles(home, func(file string) bool { return true }, fileops.ALL_CHILDREN)
	pin.D("list", list)
}
