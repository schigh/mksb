# mksb (work in progress)

This is a very simple and somewhat useless tool.  All it does (currently) is read a text file and converts its lines to equivalent `strings.Builder` lines, preserving the readability of the original text file.

The motivation behind this is because sql query builders are the bane of my existence and they tend to cause more problems than they solve for any reasonably complex query.

## Installation

`go get -u github.com/schigh/mksb`

## Usage

```
> mksb --file /path/to/some/text/file.txt
```

By default, the reader delimits the file using newline character `\n`, which is `10` in ASCII Decimal.  To change the input and/or output delimiters, use these flags:

```
--rd <int> # read delimiter, default 10
--wd <int> # write delimiter, default 10
```

You can specify the variable name for your `strings.Builder` instance by passing it in the `--sbName` flag

> note that currently (as of go 1.12), the `error` returned from any `strings.Builder` `Write...` function is _always_ nil, so we are ignoring them for now
----

For sql queries, use an output delimiter of `32` (space).  For example, the following query:

```sql
SELECT `col1`, `col2`, `col3`
FROM `some_table` t1
	INNER JOIN `some_other_table` t2
	ON (
	  t1.`col3` = t2.`col3`
	  AND
	  t1.`col2` IS NOT NULL
		)
ORDER BY t1.`col1` DESC
```

after running:

```
> mksb --wd 32 --file /path/to/some/query.sql
```

produces:

```go
sb.WriteString("SELECT `col1`, `col2`, `col3` ")
sb.WriteString("FROM `some_table` t1 ")
sb.WriteString(" INNER JOIN `some_other_table` t2 ")
sb.WriteString(" ON ( ")
sb.WriteString("   t1.`col3` = t2.`col3` ")
sb.WriteString("   AND ")
sb.WriteString("   t1.`col2` IS NOT NULL ")
sb.WriteString("         ) ")
sb.WriteString("ORDER BY t1.`col1` DESC")
```
