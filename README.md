# Gobible Plain Text

A simple tool to convert a bible version from JSON to muiltiple plain text files for each book.

## Usage

```bash
go run . file.json
```

Or, if you build the binary:

```bash
gobible-plaintext file.json
```

all output will be places in `/out/`

## Example

Here is the output for the KJV:

```txt
Genesis
1:1 In the beginning God created the heaven and the earth.
1:2 And the earth was without form, and void; and darkness was upon the face of the deep. And the Spirit of God moved upon the face of the waters.
1:3 And God said, Let there be light: and there was light.
.....
```

