# fmlint
Front Matter Lint:  Lint your Markdown Front Matter


## Rule IDs
Each lint rule is identified with a `rule-id`.  You can find a list of all `rule-id`s, with the `list` command:

```
fmlint list
```
### Developer Info
Rules, within this program, are identified by a `cobraCmd` Annotation:
```
	Annotations: map[string]string{"rule-id": "none"},
```
`rule-id`s should be a dash-cased lint rule which uniquely identifies the behavior, or `"none"` if it is for internal use and not used for linting.
