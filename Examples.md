**Note to reader: The proper rendering of this document depends on the software you are using to view it.**

**It should appear correctly in the Atom editor's Markdown Preview.**

### Examples of Using cip

In the code blocks, **$** are bash prompts, and **>** are secondary prompts. All other text in code blocks is command output. For example,

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
$ cip -x 256
0x100
```

Or print it without the 0x prefix.

```
$ cip -ob=16 256
100
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
0371163
0x1f273
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

### Sequences (Ranges)

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

Print the list on one line, separating the numbers with spaces instead of newlines.

```
$ cip -s=" " -r 10
1 2 3 4 5 6 7 8 9 10
```

Count from 0 to 50 by 5s.

```
$ cip -s=' ' -r 0 50 5
0 5 10 15 20 25 30 35 40 45 50
```

Count backwards (downwards) from 100 to 90 by 2s.

```
$ cip -s ' ' -r 100 90 2
100 98 96 94 92 90
```

Print the ASCII character set, excluding whitespace and control characters.
(Notice that when you use **-c**, there are no separators between the characters.)

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
$ cip -width=4 -ic -ob=16 -H -p=U+ -r A Z
```

Character input and output (acts like **cat** command, but for standard input only).

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

Cut and paste **0** and **9** characters into this command to get their ASCII codes in decimal.

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

This was a very simple case, but the point is that you can do this with arbitrary Unicode characters that may not be available from your keyboard.

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

Count uppercase **H** characters in a file.

```
$ cip <file -U -ic | grep U+0048 | wc -l
```

Or, if you don't know the Unicode code point for uppercase **H**.

```
$ cip <file -U -ic | grep $(cip -U -ic H) | wc -l
```

Count lowercase **L** characters in **hello, world**.

```
$ echo "hello, world" | cip -U -ic | grep $(cip -U -ic l) | wc -l
3
```

Print (once each) each character in a file that has a Unicode code point made up of more than 16 bits (higher than U+FFFF).

```
$ cat file | cip -ic -U | grep 'U+[0-9A-F][0-9A-F][0-9A-F][0-9A-F][0-9A-F]' | sort | uniq | sed -e 's/U+//' | cip -ib 16 -c
```

Print the Unicode code points for each of the above characters.

```
$ cat file | cip -ic -U | grep 'U+[0-9A-F][0-9A-F][0-9A-F][0-9A-F][0-9A-F]' | sort | uniq
```

## Tutorial Examples

### Synthesis: Creating Test Cases

It can be helpful to have some weird Unicode text to use for testing software that processes text. In this example, we will create a short file containing some odd Unicode characters.

Start off by creating a short text file.

```
$ echo "Unicode isn't always as simple as it appears." >file.txt
```

Print it out to check it.

```
$ cat file.txt
Unicode isn't always as simple as it appears.
```

Check the file's size.

```
$ wc -c file.txt
46 file.txt
$ wc -m file.txt
46 file.txt
$ ls -l file.txt
-rw-r--r-- 1 jay jay 46 Feb  7 23:07 file.txt
```

(The length of the sentence is only 45 characters, but there is a newline at the end of the line, and that gets counted too.)

They all agree that the file is 46 bytes, and 46 characters, long.

Convert the Unicode characters to hex integers and save the output in another file.

```
$ cip -ic -x <file.txt >hex.txt
```

Use your favorite text editor to edit the file.

```
$ vi hex.txt
```

Add some Unicode characters:

```
U+200B  ZERO WIDTH SPACE
U+200C  ZERO WIDTH NON-JOINER
U+200D  ZERO WIDTH JOINER
U+200E  LEFT-TO-RIGHT MARK
U+200F  RIGHT-TO-LEFT MARK
U+2029  PARAGRAPH SEPARATOR
U+202A  LEFT-TO-RIGHT EMBEDDING
U+202B  RIGHT-TO-LEFT EMBEDDING
U+202C  POP DIRECTIONAL FORMATTING
U+202D  LEFT-TO-RIGHT OVERRIDE
U+202E  RIGHT-TO-LEFT OVERRIDE
```

To add them, change the format from **U+**dddd to **0x**dddd to get a result like this:

```
$ cat hex.txt
0x55
0x6e
0x69
0x63
0x6f
0x64
0x65
0x20
0x200b
0x69
0x200c
0x73
0x200d
0x6e
0x200e
0x27
0x200f
0x74
0x2029
0x20
0x202a
0x61
0x202b
0x6c
0x202c
0x77
0x202d
0x61
0x202e
0x79
0x73
0x20
0x61
0x73
0x20
0x73
0x69
0x6d
0x70
0x6c
0x65
0x20
0x61
0x73
0x20
0x69
0x74
0x20
0x61
0x70
0x70
0x65
0x61
0x72
0x73
0x2e
0xa
```

You can see where I added the characters. They have four hex digits, rather than two (or in the case of the newline at the end of the file, just one digit).

Convert the file back into Unicode characters.

```
$ cip -c <hex.txt -c >result.txt
```

Print it out.

```
$ cat result.txt
Unicode isn't  always as simple as it appears.
```

It looks the same (or almost the same, depending on what software is rendering it), but how long is it now?

```
$ wc -c result.txt
79 result.txt
$ wc -m result.txt
57 result.txt
$ ls -l result.txt
-rw-r--r-- 1 jay jay 79 Feb  8 19:35 result.txt
```

It has grown from 46 to 79 bytes long, and **wc -m** reports 57 characters! How many code points are in the file?

```
$ wc -l hex.txt
57 file.tmp
```

That counts the number of lines in the hex-format file, so **wc -m** is actually reporting the number of code points, not the number of visible characters.

Try cutting and pasting the contents of **result.txt** in your web forms and other places where your programs expect text input, and see what happens.

In this example, I used code points from the blocks starting at U+200B and U+2029 because they are often invisible (depending on the software). You can also try the following:

```
U+FE00 ... U+FE0F Variation Selectors
U+034F Combining Grapheme Joiner (Put some combining characters after it)
U+0300 ... U+036F Diacritics (Combining Characters. Use as many consecutive as you want.)
U+E0001 ... U+E0020
U+2028 LINE SEPARATOR
U+2066  LEFT-TO-RIGHT ISOLATE
U+2067  RIGHT-TO-LEFT ISOLATE
U+2068  FIRST STRONG ISOLATE
U+2069  POP DIRECTIONAL ISOLATE
U+FE00 ... U+FE0F Variation Selectors

U+FFFD Replacement Character
U+FFFF Highest 16-bit code point
U+10FFFF Highest code point allowed in Unicode
U+FEFF Byte order mark
U+FFFE Reversed byte order mark

U+212A Kelvin sign (looks like U+004B Capital K)
U+2126 Ohm sign (looks like U+03A9 Capital Omega)
```

And don't forget your favorite odd characters like emoji, heiroglyphs, and such.

### Analysis: Finding the cause of a problem.

The scenario:

A particular input to a program is causing problems.

You've isolated the problem to a short string that you copied into the file **prob.txt**. Something in that text is causing problems with your system.

You need to understand what's going on and figure out how to fix it.

To start off, show what's in the file.

```
$ cat prob.txt
açaí
```
[Advisory: That should have displayed as four characters: a, c-cedilla, a, i-acute. Unfortunately, not all web browsers handle it correctly! If it doesn't appear correctly this time, it probably won't later in this document, either.]

Well that's just a Portuguese word. Everything is international nowadays and it's not just about blueberries anymore.

Let's try to make a duplicate of that file.

```
$ echo açaí >mine.txt
$ cat mine.txt
açaí
```

Ok, super. But it's just Unicode, and it shouldn't be causing problems. You feed **mine.txt** to your program, and it works. Let's take a closer look. How large is each file?

```
$ wc -c prob.txt mine.txt
 9 prob.txt
 7 mine.txt
16 total
```
```
$ wc -m prob.txt mine.txt
 7 prob.txt
 5 mine.txt
12 total
```

(Remember, **wc -m** counts the newline at the end of the file as a character.)

They seem to have the same content, but the byte and character counts don't match!

Let's use **cip** to dump the Unicode code points in hexidecimal.

```
$ cip -ic -x <prob.txt
0x61
0x63
0x327
0x61
0x69
0x301
0xa
$ cip -ic -x <mine.txt
0x61
0xe7
0x61
0xed
0xa
```

Compare them. They have different code points.

In **mine.txt**, **0xe7** and **0xed** (U+00E7 and U+00ED) are outside the ASCII range, which goes only to 127 (0x7f in hex).

Show code points U+00E7 and U+00ED from **mine.txt**.

```
$ cip -c 0xe7 0xed
çí
```

So those are the ç and í characters.

Show code points U+0063 and U+0069 from **prob.txt**.

```
$ cip -c U+0063 U+0069
ci
```

Those are simply plain **c** and **i** characters, with no diacritical marks.

What happens when we place U+0327 after the **c** and U+0301 after the **i**?

```
$ cip -c U+0063 U+0327 U+0069 U+0301
çí
```

Oh, it makes them into **ç** and **í** characters.

Do a web search on U+0327 and U+0301, and you will learn that they are _combining characters_ that can be placed after regular characters, and modify their appearances and meanings.

Here are two ways to encode the word **açaí**:

```
$ cip -c U+0061 U+0063 U+0327 U+0061 U+0069 U+0301
açaí
$ cip -c U+0061 U+00e7 U+0061 U+00ed
açaí
```

These are both perfectly legal Unicode, and according to the Unicode specification, your program should be able to handle both!

Sometimes you may actually want to work with one or the other form, so it's good to know about _combining characters_ and _Unicode normalization forms_, and know how to handle them and convert between them in the programming language you use.

Here are some links to help you get started:

[The Go Blog - Text normalization in Go](https://blog.golang.org/normalization)

[Wikipedia - Unicode Equivalence](https://en.wikipedia.org/wiki/Unicode_equivalence#Normalization)

[Unicode Standard Annex #15 - Unicode Normalization Forms](https://unicode.org/reports/tr15/)

