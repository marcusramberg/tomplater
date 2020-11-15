package main

import (
	"flag"
	"fmt"
	"os"

	"tomplater/pkg"
)

func main() {

	templateFile := flag.String("f", "", "Input template (go template file)")
	dataFile := flag.String("i", "", "Data file (TOML document)")
	configFile := flag.String("c", "", "Configuration file (TOML document)")

	flag.Parse()
	if *templateFile == "" || *dataFile == "" || *configFile == "" {
		printUsage("All parameters are required")
	}

	renderer := pkg.NewTemplateRenderer(*templateFile, *dataFile, *configFile)
	err := renderer.Render()
	if err != nil {
		printUsage(err.Error())
	}
	fmt.Println("Operation completed successfully.")

}

func printUsage(reason string) {
	fmt.Println(reason)
	fmt.Printf("\nusage: tomplater -f deployment.yml.tmpl -i data.toml -c myconf.toml\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}
