package lexer

import (
	"unicode"

	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
)

// Defined matchers
var (
	ampersand                         = matcher.String(`&`)
	assign                            = matcher.String(`=`)
	asterisk                          = matcher.String(`*`)
	backslash                         = matcher.String(`\`)
	backtick                          = matcher.String("`")
	binaryDigit                       = matcher.String(`01`)
	binaryIndicator                   = matcher.String(`bB`)
	braceClose                        = matcher.String(`}`)
	braceOpen                         = matcher.String(`{`)
	bracketClose                      = matcher.String(`]`)
	bracketOpen                       = matcher.String(`[`)
	caret                             = matcher.String(`^`)
	dash                              = matcher.String(`-`)
	decimalDigit                      = matcher.String(`0123456789`)
	dollar                            = matcher.String(`$`)
	dot                               = matcher.String(`.`)
	doubleQuote                       = matcher.String(`"`)
	exclamationMark                   = matcher.String(`!`)
	exponentIndicator                 = matcher.String(`eE`)
	greaterThan                       = matcher.String(`>`)
	hexDigit                          = matcher.String(`0123456789abcdefABCDEF`)
	hexIndicator                      = matcher.String(`xX`)
	lessThan                          = matcher.String(`<`)
	lowerU                            = matcher.String(`u`)
	lowerX                            = matcher.String(`x`)
	nonZeroDigit                      = matcher.String(`123456789`)
	octalDigit                        = matcher.String(`01234567`)
	octalIndicator                    = matcher.String(`oO`)
	percent                           = matcher.String(`%`)
	pipe                              = matcher.String(`|`)
	plus                              = matcher.String(`+`)
	punctuatorStart                   = matcher.String(`!%&()*+,-.:;<=>?[]^{|~`)
	regularExpressionCharPartial      = matcher.Diff(regularExpressionNonTerminator, matcher.String(`\/[`))
	regularExpressionClassCharPartial = matcher.Diff(regularExpressionNonTerminator, matcher.String(`]\`))
	regularExpressionFirstCharPartial = matcher.Diff(regularExpressionNonTerminator, matcher.String(`*\/[`))
	regularExpressionNonTerminator    = matcher.Diff(sourceCharacter, lineTerminator)
	sign                              = matcher.String(`+-`)
	singleQuote                       = matcher.String(`'`)
	slash                             = matcher.String(`/`)
	templateCharacterPartial          = matcher.Diff(sourceCharacter, matcher.Merge(backtick, backslash, dollar, lineTerminator))
	underscore                        = matcher.String(`_`)
	zero                              = matcher.String(`0`)

	sourceCharacter = matcher.New("SourceCharacter", codePoint) // 10.1

	whiteSpace            = matcher.Merge(formFeed, noBreakSpace, space, horizontalTab, unicodeSpace, verticalTab, zeroWidthJoiner, zeroWidthNoBreakSpace, zeroWidthNonJoiner) // 11.2
	formFeed              = matcher.RuneWithDesc("<FF>", '\u000C')                                                                                                             // 11.2
	noBreakSpace          = matcher.RuneWithDesc("<NBSP>", '\u00A0')                                                                                                           // 11.2
	space                 = matcher.RuneWithDesc("<SP>", '\u0020')                                                                                                             // 11.2
	horizontalTab         = matcher.RuneWithDesc("<TAB>", '\u0009')                                                                                                            // 11.2
	unicodeSpace          = matcher.New("<USP>", unicode.Z)                                                                                                                    // 11.2
	verticalTab           = matcher.RuneWithDesc("<VT>", '\u000B')                                                                                                             // 11.2
	zeroWidthJoiner       = matcher.RuneWithDesc("<ZWJ>", '\u200D')                                                                                                            // 11.1
	zeroWidthNoBreakSpace = matcher.RuneWithDesc("<ZWNBSP>", '\uFEFF')                                                                                                         // 11.1
	zeroWidthNonJoiner    = matcher.RuneWithDesc("<ZWNJ>", '\u200C')                                                                                                           // 11.1

	lineTerminator     = matcher.Merge(lineFeed, carriageReturn, lineSeparator, paragraphSeparator) // 11.3
	lineFeed           = matcher.RuneWithDesc("<LF>", '\u000A')                                     // 11.3
	carriageReturn     = matcher.RuneWithDesc("<CR>", '\u000D')                                     // 11.3
	lineSeparator      = matcher.RuneWithDesc("<LS>", '\u2028')                                     // 11.3
	paragraphSeparator = matcher.RuneWithDesc("<PS>", '\u2029')                                     // 11.3

	singleLineCommentChar = matcher.Diff(sourceCharacter, lineTerminator) // 11.4

	multiLineNotAsteriskChar               = matcher.Diff(sourceCharacter, asterisk)
	multiLineNotForwardSlashOrAsteriskChar = matcher.Diff(sourceCharacter, matcher.Merge(asterisk, slash))

	unicodeIDStart         = matcher.New("ID_Start", unicode.Other_ID_Start)
	unicodeIDContinue      = matcher.New("ID_Continue", iDContinue)
	identifierStartPartial = matcher.Merge(unicodeIDStart, dollar, underscore)
	identifierStart        = matcher.Merge(unicodeIDStart, dollar, underscore, backslash)
	identifierPartPartial  = matcher.Merge(unicodeIDContinue, dollar, zeroWidthNonJoiner, zeroWidthJoiner)
	identifierPart         = matcher.Merge(unicodeIDContinue, dollar, zeroWidthNonJoiner, zeroWidthJoiner, backslash)

	doubleStringCharacterPartial = matcher.Diff(sourceCharacter, matcher.Merge(doubleQuote, backslash, lineTerminator))
	singleStringCharacterPartial = matcher.Diff(sourceCharacter, matcher.Merge(singleQuote, backslash, lineTerminator))

	singleEscapeCharacter = matcher.String(`'"\bfnrtv`)
	escapeCharacter       = matcher.Merge(singleEscapeCharacter, decimalDigit, lowerX, lowerU)
	nonEscapeCharacter    = matcher.Diff(sourceCharacter, matcher.Merge(escapeCharacter, lineTerminator))
)

var (
	// codePoint contains all runes that are >=0 and <=unicode.MaxRune
	codePoint = &unicode.RangeTable{
		R32: []unicode.Range32{
			{0, unicode.MaxRune, 1},
		},
	}
	// iDContinue contains all runes with the property [:IDContinue:]
	iDContinue = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005a, 1}, // A-Z
			{0x005f, 0x005f, 1}, // _
			{0x0061, 0x007a, 1}, // a-z
			{0x00aa, 0x00aa, 1},
			{0x00b5, 0x00b5, 1},
			{0x00b7, 0x00b7, 1},
			{0x00ba, 0x00ba, 1},
			{0x00c0, 0x00d6, 1},
			{0x00d8, 0x00f6, 1},
			{0x00f8, 0x02c1, 1},
			{0x02c6, 0x02d1, 1},
			{0x02e0, 0x02e4, 1},
			{0x02ec, 0x02ec, 1},
			{0x02ee, 0x02ee, 1},
			{0x0300, 0x0374, 1},
			{0x0376, 0x0376, 1},
			{0x0377, 0x0377, 1},
			{0x037a, 0x037d, 1},
			{0x037f, 0x037f, 1},
			{0x0386, 0x038a, 1},
			{0x038c, 0x038c, 1},
			{0x038e, 0x03a1, 1},
			{0x03a3, 0x03f5, 1},
			{0x03f7, 0x0481, 1},
			{0x0483, 0x0487, 1},
			{0x048a, 0x052f, 1},
			{0x0531, 0x0556, 1},
			{0x0559, 0x0559, 1},
			{0x0560, 0x0588, 1},
			{0x0591, 0x05bd, 1},
			{0x05bf, 0x05bf, 1},
			{0x05c1, 0x05c1, 1},
			{0x05c2, 0x05c2, 1},
			{0x05c4, 0x05c4, 1},
			{0x05c5, 0x05c5, 1},
			{0x05c7, 0x05c7, 1},
			{0x05d0, 0x05ea, 1},
			{0x05ef, 0x05f2, 1},
			{0x0610, 0x061a, 1},
			{0x0620, 0x0669, 1},
			{0x066e, 0x06d3, 1},
			{0x06d5, 0x06dc, 1},
			{0x06df, 0x06e8, 1},
			{0x06ea, 0x06fc, 1},
			{0x06ff, 0x06ff, 1},
			{0x0710, 0x074a, 1},
			{0x074d, 0x07b1, 1},
			{0x07c0, 0x07f5, 1},
			{0x07fa, 0x07fa, 1},
			{0x07fd, 0x07fd, 1},
			{0x0800, 0x082d, 1},
			{0x0840, 0x085b, 1},
			{0x0860, 0x086a, 1},
			{0x08a0, 0x08b4, 1},
			{0x08b6, 0x08bd, 1},
			{0x08d3, 0x08e1, 1},
			{0x08e3, 0x0963, 1},
			{0x0966, 0x096f, 1},
			{0x0971, 0x0983, 1},
			{0x0985, 0x098c, 1},
			{0x098f, 0x098f, 1},
			{0x0990, 0x0990, 1},
			{0x0993, 0x09a8, 1},
			{0x09aa, 0x09b0, 1},
			{0x09b2, 0x09b2, 1},
			{0x09b6, 0x09b9, 1},
			{0x09bc, 0x09c4, 1},
			{0x09c7, 0x09c7, 1},
			{0x09c8, 0x09c8, 1},
			{0x09cb, 0x09ce, 1},
			{0x09d7, 0x09d7, 1},
			{0x09dc, 0x09dc, 1},
			{0x09dd, 0x09dd, 1},
			{0x09df, 0x09e3, 1},
			{0x09e6, 0x09f1, 1},
			{0x09fc, 0x09fc, 1},
			{0x09fe, 0x09fe, 1},
			{0x0a01, 0x0a03, 1},
			{0x0a05, 0x0a0a, 1},
			{0x0a0f, 0x0a0f, 1},
			{0x0a10, 0x0a10, 1},
			{0x0a13, 0x0a28, 1},
			{0x0a2a, 0x0a30, 1},
			{0x0a32, 0x0a32, 1},
			{0x0a33, 0x0a33, 1},
			{0x0a35, 0x0a35, 1},
			{0x0a36, 0x0a36, 1},
			{0x0a38, 0x0a38, 1},
			{0x0a39, 0x0a39, 1},
			{0x0a3c, 0x0a3c, 1},
			{0x0a3e, 0x0a42, 1},
			{0x0a47, 0x0a47, 1},
			{0x0a48, 0x0a48, 1},
			{0x0a4b, 0x0a4d, 1},
			{0x0a51, 0x0a51, 1},
			{0x0a59, 0x0a5c, 1},
			{0x0a5e, 0x0a5e, 1},
			{0x0a66, 0x0a75, 1},
			{0x0a81, 0x0a83, 1},
			{0x0a85, 0x0a8d, 1},
			{0x0a8f, 0x0a91, 1},
			{0x0a93, 0x0aa8, 1},
			{0x0aaa, 0x0ab0, 1},
			{0x0ab2, 0x0ab2, 1},
			{0x0ab3, 0x0ab3, 1},
			{0x0ab5, 0x0ab9, 1},
			{0x0abc, 0x0ac5, 1},
			{0x0ac7, 0x0ac9, 1},
			{0x0acb, 0x0acd, 1},
			{0x0ad0, 0x0ad0, 1},
			{0x0ae0, 0x0ae3, 1},
			{0x0ae6, 0x0aef, 1},
			{0x0af9, 0x0aff, 1},
			{0x0b01, 0x0b03, 1},
			{0x0b05, 0x0b0c, 1},
			{0x0b0f, 0x0b0f, 1},
			{0x0b10, 0x0b10, 1},
			{0x0b13, 0x0b28, 1},
			{0x0b2a, 0x0b30, 1},
			{0x0b32, 0x0b32, 1},
			{0x0b33, 0x0b33, 1},
			{0x0b35, 0x0b39, 1},
			{0x0b3c, 0x0b44, 1},
			{0x0b47, 0x0b47, 1},
			{0x0b48, 0x0b48, 1},
			{0x0b4b, 0x0b4d, 1},
			{0x0b56, 0x0b56, 1},
			{0x0b57, 0x0b57, 1},
			{0x0b5c, 0x0b5c, 1},
			{0x0b5d, 0x0b5d, 1},
			{0x0b5f, 0x0b63, 1},
			{0x0b66, 0x0b6f, 1},
			{0x0b71, 0x0b71, 1},
			{0x0b82, 0x0b82, 1},
			{0x0b83, 0x0b83, 1},
			{0x0b85, 0x0b8a, 1},
			{0x0b8e, 0x0b90, 1},
			{0x0b92, 0x0b95, 1},
			{0x0b99, 0x0b99, 1},
			{0x0b9a, 0x0b9a, 1},
			{0x0b9c, 0x0b9c, 1},
			{0x0b9e, 0x0b9e, 1},
			{0x0b9f, 0x0b9f, 1},
			{0x0ba3, 0x0ba3, 1},
			{0x0ba4, 0x0ba4, 1},
			{0x0ba8, 0x0baa, 1},
			{0x0bae, 0x0bb9, 1},
			{0x0bbe, 0x0bc2, 1},
			{0x0bc6, 0x0bc8, 1},
			{0x0bca, 0x0bcd, 1},
			{0x0bd0, 0x0bd0, 1},
			{0x0bd7, 0x0bd7, 1},
			{0x0be6, 0x0bef, 1},
			{0x0c00, 0x0c0c, 1},
			{0x0c0e, 0x0c10, 1},
			{0x0c12, 0x0c28, 1},
			{0x0c2a, 0x0c39, 1},
			{0x0c3d, 0x0c44, 1},
			{0x0c46, 0x0c48, 1},
			{0x0c4a, 0x0c4d, 1},
			{0x0c55, 0x0c55, 1},
			{0x0c56, 0x0c56, 1},
			{0x0c58, 0x0c5a, 1},
			{0x0c60, 0x0c63, 1},
			{0x0c66, 0x0c6f, 1},
			{0x0c80, 0x0c83, 1},
			{0x0c85, 0x0c8c, 1},
			{0x0c8e, 0x0c90, 1},
			{0x0c92, 0x0ca8, 1},
			{0x0caa, 0x0cb3, 1},
			{0x0cb5, 0x0cb9, 1},
			{0x0cbc, 0x0cc4, 1},
			{0x0cc6, 0x0cc8, 1},
			{0x0cca, 0x0ccd, 1},
			{0x0cd5, 0x0cd5, 1},
			{0x0cd6, 0x0cd6, 1},
			{0x0cde, 0x0cde, 1},
			{0x0ce0, 0x0ce3, 1},
			{0x0ce6, 0x0cef, 1},
			{0x0cf1, 0x0cf1, 1},
			{0x0cf2, 0x0cf2, 1},
			{0x0d00, 0x0d03, 1},
			{0x0d05, 0x0d0c, 1},
			{0x0d0e, 0x0d10, 1},
			{0x0d12, 0x0d44, 1},
			{0x0d46, 0x0d48, 1},
			{0x0d4a, 0x0d4e, 1},
			{0x0d54, 0x0d57, 1},
			{0x0d5f, 0x0d63, 1},
			{0x0d66, 0x0d6f, 1},
			{0x0d7a, 0x0d7f, 1},
			{0x0d82, 0x0d82, 1},
			{0x0d83, 0x0d83, 1},
			{0x0d85, 0x0d96, 1},
			{0x0d9a, 0x0db1, 1},
			{0x0db3, 0x0dbb, 1},
			{0x0dbd, 0x0dbd, 1},
			{0x0dc0, 0x0dc6, 1},
			{0x0dca, 0x0dca, 1},
			{0x0dcf, 0x0dd4, 1},
			{0x0dd6, 0x0dd6, 1},
			{0x0dd8, 0x0ddf, 1},
			{0x0de6, 0x0def, 1},
			{0x0df2, 0x0df2, 1},
			{0x0df3, 0x0df3, 1},
			{0x0e01, 0x0e3a, 1},
			{0x0e40, 0x0e4e, 1},
			{0x0e50, 0x0e59, 1},
			{0x0e81, 0x0e81, 1},
			{0x0e82, 0x0e82, 1},
			{0x0e84, 0x0e84, 1},
			{0x0e86, 0x0e8a, 1},
			{0x0e8c, 0x0ea3, 1},
			{0x0ea5, 0x0ea5, 1},
			{0x0ea7, 0x0ebd, 1},
			{0x0ec0, 0x0ec4, 1},
			{0x0ec6, 0x0ec6, 1},
			{0x0ec8, 0x0ecd, 1},
			{0x0ed0, 0x0ed9, 1},
			{0x0edc, 0x0edf, 1},
			{0x0f00, 0x0f00, 1},
			{0x0f18, 0x0f18, 1},
			{0x0f19, 0x0f19, 1},
			{0x0f20, 0x0f29, 1},
			{0x0f35, 0x0f35, 1},
			{0x0f37, 0x0f37, 1},
			{0x0f39, 0x0f39, 1},
			{0x0f3e, 0x0f47, 1},
			{0x0f49, 0x0f6c, 1},
			{0x0f71, 0x0f84, 1},
			{0x0f86, 0x0f97, 1},
			{0x0f99, 0x0fbc, 1},
			{0x0fc6, 0x0fc6, 1},
			{0x1000, 0x1049, 1},
			{0x1050, 0x109d, 1},
			{0x10a0, 0x10c5, 1},
			{0x10c7, 0x10c7, 1},
			{0x10cd, 0x10cd, 1},
			{0x10d0, 0x10fa, 1},
			{0x10fc, 0x1248, 1},
			{0x124a, 0x124d, 1},
			{0x1250, 0x1256, 1},
			{0x1258, 0x1258, 1},
			{0x125a, 0x125d, 1},
			{0x1260, 0x1288, 1},
			{0x128a, 0x128d, 1},
			{0x1290, 0x12b0, 1},
			{0x12b2, 0x12b5, 1},
			{0x12b8, 0x12be, 1},
			{0x12c0, 0x12c0, 1},
			{0x12c2, 0x12c5, 1},
			{0x12c8, 0x12d6, 1},
			{0x12d8, 0x1310, 1},
			{0x1312, 0x1315, 1},
			{0x1318, 0x135a, 1},
			{0x135d, 0x135f, 1},
			{0x1369, 0x1371, 1},
			{0x1380, 0x138f, 1},
			{0x13a0, 0x13f5, 1},
			{0x13f8, 0x13fd, 1},
			{0x1401, 0x166c, 1},
			{0x166f, 0x167f, 1},
			{0x1681, 0x169a, 1},
			{0x16a0, 0x16ea, 1},
			{0x16ee, 0x16f8, 1},
			{0x1700, 0x170c, 1},
			{0x170e, 0x1714, 1},
			{0x1720, 0x1734, 1},
			{0x1740, 0x1753, 1},
			{0x1760, 0x176c, 1},
			{0x176e, 0x1770, 1},
			{0x1772, 0x1772, 1},
			{0x1773, 0x1773, 1},
			{0x1780, 0x17d3, 1},
			{0x17d7, 0x17d7, 1},
			{0x17dc, 0x17dc, 1},
			{0x17dd, 0x17dd, 1},
			{0x17e0, 0x17e9, 1},
			{0x180b, 0x180d, 1},
			{0x1810, 0x1819, 1},
			{0x1820, 0x1878, 1},
			{0x1880, 0x18aa, 1},
			{0x18b0, 0x18f5, 1},
			{0x1900, 0x191e, 1},
			{0x1920, 0x192b, 1},
			{0x1930, 0x193b, 1},
			{0x1946, 0x196d, 1},
			{0x1970, 0x1974, 1},
			{0x1980, 0x19ab, 1},
			{0x19b0, 0x19c9, 1},
			{0x19d0, 0x19da, 1},
			{0x1a00, 0x1a1b, 1},
			{0x1a20, 0x1a5e, 1},
			{0x1a60, 0x1a7c, 1},
			{0x1a7f, 0x1a89, 1},
			{0x1a90, 0x1a99, 1},
			{0x1aa7, 0x1aa7, 1},
			{0x1ab0, 0x1abd, 1},
			{0x1b00, 0x1b4b, 1},
			{0x1b50, 0x1b59, 1},
			{0x1b6b, 0x1b73, 1},
			{0x1b80, 0x1bf3, 1},
			{0x1c00, 0x1c37, 1},
			{0x1c40, 0x1c49, 1},
			{0x1c4d, 0x1c7d, 1},
			{0x1c80, 0x1c88, 1},
			{0x1c90, 0x1cba, 1},
			{0x1cbd, 0x1cbf, 1},
			{0x1cd0, 0x1cd2, 1},
			{0x1cd4, 0x1cfa, 1},
			{0x1d00, 0x1df9, 1},
			{0x1dfb, 0x1f15, 1},
			{0x1f18, 0x1f1d, 1},
			{0x1f20, 0x1f45, 1},
			{0x1f48, 0x1f4d, 1},
			{0x1f50, 0x1f57, 1},
			{0x1f59, 0x1f59, 1},
			{0x1f5b, 0x1f5b, 1},
			{0x1f5d, 0x1f5d, 1},
			{0x1f5f, 0x1f7d, 1},
			{0x1f80, 0x1fb4, 1},
			{0x1fb6, 0x1fbc, 1},
			{0x1fbe, 0x1fbe, 1},
			{0x1fc2, 0x1fc4, 1},
			{0x1fc6, 0x1fcc, 1},
			{0x1fd0, 0x1fd3, 1},
			{0x1fd6, 0x1fdb, 1},
			{0x1fe0, 0x1fec, 1},
			{0x1ff2, 0x1ff4, 1},
			{0x1ff6, 0x1ffc, 1},
			{0x203f, 0x203f, 1},
			{0x2040, 0x2040, 1},
			{0x2054, 0x2054, 1},
			{0x2071, 0x2071, 1},
			{0x207f, 0x207f, 1},
			{0x2090, 0x209c, 1},
			{0x20d0, 0x20dc, 1},
			{0x20e1, 0x20e1, 1},
			{0x20e5, 0x20f0, 1},
			{0x2102, 0x2102, 1},
			{0x2107, 0x2107, 1},
			{0x210a, 0x2113, 1},
			{0x2115, 0x2115, 1},
			{0x2118, 0x211d, 1},
			{0x2124, 0x2124, 1},
			{0x2126, 0x2126, 1},
			{0x2128, 0x2128, 1},
			{0x212a, 0x2139, 1},
			{0x213c, 0x213f, 1},
			{0x2145, 0x2149, 1},
			{0x214e, 0x214e, 1},
			{0x2160, 0x2188, 1},
			{0x2c00, 0x2c2e, 1},
			{0x2c30, 0x2c5e, 1},
			{0x2c60, 0x2ce4, 1},
			{0x2ceb, 0x2cf3, 1},
			{0x2d00, 0x2d25, 1},
			{0x2d27, 0x2d27, 1},
			{0x2d2d, 0x2d2d, 1},
			{0x2d30, 0x2d67, 1},
			{0x2d6f, 0x2d6f, 1},
			{0x2d7f, 0x2d96, 1},
			{0x2da0, 0x2da6, 1},
			{0x2da8, 0x2dae, 1},
			{0x2db0, 0x2db6, 1},
			{0x2db8, 0x2dbe, 1},
			{0x2dc0, 0x2dc6, 1},
			{0x2dc8, 0x2dce, 1},
			{0x2dd0, 0x2dd6, 1},
			{0x2dd8, 0x2dde, 1},
			{0x2de0, 0x2dff, 1},
			{0x3005, 0x3007, 1},
			{0x3021, 0x302f, 1},
			{0x3031, 0x3035, 1},
			{0x3038, 0x303c, 1},
			{0x3041, 0x3096, 1},
			{0x3099, 0x309f, 1},
			{0x30a1, 0x30fa, 1},
			{0x30fc, 0x30ff, 1},
			{0x3105, 0x312f, 1},
			{0x3131, 0x318e, 1},
			{0x31a0, 0x31ba, 1},
			{0x31f0, 0x31ff, 1},
			{0x3400, 0x4db5, 1},
			{0x4e00, 0x9fef, 1},
			{0xa000, 0xa48c, 1},
			{0xa4d0, 0xa4fd, 1},
			{0xa500, 0xa60c, 1},
			{0xa610, 0xa62b, 1},
			{0xa640, 0xa66f, 1},
			{0xa674, 0xa67d, 1},
			{0xa67f, 0xa6f1, 1},
			{0xa717, 0xa71f, 1},
			{0xa722, 0xa788, 1},
			{0xa78b, 0xa7bf, 1},
			{0xa7c2, 0xa7c6, 1},
			{0xa7f7, 0xa827, 1},
			{0xa840, 0xa873, 1},
			{0xa880, 0xa8c5, 1},
			{0xa8d0, 0xa8d9, 1},
			{0xa8e0, 0xa8f7, 1},
			{0xa8fb, 0xa8fb, 1},
			{0xa8fd, 0xa92d, 1},
			{0xa930, 0xa953, 1},
			{0xa960, 0xa97c, 1},
			{0xa980, 0xa9c0, 1},
			{0xa9cf, 0xa9d9, 1},
			{0xa9e0, 0xa9fe, 1},
			{0xaa00, 0xaa36, 1},
			{0xaa40, 0xaa4d, 1},
			{0xaa50, 0xaa59, 1},
			{0xaa60, 0xaa76, 1},
			{0xaa7a, 0xaac2, 1},
			{0xaadb, 0xaadd, 1},
			{0xaae0, 0xaaef, 1},
			{0xaaf2, 0xaaf6, 1},
			{0xab01, 0xab06, 1},
			{0xab09, 0xab0e, 1},
			{0xab11, 0xab16, 1},
			{0xab20, 0xab26, 1},
			{0xab28, 0xab2e, 1},
			{0xab30, 0xab5a, 1},
			{0xab5c, 0xab67, 1},
			{0xab70, 0xabea, 1},
			{0xabec, 0xabec, 1},
			{0xabed, 0xabed, 1},
			{0xabf0, 0xabf9, 1},
			{0xac00, 0xd7a3, 1},
			{0xd7b0, 0xd7c6, 1},
			{0xd7cb, 0xd7fb, 1},
			{0xf900, 0xfa6d, 1},
			{0xfa70, 0xfad9, 1},
			{0xfb00, 0xfb06, 1},
			{0xfb13, 0xfb17, 1},
			{0xfb1d, 0xfb28, 1},
			{0xfb2a, 0xfb36, 1},
			{0xfb38, 0xfb3c, 1},
			{0xfb3e, 0xfb3e, 1},
			{0xfb40, 0xfb40, 1},
			{0xfb41, 0xfb41, 1},
			{0xfb43, 0xfb43, 1},
			{0xfb44, 0xfb44, 1},
			{0xfb46, 0xfbb1, 1},
			{0xfbd3, 0xfd3d, 1},
			{0xfd50, 0xfd8f, 1},
			{0xfd92, 0xfdc7, 1},
			{0xfdf0, 0xfdfb, 1},
			{0xfe00, 0xfe0f, 1},
			{0xfe20, 0xfe2f, 1},
			{0xfe33, 0xfe33, 1},
			{0xfe34, 0xfe34, 1},
			{0xfe4d, 0xfe4f, 1},
			{0xfe70, 0xfe74, 1},
			{0xfe76, 0xfefc, 1},
			{0xff10, 0xff19, 1},
			{0xff21, 0xff3a, 1},
			{0xff3f, 0xff3f, 1},
			{0xff41, 0xff5a, 1},
			{0xff66, 0xffbe, 1},
			{0xffc2, 0xffc7, 1},
			{0xffca, 0xffcf, 1},
			{0xffd2, 0xffd7, 1},
			{0xffda, 0xffdc, 1},
		},
		R32: []unicode.Range32{
			{0x00010000, 0x0001000b, 1},
			{0x0001000d, 0x00010026, 1},
			{0x00010028, 0x0001003a, 1},
			{0x0001003c, 0x0001003c, 1},
			{0x0001003d, 0x0001003d, 1},
			{0x0001003f, 0x0001004d, 1},
			{0x00010050, 0x0001005d, 1},
			{0x00010080, 0x000100fa, 1},
			{0x00010140, 0x00010174, 1},
			{0x000101fd, 0x000101fd, 1},
			{0x00010280, 0x0001029c, 1},
			{0x000102a0, 0x000102d0, 1},
			{0x000102e0, 0x000102e0, 1},
			{0x00010300, 0x0001031f, 1},
			{0x0001032d, 0x0001034a, 1},
			{0x00010350, 0x0001037a, 1},
			{0x00010380, 0x0001039d, 1},
			{0x000103a0, 0x000103c3, 1},
			{0x000103c8, 0x000103cf, 1},
			{0x000103d1, 0x000103d5, 1},
			{0x00010400, 0x0001049d, 1},
			{0x000104a0, 0x000104a9, 1},
			{0x000104b0, 0x000104d3, 1},
			{0x000104d8, 0x000104fb, 1},
			{0x00010500, 0x00010527, 1},
			{0x00010530, 0x00010563, 1},
			{0x00010600, 0x00010736, 1},
			{0x00010740, 0x00010755, 1},
			{0x00010760, 0x00010767, 1},
			{0x00010800, 0x00010805, 1},
			{0x00010808, 0x00010808, 1},
			{0x0001080a, 0x00010835, 1},
			{0x00010837, 0x00010837, 1},
			{0x00010838, 0x00010838, 1},
			{0x0001083c, 0x0001083c, 1},
			{0x0001083f, 0x00010855, 1},
			{0x00010860, 0x00010876, 1},
			{0x00010880, 0x0001089e, 1},
			{0x000108e0, 0x000108f2, 1},
			{0x000108f4, 0x000108f4, 1},
			{0x000108f5, 0x000108f5, 1},
			{0x00010900, 0x00010915, 1},
			{0x00010920, 0x00010939, 1},
			{0x00010980, 0x000109b7, 1},
			{0x000109be, 0x000109be, 1},
			{0x000109bf, 0x000109bf, 1},
			{0x00010a00, 0x00010a03, 1},
			{0x00010a05, 0x00010a05, 1},
			{0x00010a06, 0x00010a06, 1},
			{0x00010a0c, 0x00010a13, 1},
			{0x00010a15, 0x00010a17, 1},
			{0x00010a19, 0x00010a35, 1},
			{0x00010a38, 0x00010a3a, 1},
			{0x00010a3f, 0x00010a3f, 1},
			{0x00010a60, 0x00010a7c, 1},
			{0x00010a80, 0x00010a9c, 1},
			{0x00010ac0, 0x00010ac7, 1},
			{0x00010ac9, 0x00010ae6, 1},
			{0x00010b00, 0x00010b35, 1},
			{0x00010b40, 0x00010b55, 1},
			{0x00010b60, 0x00010b72, 1},
			{0x00010b80, 0x00010b91, 1},
			{0x00010c00, 0x00010c48, 1},
			{0x00010c80, 0x00010cb2, 1},
			{0x00010cc0, 0x00010cf2, 1},
			{0x00010d00, 0x00010d27, 1},
			{0x00010d30, 0x00010d39, 1},
			{0x00010f00, 0x00010f1c, 1},
			{0x00010f27, 0x00010f27, 1},
			{0x00010f30, 0x00010f50, 1},
			{0x00010fe0, 0x00010ff6, 1},
			{0x00011000, 0x00011046, 1},
			{0x00011066, 0x0001106f, 1},
			{0x0001107f, 0x000110ba, 1},
			{0x000110d0, 0x000110e8, 1},
			{0x000110f0, 0x000110f9, 1},
			{0x00011100, 0x00011134, 1},
			{0x00011136, 0x0001113f, 1},
			{0x00011144, 0x00011146, 1},
			{0x00011150, 0x00011173, 1},
			{0x00011176, 0x00011176, 1},
			{0x00011180, 0x000111c4, 1},
			{0x000111c9, 0x000111cc, 1},
			{0x000111d0, 0x000111da, 1},
			{0x000111dc, 0x000111dc, 1},
			{0x00011200, 0x00011211, 1},
			{0x00011213, 0x00011237, 1},
			{0x0001123e, 0x0001123e, 1},
			{0x00011280, 0x00011286, 1},
			{0x00011288, 0x00011288, 1},
			{0x0001128a, 0x0001128d, 1},
			{0x0001128f, 0x0001129d, 1},
			{0x0001129f, 0x000112a8, 1},
			{0x000112b0, 0x000112ea, 1},
			{0x000112f0, 0x000112f9, 1},
			{0x00011300, 0x00011303, 1},
			{0x00011305, 0x0001130c, 1},
			{0x0001130f, 0x0001130f, 1},
			{0x00011310, 0x00011310, 1},
			{0x00011313, 0x00011328, 1},
			{0x0001132a, 0x00011330, 1},
			{0x00011332, 0x00011332, 1},
			{0x00011333, 0x00011333, 1},
			{0x00011335, 0x00011339, 1},
			{0x0001133b, 0x00011344, 1},
			{0x00011347, 0x00011347, 1},
			{0x00011348, 0x00011348, 1},
			{0x0001134b, 0x0001134d, 1},
			{0x00011350, 0x00011350, 1},
			{0x00011357, 0x00011357, 1},
			{0x0001135d, 0x00011363, 1},
			{0x00011366, 0x0001136c, 1},
			{0x00011370, 0x00011374, 1},
			{0x00011400, 0x0001144a, 1},
			{0x00011450, 0x00011459, 1},
			{0x0001145e, 0x0001145e, 1},
			{0x0001145f, 0x0001145f, 1},
			{0x00011480, 0x000114c5, 1},
			{0x000114c7, 0x000114c7, 1},
			{0x000114d0, 0x000114d9, 1},
			{0x00011580, 0x000115b5, 1},
			{0x000115b8, 0x000115c0, 1},
			{0x000115d8, 0x000115dd, 1},
			{0x00011600, 0x00011640, 1},
			{0x00011644, 0x00011644, 1},
			{0x00011650, 0x00011659, 1},
			{0x00011680, 0x000116b8, 1},
			{0x000116c0, 0x000116c9, 1},
			{0x00011700, 0x0001171a, 1},
			{0x0001171d, 0x0001172b, 1},
			{0x00011730, 0x00011739, 1},
			{0x00011800, 0x0001183a, 1},
			{0x000118a0, 0x000118e9, 1},
			{0x000118ff, 0x000118ff, 1},
			{0x000119a0, 0x000119a7, 1},
			{0x000119aa, 0x000119d7, 1},
			{0x000119da, 0x000119e1, 1},
			{0x000119e3, 0x000119e3, 1},
			{0x000119e4, 0x000119e4, 1},
			{0x00011a00, 0x00011a3e, 1},
			{0x00011a47, 0x00011a47, 1},
			{0x00011a50, 0x00011a99, 1},
			{0x00011a9d, 0x00011a9d, 1},
			{0x00011ac0, 0x00011af8, 1},
			{0x00011c00, 0x00011c08, 1},
			{0x00011c0a, 0x00011c36, 1},
			{0x00011c38, 0x00011c40, 1},
			{0x00011c50, 0x00011c59, 1},
			{0x00011c72, 0x00011c8f, 1},
			{0x00011c92, 0x00011ca7, 1},
			{0x00011ca9, 0x00011cb6, 1},
			{0x00011d00, 0x00011d06, 1},
			{0x00011d08, 0x00011d08, 1},
			{0x00011d09, 0x00011d09, 1},
			{0x00011d0b, 0x00011d36, 1},
			{0x00011d3a, 0x00011d3a, 1},
			{0x00011d3c, 0x00011d3c, 1},
			{0x00011d3d, 0x00011d3d, 1},
			{0x00011d3f, 0x00011d47, 1},
			{0x00011d50, 0x00011d59, 1},
			{0x00011d60, 0x00011d65, 1},
			{0x00011d67, 0x00011d67, 1},
			{0x00011d68, 0x00011d68, 1},
			{0x00011d6a, 0x00011d8e, 1},
			{0x00011d90, 0x00011d90, 1},
			{0x00011d91, 0x00011d91, 1},
			{0x00011d93, 0x00011d98, 1},
			{0x00011da0, 0x00011da9, 1},
			{0x00011ee0, 0x00011ef6, 1},
			{0x00012000, 0x00012399, 1},
			{0x00012400, 0x0001246e, 1},
			{0x00012480, 0x00012543, 1},
			{0x00013000, 0x0001342e, 1},
			{0x00014400, 0x00014646, 1},
			{0x00016800, 0x00016a38, 1},
			{0x00016a40, 0x00016a5e, 1},
			{0x00016a60, 0x00016a69, 1},
			{0x00016ad0, 0x00016aed, 1},
			{0x00016af0, 0x00016af4, 1},
			{0x00016b00, 0x00016b36, 1},
			{0x00016b40, 0x00016b43, 1},
			{0x00016b50, 0x00016b59, 1},
			{0x00016b63, 0x00016b77, 1},
			{0x00016b7d, 0x00016b8f, 1},
			{0x00016e40, 0x00016e7f, 1},
			{0x00016f00, 0x00016f4a, 1},
			{0x00016f4f, 0x00016f87, 1},
			{0x00016f8f, 0x00016f9f, 1},
			{0x00016fe0, 0x00016fe0, 1},
			{0x00016fe1, 0x00016fe1, 1},
			{0x00016fe3, 0x00016fe3, 1},
			{0x00017000, 0x000187f7, 1},
			{0x00018800, 0x00018af2, 1},
			{0x0001b000, 0x0001b11e, 1},
			{0x0001b150, 0x0001b152, 1},
			{0x0001b164, 0x0001b167, 1},
			{0x0001b170, 0x0001b2fb, 1},
			{0x0001bc00, 0x0001bc6a, 1},
			{0x0001bc70, 0x0001bc7c, 1},
			{0x0001bc80, 0x0001bc88, 1},
			{0x0001bc90, 0x0001bc99, 1},
			{0x0001bc9d, 0x0001bc9d, 1},
			{0x0001bc9e, 0x0001bc9e, 1},
			{0x0001d165, 0x0001d169, 1},
			{0x0001d16d, 0x0001d172, 1},
			{0x0001d17b, 0x0001d182, 1},
			{0x0001d185, 0x0001d18b, 1},
			{0x0001d1aa, 0x0001d1ad, 1},
			{0x0001d242, 0x0001d244, 1},
			{0x0001d400, 0x0001d454, 1},
			{0x0001d456, 0x0001d49c, 1},
			{0x0001d49e, 0x0001d49e, 1},
			{0x0001d49f, 0x0001d49f, 1},
			{0x0001d4a2, 0x0001d4a2, 1},
			{0x0001d4a5, 0x0001d4a5, 1},
			{0x0001d4a6, 0x0001d4a6, 1},
			{0x0001d4a9, 0x0001d4ac, 1},
			{0x0001d4ae, 0x0001d4b9, 1},
			{0x0001d4bb, 0x0001d4bb, 1},
			{0x0001d4bd, 0x0001d4c3, 1},
			{0x0001d4c5, 0x0001d505, 1},
			{0x0001d507, 0x0001d50a, 1},
			{0x0001d50d, 0x0001d514, 1},
			{0x0001d516, 0x0001d51c, 1},
			{0x0001d51e, 0x0001d539, 1},
			{0x0001d53b, 0x0001d53e, 1},
			{0x0001d540, 0x0001d544, 1},
			{0x0001d546, 0x0001d546, 1},
			{0x0001d54a, 0x0001d550, 1},
			{0x0001d552, 0x0001d6a5, 1},
			{0x0001d6a8, 0x0001d6c0, 1},
			{0x0001d6c2, 0x0001d6da, 1},
			{0x0001d6dc, 0x0001d6fa, 1},
			{0x0001d6fc, 0x0001d714, 1},
			{0x0001d716, 0x0001d734, 1},
			{0x0001d736, 0x0001d74e, 1},
			{0x0001d750, 0x0001d76e, 1},
			{0x0001d770, 0x0001d788, 1},
			{0x0001d78a, 0x0001d7a8, 1},
			{0x0001d7aa, 0x0001d7c2, 1},
			{0x0001d7c4, 0x0001d7cb, 1},
			{0x0001d7ce, 0x0001d7ff, 1},
			{0x0001da00, 0x0001da36, 1},
			{0x0001da3b, 0x0001da6c, 1},
			{0x0001da75, 0x0001da75, 1},
			{0x0001da84, 0x0001da84, 1},
			{0x0001da9b, 0x0001da9f, 1},
			{0x0001daa1, 0x0001daaf, 1},
			{0x0001e000, 0x0001e006, 1},
			{0x0001e008, 0x0001e018, 1},
			{0x0001e01b, 0x0001e021, 1},
			{0x0001e023, 0x0001e023, 1},
			{0x0001e024, 0x0001e024, 1},
			{0x0001e026, 0x0001e02a, 1},
			{0x0001e100, 0x0001e12c, 1},
			{0x0001e130, 0x0001e13d, 1},
			{0x0001e140, 0x0001e149, 1},
			{0x0001e14e, 0x0001e14e, 1},
			{0x0001e2c0, 0x0001e2f9, 1},
			{0x0001e800, 0x0001e8c4, 1},
			{0x0001e8d0, 0x0001e8d6, 1},
			{0x0001e900, 0x0001e94b, 1},
			{0x0001e950, 0x0001e959, 1},
			{0x0001ee00, 0x0001ee03, 1},
			{0x0001ee05, 0x0001ee1f, 1},
			{0x0001ee21, 0x0001ee21, 1},
			{0x0001ee22, 0x0001ee22, 1},
			{0x0001ee24, 0x0001ee24, 1},
			{0x0001ee27, 0x0001ee27, 1},
			{0x0001ee29, 0x0001ee32, 1},
			{0x0001ee34, 0x0001ee37, 1},
			{0x0001ee39, 0x0001ee39, 1},
			{0x0001ee3b, 0x0001ee3b, 1},
			{0x0001ee42, 0x0001ee42, 1},
			{0x0001ee47, 0x0001ee47, 1},
			{0x0001ee49, 0x0001ee49, 1},
			{0x0001ee4b, 0x0001ee4b, 1},
			{0x0001ee4d, 0x0001ee4f, 1},
			{0x0001ee51, 0x0001ee51, 1},
			{0x0001ee52, 0x0001ee52, 1},
			{0x0001ee54, 0x0001ee54, 1},
			{0x0001ee57, 0x0001ee57, 1},
			{0x0001ee59, 0x0001ee59, 1},
			{0x0001ee5b, 0x0001ee5b, 1},
			{0x0001ee5d, 0x0001ee5d, 1},
			{0x0001ee5f, 0x0001ee5f, 1},
			{0x0001ee61, 0x0001ee61, 1},
			{0x0001ee62, 0x0001ee62, 1},
			{0x0001ee64, 0x0001ee64, 1},
			{0x0001ee67, 0x0001ee6a, 1},
			{0x0001ee6c, 0x0001ee72, 1},
			{0x0001ee74, 0x0001ee77, 1},
			{0x0001ee79, 0x0001ee7c, 1},
			{0x0001ee7e, 0x0001ee7e, 1},
			{0x0001ee80, 0x0001ee89, 1},
			{0x0001ee8b, 0x0001ee9b, 1},
			{0x0001eea1, 0x0001eea3, 1},
			{0x0001eea5, 0x0001eea9, 1},
			{0x0001eeab, 0x0001eebb, 1},
			{0x00020000, 0x0002a6d6, 1},
			{0x0002a700, 0x0002b734, 1},
			{0x0002b740, 0x0002b81d, 1},
			{0x0002b820, 0x0002cea1, 1},
			{0x0002ceb0, 0x0002ebe0, 1},
			{0x0002f800, 0x0002fa1d, 1},
			{0x000e0100, 0x000e01ef, 1},
		},
	}
)
