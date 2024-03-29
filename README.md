# fmlint
## Front Matter Lint:  Lint your Markdown Front Matter

### Why is this helpful? 
Tools like Hugo can make static websites very easy to implement, especially with healthy theme support and extensibility. 

A part of what makes static site generators (SSGs) work well is the metadata, known as *front matter*, of the content.  

While front matter is mostly open-ended, there do exist subtle failure modes of this metadata.  The purpose of this tool is to provide some *opinionated* guidance to avoid encountering those failures.


# Installation
## Releases

See [Releases](https://github.com/karldreher/fmlint/releases) and download the appropriate one for your platform.  Place this in a directory resolvable in `PATH`, or run it relative to current working dir (`./fmlint`).

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

### Usage - Disabling Lint Rules
To disable a lint rule, configure a `yaml` file with a list of [`rule-id`s](#rule-ids) which you want to disable from checking.  
#### Example
config.yaml
```yaml
disabled_rules:
  - "tags-sorted"
  - "draft-enabled"
```

Then, pass this in execution, e.g. `fmlint lint all --config config.yaml`.  

## Usage - Github Actions
This tool is **built** with CI environments in mind.  
Below is an example of how to implement a concise, helpful workflow.

```yaml
name: Lint
on:
  pull_request:
    types:
      - synchronize

jobs:  
  lint-front-matter:
    runs-on: ubuntu-22.04
    steps:
      - uses: actions/checkout@v4

      - name: Download fmlint
        # It is reccomended to pin the version and download the binary.
        # however, with the addition of another preceding step, you could get the latest version instead.
        run: curl -o fmlint.tar.gz -LO https://github.com/karldreher/fmlint/releases/download/v2.2.1/fmlint_2.2.1_linux_amd64.tar.gz

      - name: Install fmlint
        run: tar -xvf fmlint.tar.gz fmlint

      - name: Run fmlint
        run: ./fmlint lint all
```
