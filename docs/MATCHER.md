# Matchers

Matchers are how Lexx identifies Tokens, which Matchers are added to a Lexx instance defines what kind of Tokens that 
Lexx will find. Generally a Matcher is responsible for identifying one type of Token, however this is not an enforced 
rule. When a Matcher identifies a match it can return whichever kind of Token is appropriate.

Matchers are not Classes but are rather two defined function types:

### <a name="LexxMatcherInitialize"></a>LexxMatcherInitialize
```go
type LexxMatcherInitialize func() LexxMatcherMatch
```
LexxMatcherInitialize is used to perform any initialization a matcher needs in order to do its work, it will be
called once when GetNextToken is called and returns a function of type [LexxMatcherMatch](#LexxMatcherMatch). For example the FLOAT matcher uses
its initializing function to set a variable tracking if and where a '.' has been found in the input.
Essentially LexxMatcherInitialize is a closure that can keep information across multiple LexxMatcherMatch function calls.

### <a name="LexxMatcherMatch"></a>LexxMatcherMatch
```go
type LexxMatcherMatch func(r rune, currentText []rune) (token *Token, precedence int8, run bool)
```
LexxMatcherMatch is the function that gets called repeatedly with `runes` pulled from the input until it either
returns a Token and a `precedence` or false for the `run` value.
`r` is the next rune to match against, it will be 0 (zero) if the end of input is hit.
`currentText` is a history of runes already sent to the matcher.
The function should return a `token` of nil, `precedence` 0 and a `run` of true to keep getting called.
If `run` is returned as false the matcher will no longer be called with new runes.
To signify a match has been made the matcher should return a valid Token object, a `precedence` and false for `run`.

Example of a valid return with a Token 
```
({ Type: WORD, Value: "the", Line: 0, Column: 3 }, 1, false)
```
IT'S IMPORTANT that the value of Line and Column in the Token reflects how many lines and characters were advanced
in the match. Since "the" is three characters the Column value is set to 3.
For example if input is "abcd" this function will be called with
```
  ('a', ['a'])
  ('b', ['a','b'])
  ('c', ['a','b','c'])
  ('d', ['a','b','c','d'])
  ( 0,  ['a','b','c','d'])
```
Example Match

Given the input "The quick" the WORD matcher will take the following input and return the given results

```  
  ('T', ['T']) =>             (nil, 0, true)
  ('h', ['T',h']) =>          (nil, 0, true)
  ('e', ['T','h','e']) =>     (nil, 0, true)
  (' ', ['T','h','e',' ']) => ({ Type: WORD, Value: "The", Line: 0, Column: 3 }, 1, false)
```
Note that the WORD matcher had to wait until it did not have a matching character (the space) before it could
generate a Token. The generated token does not have to use all of the sent characters but it does have to use
characters in an unbroken array from the start of the sent list. It could not, for example, match 'he' and ignore
the 'T'.



// inputUnit is a group of input plugins and the shared channel they write to.
//
// ┌───────┐
// │ Input │───┐
// └───────┘   │
// ┌───────┐   │     ______
// │ Input │───┼──▶ ()_____)
// └───────┘   │
// ┌───────┐   │
// │ Input │───┘
// └───────┘