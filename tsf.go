package tsf

// Constants from the Windows Text Services Framework
const (
	// ENoInterface is returned from QueryInterface when no interface impl is
	// available for the requested IID.
	ENoInterface = 0x80004002

	// EPointer is returned from methods when an invalid or NULL pointer is
	// passed into a required field.
	EPointer = 0x80004003

	// TFMenuReady is not currently used.
	TFMenuReady = 0x00000001

	// TFPropUIStatusSaveToFile specifies that a property is serializable.
	TFPropUIStatusSaveToFile = 0x00000001

	TFTMAENoActivateTIP            = 0x00000001
	TFTMAESecureMode               = 0x00000002
	TFTMAEUIElementEnabledOnly     = 0x00000004
	TFTMAECOMLess                  = 0x00000008
	TFTMAEWOW16                    = 0x00000010
	TFTMAENoActivateKeyboardLayout = 0x00000020
	TFTMAEConsole                  = 0x00000040
)

// GUIDs for interfaces
var (
	CLSIDTFThreadMgr = GUID{
		Data1: 0x529A9E6B,
		Data2: 0x6587,
		Data3: 0x4F23,
		Data4: [8]byte{0xAB, 0x9E, 0x9C, 0x7D, 0x68, 0x3E, 0x3C, 0x50},
	}

	IIDIUnknown = GUID{
		Data1: 0x00000000,
		Data2: 0x0000,
		Data3: 0x0000,
		Data4: [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46},
	}

	IIDTFThreadMgr = GUID{
		Data1: 0xAA80E801,
		Data2: 0x2021,
		Data3: 0x11D2,
		Data4: [8]byte{0x93, 0xE0, 0x00, 0x60, 0xB0, 0x67, 0xB8, 0x6E},
	}

	IIDTFThreadMgrEx = GUID{
		Data1: 0x3e90ade3,
		Data2: 0x7594,
		Data3: 0x4cb0,
		Data4: [8]byte{0xbb, 0x58, 0x69, 0x62, 0x8f, 0x5f, 0x45, 0x8c},
	}

	IIDTFUIElementSink = GUID{
		Data1: 0xea1ea136,
		Data2: 0x19df,
		Data3: 0x11d7,
		Data4: [8]byte{0xa6, 0xd2, 0x00, 0x06, 0x5b, 0x84, 0x43, 0x5c},
	}

	IIDTFSource = GUID{
		Data1: 0x4EA48A35,
		Data2: 0x60AE,
		Data3: 0x446F,
		Data4: [8]byte{0x8F, 0xD6, 0xE6, 0xA8, 0xD8, 0x24, 0x59, 0xF7},
	}
)

// Font attribute GUIDs.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629014(v=vs.85).aspx
var (
	// TSAttrIDFont is not used.
	TSAttrIDFont = defineGUID(0x573ea825, 0x749b, 0x4f8a, 0x9c, 0xfd, 0x21, 0xc3, 0x60, 0x5c, 0xa8, 0x28)

	// TSAttrIDFontFaceName specifies the font face of the font.
	TSAttrIDFontFaceName = defineGUID(0xb536aeb6, 0x053b, 0x4eb8, 0xb6, 0x5a, 0x50, 0xda, 0x1e, 0x81, 0xe7, 0x2e)

	// TSAttrIDFontSizePts specifies the point size of the font.
	TSAttrIDFontSizePts = defineGUID(0xc8493302, 0xa5e9, 0x456d, 0xaf, 0x04, 0x80, 0x05, 0xe4, 0x13, 0x0f, 0x03)

	// TSAttrIDFontStyle is not used.
	TSAttrIDFontStyle = defineGUID(0x68b2a77f, 0x6b0e, 0x4f28, 0x81, 0x77, 0x57, 0x1c, 0x2f, 0x3a, 0x42, 0xb1)

	// TSAttrIDFontStyleAnimation is not used.
	TSAttrIDFontStyleAnimation = defineGUID(0xdcf73d22, 0xe029, 0x47b7, 0xbb, 0x36, 0xf2, 0x63, 0xa3, 0xd0, 0x04, 0xcc)

	// TSAttrIDFontStyleAnimationBlinkingBackground contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationBlinkingBackground = defineGUID(0x86e5b104, 0x0104, 0x4b10, 0xb5, 0x85, 0x00, 0xf2, 0x52, 0x75, 0x22, 0xb5)

	// TSAttrIDFontStyleAnimationLasVegasLights contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationLasVegasLights = defineGUID(0xf40423d5, 0x0f87, 0x4f8f, 0xba, 0xda, 0xe6, 0xd6, 0x0c, 0x25, 0xe1, 0x52)

	// TSAttrIDFontStyleAnimationMarchingBlackAnts contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationMarchingBlackAnts = defineGUID(0x7644e067, 0xf186, 0x4902, 0xbf, 0xc6, 0xec, 0x81, 0x5a, 0xa2, 0x0e, 0x9d)

	// TSAttrIDFontStyleAnimationMarchingRedAnts contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationMarchingRedAnts = defineGUID(0x78368dad, 0x50fb, 0x4c6f, 0x84, 0x0b, 0xd4, 0x86, 0xbb, 0x6c, 0xf7, 0x81)

	// TSAttrIDFontStyleAnimationShimmer contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationShimmer = defineGUID(0x2ce31b58, 0x5293, 0x4c36, 0x88, 0x09, 0xbf, 0x8b, 0xb5, 0x1a, 0x27, 0xb3)

	// TSAttrIDFontStyleAnimationSparkleText contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationSparkleText = defineGUID(0x533aad20, 0x962c, 0x4e9f, 0x8c, 0x09, 0xb4, 0x2e, 0xa4, 0x74, 0x97, 0x11)

	// TSAttrIDFontStyleAnimationWipeDown contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationWipeDown = defineGUID(0x5872e874, 0x367b, 0x4803, 0xb1, 0x60, 0xc9, 0x0f, 0xf6, 0x25, 0x69, 0xd0)

	// TSAttrIDFontStyleAnimationWipeRight contains a nonzero value if the text has the specified animation.
	TSAttrIDFontStyleAnimationWipeRight = defineGUID(0xb855cbe3, 0x3d2c, 0x4600, 0xb1, 0xe9, 0xe1, 0xc9, 0xce, 0x02, 0xf8, 0x42)

	// TSAttrIDFontStyleBackgroundColor specifies the RGB value of the text background. Contains 0xFF000000 if the text color is automatic.
	TSAttrIDFontStyleBackgroundColor = defineGUID(0xb50eaa4e, 0x3091, 0x4468, 0x81, 0xdb, 0xd7, 0x9e, 0xa1, 0x90, 0xc7, 0xc7)

	// TSAttrIDFontStyleBlink contains a nonzero value if the text is blinking or zero otherwise.
	TSAttrIDFontStyleBlink = defineGUID(0xbfb2c036, 0x7acf, 0x4532, 0xb7, 0x20, 0xb4, 0x16, 0xdd, 0x77, 0x65, 0xa8)

	// TSAttrIDFontStyleBold contains a nonzero value if the text is bold or zero otherwise.
	TSAttrIDFontStyleBold = defineGUID(0x48813a43, 0x8a20, 0x4940, 0x8e, 0x58, 0x97, 0x82, 0x3f, 0x7b, 0x26, 0x8a)

	// TSAttrIDFontStyleCapitalize contains a nonzero value if the text is capitalized or zero otherwise.
	TSAttrIDFontStyleCapitalize = defineGUID(0x7d85a3ba, 0xb4fd, 0x43b3, 0xbe, 0xfc, 0x6b, 0x98, 0x5c, 0x84, 0x31, 0x41)

	// TSAttrIDFontStyleColor contains the RGB value of the text color. Contains 0xFF000000 if the text color is automatic.
	TSAttrIDFontStyleColor = defineGUID(0x857a7a37, 0xb8af, 0x4e9a, 0x81, 0xb4, 0xac, 0xf7, 0x00, 0xc8, 0x41, 0x1b)

	// TSAttrIDFontStyleEmboss contains a nonzero value if the text is embossed or zero otherwise.
	TSAttrIDFontStyleEmboss = defineGUID(0xbd8ed742, 0x349e, 0x4e37, 0x82, 0xfb, 0x43, 0x79, 0x79, 0xcb, 0x53, 0xa7)

	// TSAttrIDFontStyleEngrave contains a nonzero value if the text is engraved or zero otherwise.
	TSAttrIDFontStyleEngrave = defineGUID(0x9c3371de, 0x8332, 0x4897, 0xbe, 0x5d, 0x89, 0x23, 0x32, 0x23, 0x17, 0x9a)

	// TSAttrIDFontStyleHeight containss the font height. For more information, see the Height member of the LogFont structure.
	TSAttrIDFontStyleHeight = defineGUID(0x7e937477, 0x12e6, 0x458b, 0x92, 0x6a, 0x1f, 0xa4, 0x4e, 0xe8, 0xf3, 0x91)

	// TSAttrIDFontStyleHidden contains a nonzero value if the text is hidden or zero otherwise.
	TSAttrIDFontStyleHidden = defineGUID(0xb1e28770, 0x881c, 0x475f, 0x86, 0x3f, 0x88, 0x7a, 0x64, 0x7b, 0x10, 0x90)

	// TSAttrIDFontStyleItalic contains a nonzero value if the text is italic or zero otherwise.
	TSAttrIDFontStyleItalic = defineGUID(0x8740682a, 0xa765, 0x48e1, 0xac, 0xfc, 0xd2, 0x22, 0x22, 0xb2, 0xf8, 0x10)

	// TSAttrIDFontStyleKerning contains the minimum kerning size. For more information, see TextFont::GetKerning.
	TSAttrIDFontStyleKerning = defineGUID(0xcc26e1b4, 0x2f9a, 0x47c8, 0x8b, 0xff, 0xbf, 0x1e, 0xb7, 0xcc, 0xe0, 0xdd)

	// TSAttrIDFontStyleLowercase contains a nonzero value if the text is lowercase or zero otherwise.
	TSAttrIDFontStyleLowercase = defineGUID(0x76d8ccb5, 0xca7b, 0x4498, 0x8e, 0xe9, 0xd5, 0xc4, 0xf6, 0xf7, 0x4c, 0x60)

	// TSAttrIDFontStyleOutlined contains a nonzero value if the text is outlined or zero otherwise.
	TSAttrIDFontStyleOutlined = defineGUID(0x10e6db31, 0xdb0d, 0x4ac6, 0xa7, 0xf5, 0x9c, 0x9c, 0xff, 0x6f, 0x2a, 0xb4)

	// TSAttrIDFontStyleOverline contains a nonzero value if the text is overlined or zero otherwise.
	TSAttrIDFontStyleOverline = defineGUID(0xe3989f4a, 0x992b, 0x4301, 0x8c, 0xe1, 0xa5, 0xb7, 0xc6, 0xd1, 0xf3, 0xc8)

	// TSAttrIDFontStyleOverlineDouble contains a nonzero value if the text is double overlined or zero otherwise.
	TSAttrIDFontStyleOverlineDouble = defineGUID(0xdc46063a, 0xe115, 0x46e3, 0xbc, 0xd8, 0xca, 0x67, 0x72, 0xaa, 0x95, 0xb4)

	// TSAttrIDFontStyleOverlineSingle contains a nonzero value if the text is single overlined or zero otherwise.
	TSAttrIDFontStyleOverlineSingle = defineGUID(0x8440d94c, 0x51ce, 0x47b2, 0x8d, 0x4c, 0x15, 0x75, 0x1e, 0x5f, 0x72, 0x1b)

	// TSAttrIDFontStylePosition contains the font position.
	TSAttrIDFontStylePosition = defineGUID(0x15cd26ab, 0xf2fb, 0x4062, 0xb5, 0xa6, 0x9a, 0x49, 0xe1, 0xa5, 0xcc, 0x0b)

	// TSAttrIDFontStyleProtected contains a nonzero value if the text is protected or zero otherwise.
	TSAttrIDFontStyleProtected = defineGUID(0x1c557cb2, 0x14cf, 0x4554, 0xa5, 0x74, 0xec, 0xb2, 0xf7, 0xe7, 0xef, 0xd4)

	// TSAttrIDFontStyleShadow contains a nonzero value if the text is shadowed or zero otherwise.
	TSAttrIDFontStyleShadow = defineGUID(0x5f686d2f, 0xc6cd, 0x4c56, 0x8a, 0x1a, 0x99, 0x4a, 0x4b, 0x97, 0x66, 0xbe)

	// TSAttrIDFontStyleSmallCaps contains a nonzero value if the text is lowercase or zero otherwise.
	TSAttrIDFontStyleSmallCaps = defineGUID(0xfacb6bc6, 0x9100, 0x4cc6, 0xb9, 0x69, 0x11, 0xee, 0xa4, 0x5a, 0x86, 0xb4)

	// TSAttrIDFontStyleSpacing contains the font spacing.
	TSAttrIDFontStyleSpacing = defineGUID(0x98c1200d, 0x8f06, 0x409a, 0x8e, 0x49, 0x6a, 0x55, 0x4b, 0xf7, 0xc1, 0x53)

	// TSAttrIDFontStyleStrikethrough is not used.
	TSAttrIDFontStyleStrikethrough = defineGUID(0x0c562193, 0x2d08, 0x4668, 0x96, 0x01, 0xce, 0xd4, 0x13, 0x09, 0xd7, 0xaf)

	// TSAttrIDFontStyleStrikethroughDouble contains a nonzero value if the text is double strikethrough or zero otherwise.
	TSAttrIDFontStyleStrikethroughDouble = defineGUID(0x62489b31, 0xa3e7, 0x4f94, 0xac, 0x43, 0xeb, 0xaf, 0x8f, 0xcc, 0x7a, 0x9f)

	// TSAttrIDFontStyleStrikethroughSingle contains a nonzero value if the text is single strikethrough or zero otherwise.
	TSAttrIDFontStyleStrikethroughSingle = defineGUID(0x75d736b6, 0x3c8f, 0x4b97, 0xab, 0x78, 0x18, 0x77, 0xcb, 0x99, 0x0d, 0x31)

	// TSAttrIDFontStyleSubscript contains a nonzero value if the text is subscript or zero otherwise.
	TSAttrIDFontStyleSubscript = defineGUID(0x5774fb84, 0x389b, 0x43bc, 0xa7, 0x4b, 0x15, 0x68, 0x34, 0x7c, 0xf0, 0xf4)

	// TSAttrIDFontStyleSuperscript contains a nonzero value if the text is superscript or zero otherwise.
	TSAttrIDFontStyleSuperscript = defineGUID(0x2ea4993c, 0x563c, 0x49aa, 0x93, 0x72, 0x0b, 0xef, 0x09, 0xa9, 0x25, 0x5b)

	// TSAttrIDFontStyleUnderline contains a nonzero value if the text is underlined or zero otherwise.
	TSAttrIDFontStyleUnderline = defineGUID(0xc3c9c9f3, 0x7902, 0x444b, 0x9a, 0x7b, 0x48, 0xe7, 0x0f, 0x4b, 0x50, 0xf7)

	// TSAttrIDFontStyleUnderlineDouble contains a nonzero value if the text is double underlined or zero otherwise.
	TSAttrIDFontStyleUnderlineDouble = defineGUID(0x74d24aa6, 0x1db3, 0x4c69, 0xa1, 0x76, 0x31, 0x12, 0x0e, 0x75, 0x86, 0xd5)

	// TSAttrIDFontStyleUnderlineSingle contains a nonzero value if the text is single underlined or zero otherwise.
	TSAttrIDFontStyleUnderlineSingle = defineGUID(0x1b6720e5, 0x0f73, 0x4951, 0xa6, 0xb3, 0x6f, 0x19, 0xe4, 0x3c, 0x94, 0x61)

	// TSAttrIDFontStyleUppercase contains a nonzero value if the text is uppercase or zero otherwise.
	TSAttrIDFontStyleUppercase = defineGUID(0x33a300e8, 0xe340, 0x4937, 0xb6, 0x97, 0x8f, 0x23, 0x40, 0x45, 0xcd, 0x9a)

	// TSAttrIDFontStyleWeight contains the font weight. For more information about possible values, see the Weight member of the LogFont structure.
	TSAttrIDFontStyleWeight = defineGUID(0x12f3189c, 0x8bb0, 0x461b, 0xb1, 0xfa, 0xea, 0xf9, 0x07, 0x04, 0x7f, 0xe0)
)

// List attribute GUIDs.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629016(v=vs.85).aspx
var (
	// TSAttrIDList is not used.
	TSAttrIDList = defineGUID(0x436d673b, 0x26f1, 0x4aee, 0x9e, 0x65, 0x8f, 0x83, 0xa4, 0xed, 0x48, 0x84)

	// TSAttrIDListLevelIndel contains the index level of the list. 1 is the outermost level, 2 is the next level, and so on.
	TSAttrIDListLevelIndel = defineGUID(0x7f7cc899, 0x311f, 0x487b, 0xad, 0x5d, 0xe2, 0xa4, 0x59, 0xe1, 0x2d, 0x42)

	// TSAttrIDListType is not used.
	TSAttrIDListType = defineGUID(0xae3e665e, 0x4bce, 0x49e3, 0xa0, 0xfe, 0x2d, 0xb4, 0x7d, 0x3a, 0x17, 0xae)

	// TSAttrIDListTypeArabic contains a nonzero value if the list is an arabic numeral list or zero otherwise.
	TSAttrIDListTypeArabic = defineGUID(0xbccd77c5, 0x4c4d, 0x4ce2, 0xb1, 0x02, 0x55, 0x9f, 0x3b, 0x2b, 0xfc, 0xea)

	// TSAttrIDListTypeBullet contains a nonzero value if the list is a bulleted list or zero otherwise.
	TSAttrIDListTypeBullet = defineGUID(0x1338c5d6, 0x98a3, 0x4fa3, 0x9b, 0xd1, 0x7a, 0x60, 0xee, 0xf8, 0xe9, 0xe0)

	// TSAttrIDListTypeLowerLetter contains a nonzero value if the list is a lowercase lettered list or zero otherwise.
	TSAttrIDListTypeLowerLetter = defineGUID(0x96372285, 0xf3cf, 0x491e, 0xa9, 0x25, 0x38, 0x32, 0x34, 0x7f, 0xd2, 0x37)

	// TSAttrIDListTypeLowerRoman contains a nonzero value if the list is a lowercase roman numeral list or zero otherwise.
	TSAttrIDListTypeLowerRoman = defineGUID(0x90466262, 0x3980, 0x4b8e, 0x93, 0x68, 0x91, 0x8b, 0xd1, 0x21, 0x8a, 0x41)

	// TSAttrIDListTypeUpperLetter contains a nonzero value if the list is an upper-case lettered list or zero otherwise.
	TSAttrIDListTypeUpperLetter = defineGUID(0x7987b7cd, 0xce52, 0x428b, 0x9b, 0x95, 0xa3, 0x57, 0xf6, 0xf1, 0x0c, 0x45)

	// TSAttrIDListTypeUpperRoman contains a nonzero value if the list is an uppercase roman numeral list or zero otherwise.
	TSAttrIDListTypeUpperRoman = defineGUID(0x0f6ab552, 0x4a80, 0x467f, 0xb2, 0xf1, 0x12, 0x7e, 0x2a, 0xa3, 0xba, 0x9e)
)

// Text attribute GUIDs.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629018(v=vs.85).aspx
var (
	// TSAttrIDText is not used.
	TSAttrIDText = defineGUID(0x7edb8e68, 0x81f9, 0x449d, 0xa1, 0x5a, 0x87, 0xa8, 0x38, 0x8f, 0xaa, 0xc0)

	// TSAttrIDTextAlignment is not used.
	TSAttrIDTextAlignment = defineGUID(0x139941e6, 0x1767, 0x456d, 0x93, 0x8e, 0x35, 0xba, 0x56, 0x8b, 0x5c, 0xd4)

	// TSAttrIDTextAlignmentCenter contains a nonzero value if the text is centered or zero otherwise.
	TSAttrIDTextAlignmentCenter = defineGUID(0xa4a95c16, 0x53bf, 0x4d55, 0x8b, 0x87, 0x4b, 0xdd, 0x8d, 0x42, 0x75, 0xfc)

	// TSAttrIDTextAlignmentJustify contains a nonzero value if the text is justified or zero otherwise.
	TSAttrIDTextAlignmentJustify = defineGUID(0xed350740, 0xa0f7, 0x42d3, 0x8e, 0xa8, 0xf8, 0x1b, 0x64, 0x88, 0xfa, 0xf0)

	// TSAttrIDTextAlignmentLeft contains a nonzero value if the text is left aligned or zero otherwise.
	TSAttrIDTextAlignmentLeft = defineGUID(0x16ae95d3, 0x6361, 0x43a2, 0x84, 0x95, 0xd0, 0x0f, 0x39, 0x7f, 0x16, 0x93)

	// TSAttrIDTextAlignmentRight contains a nonzero value if the text is right aligned or zero otherwise.
	TSAttrIDTextAlignmentRight = defineGUID(0xb36f0f98, 0x1b9e, 0x4360, 0x86, 0x16, 0x03, 0xfb, 0x08, 0xa7, 0x84, 0x56)

	// TSAttrIDTextEmbeddedObject contains a nonzero value if the text is an embedded object or zero otherwise.
	TSAttrIDTextEmbeddedObject = defineGUID(0x7edb8e68, 0x81f9, 0x449d, 0xa1, 0x5a, 0x87, 0xa8, 0x38, 0x8f, 0xaa, 0xc0)

	// TSAttrIDTextHyphenation contains a nonzero value if the text is hyphenated or zero otherwise.
	TSAttrIDTextHyphenation = defineGUID(0xdadf4525, 0x618e, 0x49eb, 0xb1, 0xa8, 0x3b, 0x68, 0xbd, 0x76, 0x48, 0xe3)

	// TSAttrIDTextLanguage contains the LANGID language identifier of the text.
	TSAttrIDTextLanguage = defineGUID(0xd8c04ef1, 0x5753, 0x4c25, 0x88, 0x87, 0x85, 0x44, 0x3f, 0xe5, 0xf8, 0x19)

	// TSAttrIDTextLink contains a pointer to a link object. The caller must use the QueryInterface method to obtain the desired interface, such as IUniformResourceLocator.
	TSAttrIDTextLink = defineGUID(0x47cd9051, 0x3722, 0x4cd8, 0xb7, 0xc8, 0x4e, 0x17, 0xca, 0x17, 0x59, 0xf5)

	// TSAttrIDTextOrientation specifies the angle, in tenths of degrees, between text base line and the x-axis of the device.
	TSAttrIDTextOrientation = defineGUID(0x6bab707f, 0x8785, 0x4c39, 0x8b, 0x52, 0x96, 0xf8, 0x78, 0x30, 0x3f, 0xfb)

	// TSAttrIDTextPara is not used.
	TSAttrIDTextPara = defineGUID(0x5edc5822, 0x99dc, 0x4dd6, 0xae, 0xc3, 0xb6, 0x2b, 0xaa, 0x5b, 0x2e, 0x7c)

	// TSAttrIDTextParaFirstLineIndent contains the number of points that the first line of a paragraph is indented.
	TSAttrIDTextParaFirstLineIndent = defineGUID(0x07c97a13, 0x7472, 0x4dd8, 0x90, 0xa9, 0x91, 0xe3, 0xd7, 0xe4, 0xf2, 0x9c)

	// TSAttrIDTextParaLeftIndent contains the number of points that the paragraph is indented from the left.
	TSAttrIDTextParaLeftIndent = defineGUID(0xfb2848e9, 0x7471, 0x41c9, 0xb6, 0xb3, 0x8a, 0x14, 0x50, 0xe0, 0x18, 0x97)

	// TSAttrIDTextParaLineSpacing is not used.
	TSAttrIDTextParaLineSpacing = defineGUID(0x699b380d, 0x7f8c, 0x46d6, 0xa7, 0x3b, 0xdf, 0xe3, 0xd1, 0x53, 0x8d, 0xf3)

	// TSAttrIDTextParaLineSpacingAtLeast contains the minimum number of lines for the line spacing of the paragraph.
	TSAttrIDTextParaLineSpacingAtLeast = defineGUID(0xadfedf31, 0x2d44, 0x4434, 0xa5, 0xff, 0x7f, 0x4c, 0x49, 0x90, 0xa9, 0x05)

	// TSAttrIDTextParaLineSpacingDouble contains a nonzero value if the paragraph is double spaced or zero otherwise.
	TSAttrIDTextParaLineSpacingDouble = defineGUID(0x82fb1805, 0xa6c4, 0x4231, 0xac, 0x12, 0x62, 0x60, 0xaf, 0x2a, 0xba, 0x28)

	// TSAttrIDTextParaLineSpacingExactly contains the exact number of lines for the line spacing of the paragraph.
	TSAttrIDTextParaLineSpacingExactly = defineGUID(0x3d45ad40, 0x23de, 0x48d7, 0xa6, 0xb3, 0x76, 0x54, 0x20, 0xc6, 0x20, 0xcc)

	// TSAttrIDTextParaLineSpacingMultiple contains the number of lines for the multiple line spacing of the paragraph.
	TSAttrIDTextParaLineSpacingMultiple = defineGUID(0x910f1e3c, 0xd6d0, 0x4f65, 0x8a, 0x3c, 0x42, 0xb4, 0xb3, 0x18, 0x68, 0xc5)

	// TSAttrIDTextParaLineSpacingOnePtFive contains a nonzero value if the paragraph is one and one half line spaced or zero otherwise.
	TSAttrIDTextParaLineSpacingOnePtFive = defineGUID(0x0428a021, 0x0397, 0x4b57, 0x9a, 0x17, 0x07, 0x95, 0x99, 0x4c, 0xd3, 0xc5)

	// TSAttrIDTextParaLineSpacingSingle contains a nonzero value if the paragraph is single spaced or zero otherwise.
	TSAttrIDTextParaLineSpacingSingle = defineGUID(0xed350740, 0xa0f7, 0x42d3, 0x8e, 0xa8, 0xf8, 0x1b, 0x64, 0x88, 0xfa, 0xf0)

	// TSAttrIDTextParaRightIndent contains the number of points that the paragraph is indented from the right.
	TSAttrIDTextParaRightIndent = defineGUID(0x2c7f26f9, 0xa5e2, 0x48da, 0xb9, 0x8a, 0x52, 0x0c, 0xb1, 0x65, 0x13, 0xbf)

	// TSAttrIDTextParaSpaceAfter contains the number of points of spacing after the paragraph.
	TSAttrIDTextParaSpaceAfter = defineGUID(0x7b0a3f55, 0x22dc, 0x425f, 0xa4, 0x11, 0x93, 0xda, 0x1d, 0x8f, 0x9b, 0xaa)

	// TSAttrIDTextParaSpaceBefore contains the number of points of spacing before the paragraph.
	TSAttrIDTextParaSpaceBefore = defineGUID(0x8df98589, 0x194a, 0x4601, 0xb2, 0x51, 0x98, 0x65, 0xa3, 0xe9, 0x06, 0xdd)

	// TSAttrIDTextReadOnly contains zero if the text is read-only or nonzero otherwise.
	TSAttrIDTextReadOnly = defineGUID(0x85836617, 0xde32, 0x4afd, 0xa5, 0x0f, 0xa2, 0xdb, 0x11, 0x0e, 0x6e, 0x4d)

	// TSAttrIDTextRightToLeft contains zero if the text is right-to-left reading or nonzero otherwise.
	TSAttrIDTextRightToLeft = defineGUID(0xca666e71, 0x1b08, 0x453d, 0xbf, 0xdd, 0x28, 0xe0, 0x8c, 0x8a, 0xaf, 0x7a)

	// TSAttrIDTextVerticalWriting specifies if the text is vertical or horizontal. Contains zero if the text is horizontal or nonzero if the text is vertical.
	TSAttrIDTextVerticalWriting = defineGUID(0x6bba8195, 0x046f, 0x4ea9, 0xb3, 0x11, 0x97, 0xfd, 0x66, 0xc4, 0x27, 0x4b)
)

// App attribute GUIDs.
var (
	// TSAttrIDApp is not used.
	TSAttrIDApp = defineGUID(0xa80f77df, 0x4237, 0x40e5, 0x84, 0x9c, 0xb5, 0xfa, 0x51, 0xc1, 0x3a, 0xc7)

	// TSAttrIDAppIncorrectGrammar contains a nonzero value if the text contains a grammar error or zero otherwise.
	TSAttrIDAppIncorrectGrammar = defineGUID(0xbd54e398, 0xad03, 0x4b74, 0xb6, 0xb3, 0x5e, 0xdb, 0x19, 0x99, 0x63, 0x88)

	// TSAttrIDAppIncorrectSpelling contains a nonzero value if the text contains a spelling error or zero otherwise.
	TSAttrIDAppIncorrectSpelling = defineGUID(0xf42de43c, 0xef12, 0x430d, 0x94, 0x4c, 0x9a, 0x08, 0x97, 0x0a, 0x25, 0xd2)

	// TSAttrIDOthers is not used.
	TSAttrIDOthers = defineGUID(0xb3c32af9, 0x57d0, 0x46a9, 0xbc, 0xa8, 0xda, 0xc2, 0x38, 0xa1, 0x30, 0x57)
)

// Speech recognition constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629029(v=vs.85).aspx
const (
	// TFDictationOn specifies whether or not dictation is active.
	//
	// If this bit is 1, speech dictation is active. Otherwise, speech
	// dictation is inactive. This value cannot be combined with
	// TF_COMMANDING_ON. Used with the
	// GUID_COMPARTMENT_SPEECH_DICTATIONSTAT compartment.
	TFDictationOn = 0x00000001

	// TFDictationEnabled specifies whether or not dictation is enabled.
	//
	// If this bit is 1, speech dictation is enabled. Otherwise, speech
	// dictation is disabled. Used with the
	// GUID_COMPARTMENT_SPEECH_DICTATIONSTAT compartment.
	TFDictationEnabled = 0x00000002

	// TFCommandingEnabled specifies whether or not speech command is enabled.
	//
	// If this bit is 1, speech command is enabled. Otherwise, speech command
	// is disabled. Used with the GUID_COMPARTMENT_SPEECH_DICTATIONSTAT
	// compartment.
	TFCommandingEnabled = 0x00000004

	// TFCommandingOn specifies whether or not speech command is active.
	//
	// If this bit is 1, speech command is active. Otherwise, speech command
	// is inactive. This value cannot be combined with TF_DICTATION_ON. Used
	// with the GUID_COMPARTMENT_SPEECH_DICTATIONSTAT compartment.
	TFCommandingOn = 0x00000008

	// TFSpeechUIShown is not currently used.
	TFSpeechUIShown = 0x00000010

	// TFDictationOn is not currently used. Used with the
	// GUID_COMPARTMENT_SPEECH_UI_STATUS compartment.
	TFShowBalloon = 0x00000001

	// TFDisableBalloon specifies whether or not balloon display is disabled.
	//
	// If this bit is 1, balloon display for the current speech mode is
	// disabled. Otherwise, balloon display for the current speech mode is
	// enabled. Used with the GUID_COMPARTMENT_SPEECH_UI_STATUS compartment.
	TFDisableBalloon = 0x00000002

	// TFDisableSpeech specifies whether or not speech input is disabled.
	//
	// If this bit is 1, speech input is disabled. Otherwise, speech input is
	// enabled. Used with the GUID_COMPARTMENT_SPEECH_DISABLED compartment.
	TFDisableSpeech = 0x00000001

	// TFDisableDictation specifies whether or not dictation is disabled.
	//
	// If this bit is 1, speech dictation is disabled. Otherwise, speech
	// dictation is enabled. Used with the GUID_COMPARTMENT_SPEECH_DISABLED
	// compartment.
	TFDisableDictation = 0x00000002

	// TFDisableCommanding specifies whether or not speech command is
	// disabled.
	//
	// If this bit is 1, speech command is disabled. Otherwise, speech
	// command is enabled. Used with the GUID_COMPARTMENT_SPEECH_DISABLED
	TFDisableCommanding = 0x00000004
)

// Constants used for input conversion mode.
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa380396(v=vs.85).aspx
const (
	// TFConversionModeAlphaNumeric is set to 1 if ALPHANUMERIC mode.
	TFConversionModeAlphaNumeric = 0x0000

	// TFConversionModeNative is set to 1 if NATIVE mode; 0 if ALPHANUMERIC mode.
	TFConversionModeNative = 0x0001

	// TFConversionModeKatakana is set to 1 if KATAKANA mode; 0 if HIRAGANA mode.
	TFConversionModeKatakana = 0x0002

	// TFConversionModeFullShape is set to 1 if full shape mode; 0 if half shape mode.
	TFConversionModeFullShape = 0x0008

	// TFConversionModeRoman is set to 1 to prevent processing of conversions by IME; 0 if not.
	TFConversionModeRoman = 0x0010

	// TFConversionModeCharInput is set to 1 if character code input mode; 0 if not.
	TFConversionModeCharInput = 0x0020

	// TFConversionModeSoftKeyboard is set to 1 if Soft Keyboard mode; 0 if not.
	TFConversionModeSoftKeyboard = 0x0080

	// TFConversionModeNoConversion is set to 1 to prevent processing of conversions by IME; 0 if not.
	TFConversionModeNoConversion = 0x0100

	// TFConversionModeSymbol is set to 1 if SYMBOL conversion mode; 0 if not.
	TFConversionModeSymbol = 0x0400

	// TFConversionModeFixed is set to 1 if fixed conversion mode; 0 if not.
	TFConversionModeFixed = 0x0800
)

// Sentence mode values.
const (
	// TFSentenceModeNone specifies that there is no information for sentence.
	TFSentenceModeNone = 0x0000

	// TFSentenceModePlauralClause specifies that the IME uses plural clause
	// information to carry out conversion processing.
	TFSentenceModePlauralClause = 0x0001

	// TFSentenceModeSingleConvert specifies that the IME carries out
	// conversion processing in single-character mode.
	TFSentenceModeSingleConvert = 0x0002

	// TFSentenceModeAutomatic specifies that the IME carries out conversion
	// processing in automatic mode.
	TFSentenceModeAutomatic = 0x0004

	// TFSentenceModePhrasePredict specifies that the IME uses phrase
	// information to predict the next character.
	TFSentenceModePhrasePredict = 0x0008

	// TFSentenceModeConversation specifies that the IME uses conversation
	// mode. This is useful for chat applications.
	TFSentenceModeConversation = 0x0010
)

// Language bar constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629003(v=vs.85).aspx
const (
	// TFDTLBIUseProfileIcon specifies that the system language bar item
	// should display the icon specified for the language profile.
	TFDTLBIUseProfileIcon = 0x00000001

	// TFInvalidMenuItem is not used.
	TFInvalidMenuItem = 0xffffffff

	// TFLBIBMPFVertical is not used.
	TFLBIBMPFVertical = 0x00000001

	// TFLBIDescMaxLen is the length of TFLangBarItemInfo.Description.
	TFLBIDescMaxLen = 32
)

// Langauge bar LBI constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629079(v=vs.85).aspx
const (
	// TFLBIIcon specifies that the icon of the item has changed. The language
	// bar calls ITfLangBarItemBitmapButton::GetIcon in response to this
	// notification.
	TFLBIIcon = 0x00000001

	// TFLBIText specifies that the text of a button or bitmap button item has
	// changed. The language bar calls ITfLangBarItemButton::GetText or
	// ITfLangBarItemBitmapButton::GetText, whichever is appropriate, in
	// response to this notification.
	TFLBIText = 0x00000002

	// TFLBITooltip specifies that the tooltip text of the item changed. The
	// language bar calls ITfLangBarItem::GetTooltipString in response to this
	// notification.
	TFLBITooltip = 0x00000004

	// TFLBIBitmap specifies that the bitmap of a bitmap or bitmap button
	// item changed. The language bar calls ITfLangBarItemBitmap::DrawBitmap
	// or ITfLangBarItemBitmapButton::DrawBitmap, whichever is appropriate, in
	// response to this notification.
	TFLBIBitmap = 0x00000008

	// TFLBIBalloon specifies that the information for a balloon item changed.
	// The language bar calls ITfLangBarItemBalloon::GetBalloonInfo in
	// response to this notification.
	TFLBIBalloon = 0x00000010

	// TFLBIStatus specifies that the item status changed. The language bar
	// calls ITfLangBarItem::GetStatus in response to this notification.
	TFLBIStatus = 0x00010000

	TFLBIBmpAll    = TFLBIBitmap | TFLBITooltip
	TFLBIBmpBtnAll = TFLBIBitmap | TFLBIText | TFLBITooltip
	TFLBIBtnAll    = TFLBIIcon | TFLBIText | TFLBITooltip
)

// Langauge bar LBIStatus constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629077(v=vs.85).aspx
const (
	// TFLBIStatusHidden specifies that the item is hidden. This style is
	// ignored if the item does not include the TFLBIStyleHiddenStatusControl
	// style.
	TFLBIStatusHidden = 0x00000001

	// TFLBIStatusDisabled specifies that the item is disabled.
	TFLBIStatusDisabled = 0x00000002

	// TFLBIStatusBtnToggle specifies that the item is in the toggled or
	// pressed state. This style is ignored if the item does not include
	// the TFLBIStyleBtnToggle style.
	TFLBIStatusBtnToggle = 0x00010000
)

// Language bar LBIStyle constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629078(v=vs.85).aspx
const (
	// TFLBIStyleHiddenStatusControl specifies that the item can be hidden or
	// shown dynamically using the TF_LBI_STATUS_HIDDEN value in the
	// ITfLangBarItem::GetStatus method. If this value is not present, the item
	// cannot be hidden in this manner.
	TFLBIStyleHiddenStatusControl = 0x00000001

	// TFLBIStyleShowInTray specifies that the item will be displayed in the
	// notification icon tray in addition to the language bar. This flag is not
	// currently supported.
	TFLBIStyleShowInTray = 0x00000002

	// TFLBIStyleHideOnNoOtherItems specifies that the language bar is hidden if
	// all items in the language bar contain this style. If any item in the
	// language bar does not contain this style, the language bar is displayed.
	TFLBIStyleHideOnNoOtherItems = 0x00000004

	// TFLBIStyleShowInTrayOnly specifies that the item will only be displayed in
	// the notification icon tray and not in the language bar. This flag is not
	// currently supported.
	TFLBIStyleShowInTrayOnly = 0x00000008

	// TFLBIStyleHiddenByDefault specifies that the item is not displayed in the
	// toolbar until it is selected from the language bar options menu. This flag
	// is ignored if the TFLBIStyleHiddenStatusControl is set or the user has
	// already changed the hidden/shown state using the language bar options menu.
	TFLBIStyleHiddenByDefault = 0x00000010

	// TFLBIStyleTextColorIcon specifies that tny black pixel within the icon
	// will be converted to the text color of the selected theme. The icon must
	// be monochrome.
	TFLBIStyleTextColorIcon = 0x00000020

	// TFLBIStyleBtnButton specifies that the item is a push button.
	// ITfLangBarItemButton::OnClick is called when the item is pressed.
	TFLBIStyleBtnButton = 0x00010000

	// TFLBIStyleBtnMenu specifies that the item is a menu.
	// ITfLangBarItemButton::InitMenu is called when the item is pressed.
	//
	// If this style is combined with TFLBIStyleBtnButton, a drop-down arrow will
	// be displayed for the item in addition to the text. The drop-down arrow
	// functions as the menu button and clicking it will cause
	// ITfLangBarItemButton::InitMenu to be called. Clicking the text portion of
	// the button will cause ITfLangBarItemButton::OnClick to be called.
	TFLBIStyleBtnMenu = 0x00020000

	// TFLBIStyleBtnToggle specifies that the item is a toggle button and operates
	// similar to a check box. ITfLangBarItemButton::OnClick is called when the
	// item is pressed.
	TFLBIStyleBtnToggle = 0x00040000
)

// Language bar LBMenuF constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629080(v=vs.85).aspx
const (
	// TFLBMenuFChecked specifies that the menu item is checked.
	TFLBMenuFChecked = 0x01

	// TFLBMenuFSubMenu specifies that the menu item is a submenu.
	TFLBMenuFSubMenu = 0x02

	// TFLBMenuFSeparator specifies that the menu item is a separator.
	TFLBMenuFSeparator = 0x04

	// TFLBMenuFRadioChecked specifies that the menu item is a radio check mark.
	TFLBMenuFRadioChecked = 0x08

	// TFLBMenuFGrayed specifies that the menu item is disabled.
	TFLBMenuFGrayed = 0x10
)

// Floating language bar display settings constants.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629190(v=vs.85).aspx
const (
	// TFSFTShowNormal specifies that language bar should be displayed as a
	// floating window. This constant cannot be combined with the TFSFTDock,
	// TFSFTMinimized, TFSFTHidden, or TFSFTDeskBand constants.
	TFSFTShowNormal = 0x00000001

	// TFSFTDock specifies that the language bar should dock in its own task
	// pane. This constant cannot be combined with the TFSFTShowNormal,
	// TFSFTMinimized, TFSFTHidden, or TFSFTDeskBand constants. Available only
	// on Windows XP.
	TFSFTDock = 0x00000002

	// TFSFTMinimized specifies that the language bar should be displayed as a
	// single icon in the system tray. This constant cannot be combined with
	// the TFSFTShowNormal, TFSFTDock, TFSFTHidden, or TFSFTDeskBand constants.
	// In Windows XP, use TFSFTDeskBand instead.
	TFSFTMinimized = 0x00000004

	// TFSFTHidden specifies that the language bar should be hidden. This
	// constant cannot be combined with the TFSFTShowNormal, TFSFTDock,
	// TFSFTMinimized, or TFSFTDeskBand constants.
	TFSFTHidden = 0x00000008

	// TFSFTNoTransparency specifies that language bar should be opaque.
	TFSFTNoTransparency = 0x00000010

	// TFSFTLowTransparency specifies that language bar should be partially
	// transparent. Available only on Windows 2000 or later.
	TFSFTLowTransparency = 0x00000020

	// TFSFTHighTransparency specifies that language bar should be highly
	// transparent. Available only on Windows 2000 or later.
	TFSFTHighTransparency = 0x00000040

	// TFSFTLabels specifies that text labels should display next to language
	// bar icons.
	TFSFTLabels = 0x00000080

	// TFSFTNoLabels specifies that language bar icon text labels should be
	// hidden.
	TFSFTNoLabels = 0x00000100

	// TFSFTExtraIconsOnMinimized specifies that text service icons on the
	// taskbar should be displayed when the language bar is minimized.
	TFSFTExtraIconsOnMinimized = 0x00000200

	// TFSFTNoExtraIconsOnMinimized specifies that text service icons on the
	// taskbar should be hidden when the language bar is minimized.
	TFSFTNoExtraIconsOnMinimized = 0x00000400

	// TFSFTDeskBand specifies that the language bar should dock in the
	// righthand end of the system task bar (immediately left of the system
	// tray/clock). This constant cannot be combined with the TFSFTShowNormal,
	// TFSFTDock, TFSFTMinimized, or TFSFTHidden constants.
	// Available only on Windows XP.
	TFSFTDeskBand = 0x00000800
)

// Types from the Windows Text Services Framework
type (
	// TFClientID is the data type used to identify the client.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629049(v=vs.85).aspx
	TFClientID uint32

	// TFEditCookie identifies an edit session that has a lock.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629050(v=vs.85).aspx
	TFEditCookie uint32

	// TFGuidAtom identifies a GUID within TSF. A TFGuidAtom is obtained by
	// calling TFCategoryMgr::RegisterGUID
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629052(v=vs.85).aspx
	TFGuidAtom uint32

	// TSAttrID is used to identify a text attribute.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629205(v=vs.85).aspx
	TSAttrID GUID

	// TSViewCookie identifies a context view.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629203(v=vs.85).aspx
	TSViewCookie uint32

	// ColorType specifies types of colors that are available for a soft
	// keyboard.
	//
	//
	ColorType int

	// InputScope specifies which input scopes are applied to a given field.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms538181(v=vs.85).aspx
	InputScope int

	// TFDAAttrInfo specfies text conversion data in the TFDisplayAttribute
	// structure
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629063(v=vs.85).aspx
	TFDAAttrInfo int

	// TFDAColorType specify the format of the color contained in the
	// TFDAColor structure.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629065(v=vs.85).aspx
	TFDAColorType int

	// TFDALineStyle specifies the underline style of a display attribute
	// in the TFDAColor structure.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629066(v=vs.85).aspx
	TFDALineStyle int

	// TFActiveSelEnd specifies which end of a selected range of text is active.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629046(v=vs.85).aspx
	TFActiveSelEnd int

	// TFAnchor specifies the start anchor or end anchor of an ITfRange object.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629047(v=vs.85).aspx
	TFAnchor int

	// TFCandidateResult specifies the result of a reconversion operation
	// performed on a given candidate string.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629048(v=vs.85).aspx
	TFCandidateResult int

	// TFGravity specifies the type of gravity associated with the anchor
	// of a TFRange object
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629051(v=vs.85).aspx
	TFGravity int

	// TFIntegratableCandidateListSelectionStyle specifies the integratable
	// candidate list selection styles.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/hh920975(v=vs.85).aspx
	TFIntegratableCandidateListSelectionStyle int

	// TFLayoutCode specifies the type of layout change in an
	// TFTextLayoutSink::OnLayoutChange notification.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629053(v=vs.85).aspx
	TFLayoutCode int

	// TFLBBalloonStyle specifies a language bar balloon style.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629054(v=vs.85).aspx
	TFLBBalloonStyle int

	// TFLBIClick specifies which mouse button was used to click a toolbar item.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629055(v=vs.85).aspx
	TFLBIClick int

	// TFSAPIObject specifies a specific type of Speech API (SAPI) object.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629056(v=vs.85).aspx
	TFSAPIObject int

	// TFShiftDir specifies which direction a range anchor is moved.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629057(v=vs.85).aspx
	TFShiftDir int

	// TitlebarType specifies types of titlebars that are available for a
	// soft keyboard window.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd388707(v=vs.85).aspx
	TitlebarType int

	// TKBLayoutType specifies the type of layout for a soft keyboard.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/hh802866(v=vs.85).aspx
	TKBLayoutType int

	// TSActiveSelEnd specifies which end of a text store selection is active.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629196(v=vs.85).aspx
	TSActiveSelEnd int

	// TSGravity specifies the gravity type associated with an Anchor object.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629199(v=vs.85).aspx
	TSGravity int

	// TSLayoutCode specifies the type of layout change in an
	// TextStoreACPSink::OnLayoutChange or
	// TextStoreAnchorSink::OnLayoutChange notification.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629200(v=vs.85).aspx
	TSLayoutCode int

	// TSRunType specifies if a text run is visible, hidden, or is a private
	// data type embedded in the text run.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629201(v=vs.85).aspx
	TSRunType int

	// TSShiftDir specifies which direction an anchor is moved.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629202(v=vs.85).aspx
	TSShiftDir int

	// TypeMode specifies type modes that are available for a soft keyboard.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd388738(v=vs.85).aspx
	TypeMode int

	// TFDAColor contains color data used in the display attributes for a
	// range of text.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629064(v=vs.85).aspx
	TFDAColor struct {
		Type         TFDAColorType
		ColorOrIndex uint32
	}

	// TFDisplayAttribute contains display attribute data for rendering text.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629067(v=vs.85).aspx
	TFDisplayAttribute struct {
		TextColor      TFDAColor
		BGColor        TFDAColor
		UnderlineStyle TFDALineStyle
		BoldLine       int32
		UnderlineColor TFDAColor
		Attr           TFDAAttrInfo
	}

	// TFHaltCond is used to contain conditions of a range shift.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms629071(v=vs.85).aspx
	TFHaltCond struct {
		HaltRange *ITfRange
		HaltPos   *TFAnchor
		Flags     uint32
	}

	// TFInputProcessorProfile contains data for the input processor profile.
	//
	// https://msdn.microsoft.com/en-us/library/windows/desktop/aa383446(v=vs.85).aspx
	TFInputProcessorProfile struct {
		ProfileType   uint32
		LangID        uint16
		ClsID         GUID
		ProfileGUID   GUID
		CatID         GUID
		HKLSubstitute uintptr
		Caps          uint32
		HKL           uintptr
		Flags         uint32
	}

	// TFLangBarItemInfo is used to hold information about a language bar item.
	TFLangBarItemInfo struct {
		ClsIDService GUID
		ItemGUID     GUID
		Style        uint32
		Sort         uint32
		Description  [TFLBIDescMaxLen]uint16
	}

	// TFLanguageProfile contains information about a language profile.
	TFLanguageProfile struct {
		ClsID       GUID
		LangID      uint16
		CatID       GUID
		Active      uint32
		ProfileGUID GUID
	}

	// TFLattElement contains information about a lattice element.
	TFLattElement struct {
		FrameStart uint32
		FrameLen   uint32
		Flags      uint32
		Cost       int32
		Text       BStr
	}

	// TFPropertyVal contains property value data.
	TFPropertyVal struct {
		ID    GUID
		Value GUID
	}

	// TFRenderingMarkup contains a range and the display attribute
	// information.
	TFRenderingMarkup struct {
		Range       *ITfRange
		DisplayAttr TFDisplayAttribute
	}
)

// Elements of the ColorType enumeration are used to specify types of colors
// that are available for a soft keyboard.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/dd374683(v=vs.85).aspx
const (
	// BackColor specifies background color.
	BackColor ColorType = 0

	// UnSelForeColor specifies unselected foreground color.
	UnSelForeColor = 1

	// UnSelTextColor specifies unselected text color.
	UnSelTextColor = 2

	// SelForeColor specifies selected foreground color.
	SelForeColor = 3

	// SelTextColor specifies selected text color.
	SelTextColor = 4

	// MaxColorType specifies maximum color type.
	MaxColorType = 5
)
