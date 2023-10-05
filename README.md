# pratt

The `pratt` command is a stripped-down expression parser based on the
[Pratt parser]. To help understand Pratt parsing, the command also
provides a naive, recursive, right-associative expression parser and a naive,
iterative, left-associative expression parser.

Try it with:

```
go run ./main.go '1 + 2 * 3'
```

## Limitations

`pratt` is intended for demonstration and education purposes only. It is
purposefully thin on error handling. The input expression is expected to be a
valid, using known operators only: `<`, `+`, `-`, `*`, `/`. Operators and
operands must be separated by a single space, e.g. `1 + 2 * 3`.

[Pratt parser]: https://en.wikipedia.org/wiki/Pratt_parser

## `svg` command

The `svg` command is meant to be used with the `pratt` command. It generates
and opens an SVG image of the expression tree in the default SVG viewer.

Try it with:

```
go install ./...
pratt '1 * 2 + 3 * 4' | svg -o
```

## Installation

You can install the `pratt` command without cloning and rebuilding this repo.

### Linux and Windows

Download the [latest release] for your platform, unzip it and add `pratt`
to your path.

### macOS

Use [Homebrew] to install `pratt`.

    brew install evylang/tap/pratt

[latest release]: https://github.com/evylang/pratt/releases/latest
[Homebrew]: https://brew.sh/
