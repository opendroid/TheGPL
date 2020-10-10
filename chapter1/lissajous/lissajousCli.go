package lissajous

import (
	"flag"
	"fmt"
	"github.com/opendroid/the-gpl/gplCLI"
	"io"
	"os"
)

type CLI struct {
	set *flag.FlagSet
}

// bitCountCmd allows to refer call send this module the CLI argument
var cmd CLI
var cycles, size, frames *int
var outFileName *string
// InitCli for the "bits" command
func InitCli() {
	cmd.set = flag.NewFlagSet("lissajous", flag.ExitOnError)
	cycles = cmd.set.Int("cycles", 2, "number of cycles")
	size = cmd.set.Int("size", 512, "size of square image")
	frames = cmd.set.Int("frames", 10, "number of frames")
	outFileName = cmd.set.String("file", "lis.gif", "name of output file")
	gplCLI.Add("lissajous", cmd)
}

// ExecCmd run lissajous count from CLI
func (l CLI) ExecCmd(args []string) {
	err := l.set.Parse(args)
	if err != nil {
		fmt.Printf("ExecCmd: Lissajous %s\n", err.Error())
		return
	}
	f, err := os.Create(*outFileName)
	if err != nil {
		fmt.Printf("ExecCmd: Write to file %s: %s\n", *outFileName, err.Error())
		return
	}
	saveSquareImage(*cycles, *size, *frames, f)
}

// DisplayHelp prints help on command line for lissajous module
func (l CLI) DisplayHelp() {
	fmt.Println("\nUsage: the-gpl lissajous. Saves a lissajous figure of n cycles, n frames and square size.")
	l.set.PrintDefaults()
}

// saveImageSquare tests lissajous package
func saveSquareImage(cycles, size, frames int, w io.Writer) {
	config := Config{
		Cycles:     cycles,
		Resolution: 0.000001,
		Size:       size,
		NFrames:    frames,
		DelayMS:    10,
	}
	Lissajous(config, w)
}