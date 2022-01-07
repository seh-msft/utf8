# utf8

utf8 provides unicode code point values for input runes and the unicode rune (if printable) for a given unicode code point. 

With no arguments, prints a table of unicode code points in the range [`-from`, `-to`] in the style of ascii(1). 

By default, table output suppresses non-printable characters (toggled by `-g`). 

A one-element-per-line list useful for scripting can be emitted with `-l`. 

## Build

	; go build

## Test

	; go test

## Install

	; go install

## Usage

```
» utf8 -h
Usage of utf8:
  -cols uint
        Number of columns in the table (default 5)
  -from uint
        Beginning code point for table (default 20)
  -g    Also print non-'graphic' runes in the table
  -l    list mode instead of table
  -to uint
        Ending code point for table (default 129)
»
```

## Examples

Several input forms are permitted:

```
» utf8 '\u222A' 'u222B' '222C' '2318' '0x1234' '∫' '0'
U+222A '∪'
U+222B '∫'
U+222C '∬'
U+2318 '⌘'
U+1234 'ሴ'
U+222B '∫'
U+0030 '0'
»
```

With no arguments, a table is emitted:

```
» utf8
|U+0020 ' ' |U+0021 '!' |U+0022 '"' |U+0023 '#' |U+0024 '$' |
|U+0025 '%' |U+0026 '&' |U+0027 ''' |U+0028 '(' |U+0029 ')' |
|U+0030 '0' |U+0031 '1' |U+0032 '2' |U+0033 '3' |U+0034 '4' |
|U+0035 '5' |U+0036 '6' |U+0037 '7' |U+0038 '8' |U+0039 '9' |
|U+0040 '@' |U+0041 'A' |U+0042 'B' |U+0043 'C' |U+0044 'D' |
|U+0045 'E' |U+0046 'F' |U+0047 'G' |U+0048 'H' |U+0049 'I' |
|U+0050 'P' |U+0051 'Q' |U+0052 'R' |U+0053 'S' |U+0054 'T' |
|U+0055 'U' |U+0056 'V' |U+0057 'W' |U+0058 'X' |U+0059 'Y' |
|U+0060 '`' |U+0061 'a' |U+0062 'b' |U+0063 'c' |U+0064 'd' |
|U+0065 'e' |U+0066 'f' |U+0067 'g' |U+0068 'h' |U+0069 'i' |
|U+0070 'p' |U+0071 'q' |U+0072 'r' |U+0073 's' |U+0074 't' |
|U+0075 'u' |U+0076 'v' |U+0077 'w' |U+0078 'x' |U+0079 'y' |
|U+0100 'Ā' |U+0101 'ā' |U+0102 'Ă' |U+0103 'ă' |U+0104 'Ą' |
|U+0105 'ą' |U+0106 'Ć' |U+0107 'ć' |U+0108 'Ĉ' |U+0109 'ĉ' |
|U+0110 'Đ' |U+0111 'đ' |U+0112 'Ē' |U+0113 'ē' |U+0114 'Ĕ' |
|U+0115 'ĕ' |U+0116 'Ė' |U+0117 'ė' |U+0118 'Ę' |U+0119 'ę' |
|U+0120 'Ġ' |U+0121 'ġ' |U+0122 'Ģ' |U+0123 'ģ' |U+0124 'Ĥ' |
|U+0125 'ĥ' |U+0126 'Ħ' |U+0127 'ħ' |U+0128 'Ĩ' |U+0129 'ĩ' |
»
```

A list may be emitted:

```
» utf8 -l -from 20 -to 25
U+0020 ' '
U+0021 '!'
U+0022 '"'
U+0023 '#'
U+0024 '$'
U+0025 '%'
»
```
