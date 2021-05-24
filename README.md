# ISO/IEC 7064:2003

[![GoDoc](https://godoc.org/github.com/jacoelho/go-iso7064?status.svg)](https://pkg.go.dev/github.com/jacoelho/go-iso7064?tab=overview)

Go check character system implementation as described in [ISO/IEC 7064:2003](https://www.iso.org/standard/31531.html).

## Instalation

```go
go get -u github.com/jacoelho/go-iso7064/iso7064
```

## Usage

```go
iso7064.Modulo11Radix2("079")
// output: "X"
```

Pure systems:

| Designation   | Input (string) | Output (string)          |
| ------------- | -------------- | ------------------------ |
| `MOD 11-2`    | numeric        | 1 digit or `X`           |
| `MOD 37-2`    | alpha numeric  | 1 digit or letter or `*` |
| `MOD 97-10`   | numeric        | 2 digits                 |
| `MOD 661-26`  | alphabetic     | 2 letters                |
| `MOD 1271-36` | alpha numeric  | 2 digits or letters      |

Hybrid systems:

* `MOD 11-10`
* `MOD 27-26`
* `MOD 37-36`

## Credits

Some test cases were picked from:
* https://github.com/konfirm/node-iso7064

## License

GNU General Public License v3.0 or later

See [LICENSE](LICENSE) to see the full text.
