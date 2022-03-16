package main

import (
	"github.com/lekan-pvp/mycheck/test/addlint"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
)

func main() {
	multichecker.Main(
		printf.Analyzer,
		shadow.Analyzer,
		addlint.CheckUseOsExit,
	)
}
