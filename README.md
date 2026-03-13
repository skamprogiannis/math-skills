# math-skills

CLI program that reads a text file (one number per line) and prints:

- Average
- Median
- Variance
- Standard Deviation

All values are printed as rounded integers.

## Requirements

- Go 1.22+

## Run

```bash
go run . <path-to-data-file>
```

Example:

```bash
go run . data/data.txt
```

## Data format

Input file must contain one value per line, for example:

```text
189
113
121
114
145
```

Blank lines are ignored. Non-numeric lines are skipped with a warning.

