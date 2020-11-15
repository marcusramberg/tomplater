package pkg

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/pelletier/go-toml"
)

// TemplateRenderer is glue between go-toml and text/template
type TemplateRenderer struct {
	TemplateFile string
	DataFile     string
	ConfigFile   string
}

// NewTemplateRenderer is TemplateRenderer constructor
func NewTemplateRenderer(templateFile string, dataFile string, configFile string) *TemplateRenderer {
	return &TemplateRenderer{templateFile, dataFile, configFile}
}

func getTemplate(templateFile string) (*template.Template, error) {

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		return nil, fmt.Errorf("Invalid template: %s (Not found)", templateFile)
	} else if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func getData(dataFile string) (map[string]interface{}, error) {
	data, err := toml.LoadFile(dataFile)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		return nil, fmt.Errorf("Invalid data: %s (Not found)", dataFile)
	} else if err != nil {
		return nil, fmt.Errorf("Failed to parse datafile %s: %s", dataFile, err.Error())
	}
	return data.ToMap(), nil
}

func getOutput(configFile string) (string, error) {
	config, err := toml.LoadFile(configFile)
	if err != nil {
		return "", err
	}
	return config.Get("output_file").(string), nil

}

// Render is responsible for Rendering file to specified output.
func (t *TemplateRenderer) Render() error {
	template, err := getTemplate(t.TemplateFile)
	if err != nil {
		return err
	}
	data, err := getData(t.DataFile)
	if err != nil {
		return err
	}
	output, err := getOutput(t.ConfigFile)
	if err != nil {
		return err
	}
	return renderTemplate(template, data, output)
}

func renderTemplate(tmpl *template.Template, data map[string]interface{}, outputFile string) error {
	fmt.Printf("Rendering %s with %v to %s\n", tmpl.ParseName, data, outputFile)
	destination, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("Could not create output file: %v", err)
	}
	defer destination.Close()
	return tmpl.Execute(destination, data)
}
