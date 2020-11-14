# Templater

Simple TOML based go template renderer.

## Build

Run `make` to build bin/templater

To run tests, you can do `make test`

## Usage

```sh
usage: templater -f deployment.yml.tmpl -i data.yml -c myconf.toml

-c string
Configuration file (TOML document)
-f string
Input template (go template file)
-i string
Data file (TOMLdocument)
```

This program assumes the following:

- Your input template should use the [standard go template format](https://golang.org/pkg/text/template/).

- The data file should be a [TOML document](https://toml.io/en/), and the values will be available as variables in the template

- The config file TOML currently only supports one key, `output_file`, to
  specify filename we write the template to.

## Improvements

- By default, it would be nicer if Templater defaulted to stdin and stdout for input/output
- Unless we come up with some more configuration, it would probably be nicer to specify the output directly as a parameter.
- Check for existing file and only overwrite with -force flag or similar.
- Support long versions of the flags?
- Improve error message for corner cases when input files are not correctly formatted.
- Increase test coverage
- Use [gtf](https://github.com/leekchan/gtf) or something similar to provide more template functions.

## License

This project is licensed under the terms of the MIT license.
