### Examples of Using cip

In the code blocks, `$` are bash prompts, and `>` are secondary prompts. All other text in code blocks is command output. For example,

```
$ echo "hello, world"
hello, world
```

### Numeric and Unicode Conversions

Convert 377 (base 8) and 7ff (base 16) to decimal.

```
$ cip 0377 0x7ff
255
2047
```

Print the Unicode code point for the H character.

```
$ cip -U -ic H
U+0048
```

Convert 8080 (base 16) to binary.

```
$ cip -b 0x8080
1000000010000000
```

Convert 256 (decimal) to hex, and print with 0x prefix.

```
$ cip -x -p=0x 256
0x100
```

Print 3 shifted left 5 times in binary.

```
$ echo $((3 << 5)) | cip -b
1100000
```

Print 127603 in binary, octal, and hexidecimal.

```
$ for opt in -b -o -h
> do
> 	cip $opt 127603
> done
11111001001110011
371163
1f273
```

Convert Unix octal permission bits 644, 600, 755, and 700 to binary.

```
$ cip -ib=8 -b 644 600 755 700
110100100
110000000
111101101
111000000
```

or

```
$ cip -b -ib 8 644 600 755 700
```

or

```
$ cip -b 0644 0600 0755 0700
```

or

```
$ echo 644 600 755 700 | cip -ib=8 -b
```

or

```
$ echo 0644 0600 0755 0700 | cip -b
```

Print the largest unsigned 8, 16, 32, and 64-bit integers in base 10.

```
$ cip 0xff 0xffff 0xffffffff 0xffffffffffffffff
255
65535
4294967295
18446744073709551615
```

Print the highest code point allowed in Unicode.

```
$ cip 0x10ffff
114111
```

Print the hex value of Unicode Greek letters alpha and pi.

```
$ cip -U -ic απ
U+03B1
U+03C0
```

A round trip: Read Unicode characters and output decimal integers, pipe into cip and output as Unicode characters.

```
$ echo "hello, world" | cip -ic | cip -c
hello, world
```

Print C-cedillas and C-acutes.

```
$ cip -c 0xc7 0xe7 0x106 0x107
ÇçĆć
```

### Sequences

Print numbers from 1 to 10.

```
$ cip -r 10
1
2
3
4
5
6
7
8
9
10
```

Print the list one one line, separating the numbers with spaces instead of newlines.

```
$ cip -s=" " -r 10
1 2 3 4 5 6 7 8 9 10
```

Count from 0 to 50 by 5s.

```
$ cip -s=' ' -r 0 50 5
0 5 10 15 20 25 30 35 40 45 50
```

Print the ASCII character set, excluding whitespace and control characters.
(Notice that when you use -c, there are no separators between the characters.)

```
$ cip -c -r 33 127
!"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
```

or

```
$ cip -c -ic -r '!' '~'
```

Convert Unicode characters A through Z to U+XXXX code points.

```
$ cip -ic -U -r A Z
U+0041
U+0042
U+0043
U+0044
[...]
U+0057
U+0058
U+0059
U+005A
```

or

```
$ cip -width=4 -ic -X -p=U+ -r A Z
```

Character input and output (acts like cat command, but for stdin only).

```
$ cat file | cip -ic -c
```

### Short Sessions with cip

#### Working with ASCII (and Unicode) codes

Find codes for 0 and 9.

```
$ cip -c -r 32 127
 !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
```

Cut and paste '0' and '9' characters into this command to get their ASCII codes in decimal.

```
$ cip -ic 0 9
48
57
```

Print the sequence.

```
$ cip -c -r 48 57
0123456789
```

The above is the same as

```
$ cip -ic -c -r 0 9
```

This was a very simple case, but the point is that you can do this with arbitrary Unicode characters, which may not be available from your keyboard.

#### Now with Greek

Print the Greek Unicode block.

```
$ cip -c -r 0x370 0x3ff
ͰͱͲͳʹ͵Ͷͷ͸͹ͺͻͼͽ;Ϳ΀΁΂΃΄΅Ά·ΈΉΊ΋Ό΍ΎΏΐΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡ΢ΣΤΥΦΧΨΩΪΫάέήίΰαβγδεζηθικλμνξοπρςστυφχψωϊϋόύώϏϐϑϒϓϔϕϖϗϘϙϚϛϜϝϞϟϠϡϢϣϤϥϦϧϨϩϪϫϬϭϮϯϰϱϲϳϴϵ϶ϷϸϹϺϻϼϽϾϿ
```

(The range U+0370 to U+03FF can be found in the Unicode Specification. For a list, try https://unicode.org/charts/ )

Copy and paste the alpha and omega from that into this command to print the lowercase Greek alphabet.

```
$ cip -c -ic -r α ω
αβγδεζηθικλμνξοπρςστυφχψω
```

Print every 3rd character of the lowercase Greek alpahabet.

```
$ cip -c -ic -r α ω 3
αδηκνπσφω
```

Print the lowercase Greek alphabet backwards.

```
$ cip -c -ic -r ω α
ωψχφυτσςρποξνμλκιθηζεδγβα
```

### Analysis

Count 'H' characters in a file.

```
$ cip <file -U -ic | grep U+0048 | wc -l
```

Or, if you don't know the Unicode code point for H.

```
$ cip <file -U -ic | grep $(cip -U -ic H) | wc -l
```

Count lowercase 'L' characters in "hello, world".

```
$ echo "hello, world" | cip -U -ic | grep $(cip -U -ic l) | wc -l
```

Print (once each) each character in a file that has a Unicode code point made up of more than 16 bits (higher than U+FFFF).

```
$ cat file | cip -ic -U | grep 'U+[0-9A-F][0-9A-F][0-9A-F][0-9A-F][0-9A-F]' | sort | uniq | sed -e 's/U+//' | cip -ib 16 -c
```

Print the Unicode code points for each of the above characters.

```
$ cat file | cip -ic -U | grep 'U+[0-9A-F][0-9A-F][0-9A-F][0-9A-F][0-9A-F]' | sort | uniq
```
