// Do various units conversions.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jreisinger/lenconv"
	"gopl.io/ch2/tempconv"
)

var t = flag.Bool("t", false, "convert temperatures")
var l = flag.Bool("l", false, "convert lengths")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		u, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}

		// Temparatures conversions.
		if *t {
			f := tempconv.Fahrenheit(u)
			c := tempconv.Celsius(u)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}

		// Length conversions.
		if *l {
			cm := lenconv.Cm(u)
			in := lenconv.Inch(u)
			fmt.Printf("%s = %s, %s = %s\n",
				cm, lenconv.CToI(cm), in, lenconv.IToC(in))
		}
	}
}
