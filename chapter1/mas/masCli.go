package mas

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/gplCLI"
)

// Command line help func

// CLI wrapper for *flag.FlagSet
type CLI struct {
	set *flag.FlagSet
}

// cmd allows to refer call send this module the CLI argument
var cmd CLI
var callMethod *string // Flag that stores value for -name="callMethod"
var n1 *int // Flag that stores value for: the-gpl mas -comp  -n1=123 -n2=345
var n2 *int // Flag that stores value for -n2=234

// InitCli for command: the-gpl mas -fn=array
func InitCli() {
	cmd.set = flag.NewFlagSet("mas", flag.ContinueOnError)
	callMethod = cmd.set.String("fn", "array", "[array comp slice]")
	n1 = cmd.set.Int("n1", 123, "first number")
	n2 = cmd.set.Int("n2", -46, "second number")
	gplCLI.Add("mas", cmd)
}

// ExecCmd run bit count from CLI
func (m CLI) ExecCmd(args []string) {
	err := m.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: MAS (Maps Arrays Slices) Parse Error %s\n", err.Error())
		return
	}

	switch *callMethod {
	case "array":
		IterateOverArray()
	case "comp":
		compResult, diff := CompareNumbers(*n1, *n2)
		fmt.Printf("mas:CompareNumbers: ints: %d == %d, => %t, differance: %d\n", *n1, *n2, compResult, diff)
	case "slice":
		AddToSlices()
	}
}

// DisplayHelp prints help on command line for bits module
func (m CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl mas. (array comp slice)")
	m.set.PrintDefaults()
}
