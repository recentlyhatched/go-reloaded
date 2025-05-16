# Using regular expressions
This program uses the following regular expression to match the content from sample.txt which is used to break the text up into a []string slice before performing operations (e.g. uppercase, capitalize) on the text. The text is later converted into one string before being serialized into results.txt

## Breakdown of regular expresion
re := regexp.MustCompile(`\([a-z]+(,\s*\d+)?\)|[,]+|[^\s()]+`)

1. \`\` backticks at either end of the expression represents passing in a regex literal

From right to left:

2. `\(` - since opening and closing brackets have reserved functions, we use a backslash to escape it so we can match a literal opening bracket


3. [] - represents a character set where any of the characters in the sqaure brackets can be matched

4. [a-z] - macthes any character between lowercase a and lowercase z

5. [a-z]+ - + means 1 or more matches so the regex pattern matches any lowercase character 1 or more time e.g. "a" or "ab" or "abc" or "dcba"

6. () - represents a group which is used to match a pattern that involves more than one character or whitespace

7. , - matches a literal comma

8. \s - matches whitespace

9. \s* - * matches 0 or more characters - in this case, 0 or more whitespaces

10. \d - matches a digit e.g. 0-9 (could have been represented as [0-9])

11. \d+ - + matches 1 or more characters - in this case, 1 or more digits

12. `\)` - matches a literal closing bracket

13. `(,\s*\d+)?\)` - example match: (cap, 6)

14. | - means "or"

15. [,]+ - matches a comma repeated 1 or more times

16. |[^] - ^ means "not" in this case, within the character set (square bracket), preceded by | which means "or"

16. |[^\s()]+ - matches characters that are not whitespace, opening brackets or closing brackets, repeated 1 or more times - e.g. "hello!"


tokens := re.FindAllString(contentString, -1)


17. So it matches "(cap, 6)" -> or ->  "," -> or -> "hello!" - as tokens and order of the matching matters


### helpful resources
[regex tutorial](https://youtu.be/sa-TUpSx1JA?feature=shared)