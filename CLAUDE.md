# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Purpose

`package-sorter` is a Go CLI tool that sorts packages based on predefined rules. 


## Commands

Once `go.mod` is initialized:

```bash
# Build
go build ./...

# Run
go run main.go [args]

# Test all
go test ./...

# Test single package
go test ./internal/sorter/...

# Test single test
go test -run TestFunctionName ./...

# Lint (requires golangci-lint)
golangci-lint run
```

## Architecture

The project is a Go CLI application. Expected structure once implemented:

- **`main.go`** — entry point, parses CLI flags, delegates to core logic
- **`internal/sorter/`** — core sorting logic and predefined rules
- **`internal/rules/`** — rule definitions and configuration loading

The sorting rules are predefined but can be extended or overridden through the use of a file-based configuration.


## Rules
### Objective

Imagine you work in Smarter Technology’s robotic automation factory, and your objective is to write a function for one of its robotic arms that will dispatch the packages to the correct stack according to their volume and mass.

### Rules

Sort the packages using the following criteria:

- A package is **bulky** if its volume (Width x Height x Length) is greater than or equal to 1,000,000 cm³ or when one of its dimensions is greater or equal to 150 cm.
- A package is **heavy** when its mass is greater or equal to 20 kg.

You must dispatch the packages in the following stacks:

- **STANDARD**: standard packages (those that are not bulky or heavy) can be handled normally.
- **SPECIAL**: packages that are either heavy or bulky can't be handled automatically.
- **REJECTED**: packages that are **both** heavy and bulky are rejected.

### Implementation

Implement the function **`sort(width, height, length, mass)`** (units are centimeters for the dimensions and kilogram for the mass). This function must return a string: the name of the stack where the package should go.
Use Cobra for CLI
  Don't require order for variables by allowing the use of flags
  Allow default order, just don't require it
Use Logrus for logging (allow debug mode)

## Test case considerations
* Test cases should be written in a way that they are easy to understand and maintain.
* Test cases should be written in Go

### Testing scenarios
* Standard packages
* Bulky packages
* Heavy packages
* Rejected packages
* Invalid input on one field
* Invalid input on all fields
* Boundaries for Standard, Heavy, Bulky and Rejected stacks