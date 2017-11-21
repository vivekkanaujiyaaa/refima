package kernel

import (
	"fmt"
	"testing"

	"github.com/xalabs/refima/config"
)

var kernel *Kernel

func init() {
	c := config.Config{}
	kernel = NewKernel(c)
}

func TestGetDirContent(t *testing.T) {
	contents, _ := kernel.GetDirContent("./")
	fmt.Println(string(contents))

}
