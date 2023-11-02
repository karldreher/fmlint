# fmlint
## Front Matter Lint:  Lint your Markdown Front Matter

### Why is this helpful? 
Tools like Hugo can make static sites very easy to implement, especially with healthy theme support and extensibility. 

A part of what makes Hugo work well is the metadata, known as *front matter*, of the content.  

While front matter is mostly open-ended, there do exist subtle failure modes of this metadata.  The purpose of this tool is to provide some *opinionated* guidance to avoid encountering those failures.


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

You can learn more about the program with `fmlint --help.`  `--help` is available from many contexts and will give you the most information when used interactively. 

### Rule IDs
Each lint rule is identified with a `rule-id`.  You can find a list of all `rule-id`s, with the `list` command:

```
fmlint list
```

Each entry will have a brief description of the lint command.  For a long one, try `fmlint lint <command> --help`.


### Usage - Linting Front Matter

This expects a directory structure which has Markdown with [Front Matter](https://gohugo.io/content-management/front-matter/), as is typically part of a Hugo content directory.  This tool *should* be compatible with Jekyll content directories as well.  


Using a command identified by [`fmlint list`](#rule-ids), you can then use than lint rule using `fmlint lint <command>`


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
