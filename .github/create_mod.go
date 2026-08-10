// Command create_mod builds a Go module zip in proxy.golang.org layout from a
// checkout directory, using the canonical golang.org/x/mod/zip packer.
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <src-dir> <out-zip>")
		os.Exit(2)
	}
	modPath, version, srcDir, outZip := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(outZip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %s: %v\n", outZip, err)
		os.Exit(1)
	}

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "CreateFromDir %s: %v\n", srcDir, err)
		f.Close()
		os.Exit(1)
	}
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "close %s: %v\n", outZip, err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s for %s@%s\n", outZip, modPath, version)
}
