package pkgb

import (
	"fmt"
	"github.com/tanyfx/pkgd"
)

func PrintInfo() {
	fmt.Printf("pkgb, branch forkb:\t%s\n", pkgd.GetInfo())
}
