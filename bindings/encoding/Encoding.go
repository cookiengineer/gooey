//go:build wasm

package encoding

type Encoding string

const (

	// Web standard encoding
	EncodingUTF8 Encoding = "utf-8"

	// Legacy multi-byte encodings
	EncodingUTF16BE Encoding = "utf-16be"
	EncodingUTF16LE Encoding = "utf-16le"

	// Legacy single-byte encodings
	EncodingIBM866      Encoding = "ibm866"
	EncodingISO_8859_2  Encoding = "iso-8859-2"
	EncodingISO_8859_3  Encoding = "iso-8859-3"
	EncodingISO_8859_4  Encoding = "iso-8859-4"
	EncodingISO_8859_5  Encoding = "iso-8859-5"
	EncodingISO_8859_6  Encoding = "iso-8859-6"
	EncodingISO_8859_7  Encoding = "iso-8859-7"
	EncodingISO_8859_8  Encoding = "iso-8859-8"
	EncodingISO_8859_8I Encoding = "iso-8859-8i"
	EncodingISO_8859_10 Encoding = "iso-8859-10"
	EncodingISO_8859_13 Encoding = "iso-8859-13"
	EncodingISO_8859_14 Encoding = "iso-8859-14"
	EncodingISO_8859_15 Encoding = "iso-8859-15"
	EncodingISO_8859_16 Encoding = "iso-8859-16"

	// Legacy single-byte encodings
	EncodingKOI8R          Encoding = "koi8-r"
	EncodingKOI8U          Encoding = "koi8-u"
	EncodingMACINTOSH      Encoding = "macintosh"
	EncodingWINDOWS_874    Encoding = "windows-874"
	EncodingWINDOWS_1250   Encoding = "windows-1250"
	EncodingWINDOWS_1251   Encoding = "windows-1251"
	EncodingWINDOWS_1252   Encoding = "windows-1252"
	EncodingWINDOWS_1253   Encoding = "windows-1253"
	EncodingWINDOWS_1254   Encoding = "windows-1254"
	EncodingWINDOWS_1255   Encoding = "windows-1255"
	EncodingWINDOWS_1256   Encoding = "windows-1256"
	EncodingWINDOWS_1257   Encoding = "windows-1257"
	EncodingWINDOWS_1258   Encoding = "windows-1258"
	EncodingX_MAC_CYRILLIC Encoding = "x-mac-cyrillic"

	// Legacy multi-byte encodings
	EncodingBIG5        Encoding = "big5"
	EncodingGBK         Encoding = "gbk"
	EncodingGB18030     Encoding = "gb18030"
	EncodingEUCJP       Encoding = "euc-jp"
	EncodingEUCKR       Encoding = "euc-kr"
	EncodingISO_2022_JP Encoding = "iso-2022-jp"
	EncodingSHIFTJIS    Encoding = "shift-jis"

)

func (encoding Encoding) String() string {
	return string(encoding)
}
