# fmlint
Front Matter Lint:  Lint your Markdown Front Matter


# Installation
This tool can be installed either with `go install`, or through a binary release download.  

### `go install`
```bash
# ensure that $GOPATH is set!
go install github.com/karldreher/fmlint@latest
# If desired, replace @latest with any specific tag.
```

### Releases

See [Releases](https://github.com/karldreher/fmlint/releases) and download the appropriate one for your platform.  

## Usage
Once installed, the program can be run from a command line.

You can learn more about the program with `fmlint --help.`

### Usage - Linting Front Matter

This expects a directory structure which has Markdown with [Front Matter](https://gohugo.io/content-management/front-matter/), as is typically part of a Hugo content directory.  This tool *should* be compatible with Jekyll content directories as well.  

The way to lint front matter with this tool is using `fmlint lint <subcommand>`

Currently, the only `lint` sub-command available is `tags`.  You can read more about what it does, with `fmlint lint tags --help`.

```bash
# Assumes a content/ directory below current working directory.  Otherwise, supply it with --folder.
fmlint tags
```


### Rule IDs
Each lint rule is identified with a `rule-id`.  You can find a list of all `rule-id`s, with the `list` command:

```
fmlint list
```

## Usage - Github Actions
This tool is **built** with CI environments in mind.  
Below is an example of how to implement a concise, helpful workflow.

```yaml
name: Lint
on:
  pull_request:
    types:
      - opened
      - edited
      - synchronize

jobs:  
  lint-front-matter:
    runs-on: ubuntu-22.04
    steps: 
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v4
        with:
          go-version: 1.21

      - name: Install fmlint
        run: go install github.com/karldreher/fmlint@latest
        
      - name: Run fmlint
        run: fmlint lint tags
```
