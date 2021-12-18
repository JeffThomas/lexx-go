// Package matchers contains a set of pre-made default matchers for Lexx.
package matchers

// LexxMatcherMatch is the function that gets called repeatedly with runes pulled from the input until it either
// returns a Token and a precedence or false for the run value.
//  r is the next rune to match against, it will be 0 (zero) if the end of input is hit.
//  currentText is a history of runes already sent to the matcher.
// The function should return a token of nil, precedence 0 and a run of true to keep getting called.
// If run is returned as false the matcher will no longer be called with new runes.
// To signify a match has been made the matcher should return a valid Token object, a precedence and false for run.
type LexxMatcherMatch func(r rune, currentText []rune) (token *Token, precedence int8, run bool)

// LexxMatcherInitialize is used to perform any initialization a matcher needs in order to do its work, it will be
// called once when GetNextToken is called and should return the function above. For example the FLOAT matcher uses
// its initializing function to set a variable tracking if and where a '.' has been found in the input.
// Essentially LexxMatcherInitialize is a closure that can keep information across multiple LexxMatcherMatch function calls.
type LexxMatcherInitialize func() LexxMatcherMatch
