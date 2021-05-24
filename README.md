# ISO/IEC 7064:2003

[![GoDoc](https://godoc.org/github.com/jacoelho/go-iso7064?status.svg)](https://pkg.go.dev/github.com/jacoelho/go-iso7064?tab=overview)

Check character system implementation in [Go](https://golang.org) as described in [ISO/IEC 7064:2003](https://www.iso.org/standard/31531.html).

## Instalation

```bash
go get -u github.com/jacoelho/go-iso7064/iso7064
```

## Usage

Import:
```go
import "github.com/jacoelho/go-iso7064/iso7064"
```

Usage
```go
iso7064.Modulo11Radix2("079")
// output: "X"
```

More examples in godoc.

## Normalise

Check systems are not validating if the input is within the expected values, if required, the following function are provided:

```go
iso7064.NormaliseNumeric("1A3 4")
// output: "134"
```

```go
iso7064.NormaliseAlphabetic("aB1c")
// output: "ABC"
```

```go
iso7064.NormaliseAlphaNumeric("aB1 c")
// output: "AB1C"
```

## Implementation
Pure systems:

| Designation   | Input (string) | Output (string)          |
| ------------- | -------------- | ------------------------ |
| `MOD 11-2`    | numeric        | 1 digit or `X`           |
| `MOD 37-2`    | alpha numeric  | 1 digit or letter or `*` |
| `MOD 97-10`   | numeric        | 2 digits                 |
| `MOD 661-26`  | alphabetic     | 2 letters                |
| `MOD 1271-36` | alpha numeric  | 2 digits or letters      |

Hybrid systems:

| Designation | Input (string) | Output (string)   |
| ----------- | -------------- | ----------------- |
| `MOD 11,10` | numeric        | 1 digit           |
| `MOD 27,26` | alphabetic     | 1 letter          |
| `MOD 37,36` | alpha numeric  | 1 digit or letter |

## Credits

Some test cases were picked from:
* https://github.com/konfirm/node-iso7064

## License

GNU General Public License v3.0 or later

See [LICENSE](LICENSE) to see the full text.
