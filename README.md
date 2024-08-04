# go-template-runner

Simple tool to run go templates against JSON files

## Installation

```bash
go install github.com/alejo47/go-template-runner/cmd/tmpl@latest
```

## Usage


### Reading template from stdin

This will prompt the user to write the template, use `Ctrl+D` to send an EOF prompting the program to continue.
```sh
templ -json data.json
```

### Reading template from pipe

```sh
cat template.file | templ -json data.json
```
