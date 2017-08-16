package pkgb

import (
	"fmt"
	"github.com/tanyfx/pkgd"
)

func PrintInfo() {
	fmt.Printf("pkgb, branch master:\t%s\n", pkgd.GetInfo())
}
