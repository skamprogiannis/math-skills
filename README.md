# Math Skills

Math Skills is a small Go command-line program that summarizes a numeric data
set. It reads one integer per line and prints the average, median, population
variance, and population standard deviation.

The project was built as part of the Zone01 Athens curriculum to practise
statistical calculations, file parsing, and automated testing in Go.

## Features

- Reads signed integers from a text file.
- Ignores blank lines and surrounding whitespace.
- Warns about invalid lines while continuing with valid values.
- Sorts a copy of the data when calculating the median, leaving the input
  unchanged.
- Calculates population variance by dividing the squared deviations by the
  number of values (`N`).
- Rounds every reported result to the nearest integer with Go's `math.Round`;
  half values round away from zero.

## Requirements

- Go 1.22 or newer

No third-party dependencies are required.

## Run

```bash
go run . examples/data.txt
```

Expected output:

```text
Average: 136
Median: 121
Variance: 825
Standard Deviation: 29
```

Supply another text file to analyze a different data set:

```bash
go run . path/to/data.txt
```

If no file is supplied, the file cannot be opened, or the file contains no
valid integers, the program reports an error on standard error and exits with a
non-zero status. Invalid non-empty lines produce warnings on standard error.

## Input format

Input must contain one base-10 integer per line. Blank lines are ignored.

```text
189
113
121
114
145
```

## How it works

The program has two small layers:

- `main.go` handles command-line arguments, error reporting, and output.
- `internal/` reads and validates input, then calculates a `Summary` in one
  pass over the values plus a sorted copy for the median.

For values x1 through xN, variance is calculated as the population variance:

```text
variance = sum((xi - mean)^2) / N
standard deviation = sqrt(variance)
```

The unrounded mean is used for variance and standard-deviation calculations;
rounding is applied only when constructing the final summary.

## Test

Run the unit tests:

```bash
go test ./...
```

Run the same race-enabled suite used by continuous integration:

```bash
go test -race ./...
```

The tests cover input normalization, invalid input, missing files, integer
rounding, and the public calculation helpers.

## Status

The command implements the Zone01 Math Skills requirements and is maintained as
a completed educational project. Its output intentionally consists of rounded
integers rather than full-precision floating-point values.
