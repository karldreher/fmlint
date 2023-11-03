.PHONY: new-lint-rule default

default: 
# Do nothing, this makefile is only for convenience in creating new rules
	@ exit

new-lint-rule:
	@echo "Enter a new lint command name: "
	@read lintcmd; \
	cobra-cli add $$lintcmd --parent lintCmd; \
	mv cmd/$$lintcmd.go cmd/lint_$$lintcmd.go; \
	echo "Moved to cmd/lint_$$lintcmd.go"
