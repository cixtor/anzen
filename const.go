package main

// clusterPrefixLength defines the size of the subset of characters taken from
// the left side of the hashed URL that we are analyzing. The bigger the number
// the bigger the pool of servers we expect to communicate with.
//
// If the length is set to one the size of the server pool will be sixteen, one
// server for each hexadecimal character [0-9] + [a-f]. If the length is two
// then the size of the server pool will be 256, one for every combination like
// so: 00, 01, 02, ... 0a, 0b, 0c and so on until [ff].
//
// The maximum size would then be sixty-four which is the size of hashed data
// using SHA256. Because zero-zero is a valid prefix for a SHA256 sum it means
// we have to consider repetitions in the calculation of the permutations. So
// theoretically speaking, the size of the server pool would be equal to:
//
//   115,792,089,237,316,...
//   195,423,570,985,008,...
//   687,907,853,269,984,...
//   665,640,564,039,457,...
//   584,000,000,000,000,000
//
// PS: I don’t even know how to read that number :thinking_face:
const clusterPrefixLength int = 1

// ttUnspecified represents an unknown threat type.
const ttUnspecified string = "UNSPECIFIED"

// ttMalware represents a malware threat type.
const ttMalware string = "MALWARE"

// ttSocialEngineering represents a social engineering threat type.
const ttSocialEngineering string = "SOCIAL_ENGINEERING"

// ttUnwantedSoftware represents an unwanted software threat type.
const ttUnwantedSoftware string = "UNWANTED_SOFTWARE"

// ttPotentiallyHarmful represents a potentially harmful application threat type.
const ttPotentiallyHarmful string = "POTENTIALLY_HARMFUL"

// ttNone represents no malware infection.
const ttNone string = "NONE"
