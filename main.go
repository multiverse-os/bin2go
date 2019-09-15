package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("[singularity] bin2go: Ultra simple binary conversion")
	fmt.Println("==========================================================================================")
	fmt.Println("`bin2go` is an ultra simple binary-to-GoSource conversion tool that provides no compression")
	fmt.Println("or file system, like the popular `go-bindata`.\n")

	args := os.Args
	if len(args) < 2 {
		fmt.Println("[error] failed to provide arguments, try again using the following:\n")
		fmt.Println("    bin2go [binary path (/usr/bin/ruby)] [filename.go]\n")
	} else {
		var binaryFilename string
		if len(args[1]) == 0 {
			fmt.Println("[error] failed to provide path to binary:\n")
			fmt.Println("    bin2go [binary path (/usr/bin/ruby)] [filename.go]\n")
		} else {
			binaryFilename = args[1]
		}

		var outputFilename string
		if len(args) == 2 || len(args[2]) == 0 {
			outputFilename = "./bin2go.go"
		} else {
			outputFilename = args[2]
		}

		outputFile, err := os.Create(outputFilename)
		if err != nil {
			fmt.Println("[bin2go] failed to open output file:", err)
			os.Exit(1)
		}
		defer outputFile.Close()

		basename := strings.Split(filepath.Base(binaryFilename), ".")[0]
		if len(basename) == 0 {
			fmt.Println("[bin2go] binary name must not be empty `"+basename+"`:", err)
			os.Exit(1)
		}

		outputFile.Write([]byte("package " + strings.ToLower(basename) + "\n\nvar (\n\tBinary = []byte{"))
		binaryFile, err := ioutil.ReadFile(binaryFilename)
		if err != nil {
			fmt.Println("[bin2go] failed to open binary file:", err)
			os.Exit(1)
		}

		var data []string
		for _, binaryData := range binaryFile {
			byteString := fmt.Sprintf("%v", binaryData)
			data = append(data, byteString)
		}

		outputFile.WriteString(strings.Join(data, ", "))
		outputFile.Write([]byte("}\n)"))
		fmt.Println("[bin2go] Binary to byte slice conversion completed!")
	}
}
