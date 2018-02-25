package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include <gdnative/string.h>
*/
import "C"

type String struct {
	base *C.godot_string
}

func (t *String) getBase() *C.godot_string {
	return t.base
}

// NewString godot_string_new [[godot_string * r_dest]]

//func NewString(dest String, ) *String {
//	return &String{}
//}

// NewStringCopy godot_string_new_copy [[godot_string * r_dest] [const godot_string * p_src]]

//func NewStringCopy(dest String, src ConstString, ) *String {
//	return &String{}
//}

// NewStringWithWideString godot_string_new_with_wide_string [[godot_string * r_dest] [const wchar_t * p_contents] [const int p_size]]

//func NewStringWithWideString(dest String, contents ConstWcharT, size ConstInt, ) *String {
//	return &String{}
//}

// OperatorIndex godot_string_operator_index [[godot_string * p_self] [const godot_int p_idx]]

// OperatorIndexConst godot_string_operator_index_const [[const godot_string * p_self] [const godot_int p_idx]]

// WideStr godot_string_wide_str [[const godot_string * p_self]]

// OperatorEqual godot_string_operator_equal [[const godot_string * p_self] [const godot_string * p_b]]

// OperatorLess godot_string_operator_less [[const godot_string * p_self] [const godot_string * p_b]]

// OperatorPlus godot_string_operator_plus [[const godot_string * p_self] [const godot_string * p_b]]

// Length godot_string_length [[const godot_string * p_self]]

// CasecmpTo godot_string_casecmp_to [[const godot_string * p_self] [const godot_string * p_str]]

// NocasecmpTo godot_string_nocasecmp_to [[const godot_string * p_self] [const godot_string * p_str]]

// NaturalnocasecmpTo godot_string_naturalnocasecmp_to [[const godot_string * p_self] [const godot_string * p_str]]

// BeginsWith godot_string_begins_with [[const godot_string * p_self] [const godot_string * p_string]]

// BeginsWithCharArray godot_string_begins_with_char_array [[const godot_string * p_self] [const char * p_char_array]]

// Bigrams godot_string_bigrams [[const godot_string * p_self]]

// Chr godot_string_chr [[wchar_t p_character]]

// EndsWith godot_string_ends_with [[const godot_string * p_self] [const godot_string * p_string]]

// Find godot_string_find [[const godot_string * p_self] [godot_string p_what]]

// FindFrom godot_string_find_from [[const godot_string * p_self] [godot_string p_what] [godot_int p_from]]

// Findmk godot_string_findmk [[const godot_string * p_self] [const godot_array * p_keys]]

// FindmkFrom godot_string_findmk_from [[const godot_string * p_self] [const godot_array * p_keys] [godot_int p_from]]

// FindmkFromInPlace godot_string_findmk_from_in_place [[const godot_string * p_self] [const godot_array * p_keys] [godot_int p_from] [godot_int * r_key]]

// Findn godot_string_findn [[const godot_string * p_self] [godot_string p_what]]

// FindnFrom godot_string_findn_from [[const godot_string * p_self] [godot_string p_what] [godot_int p_from]]

// FindLast godot_string_find_last [[const godot_string * p_self] [godot_string p_what]]

// Format godot_string_format [[const godot_string * p_self] [const godot_variant * p_values]]

// FormatWithCustomPlaceholder godot_string_format_with_custom_placeholder [[const godot_string * p_self] [const godot_variant * p_values] [const char * p_placeholder]]

// HexEncodeBuffer godot_string_hex_encode_buffer [[const uint8_t * p_buffer] [godot_int p_len]]

// HexToInt godot_string_hex_to_int [[const godot_string * p_self]]

// HexToIntWithoutPrefix godot_string_hex_to_int_without_prefix [[const godot_string * p_self]]

// Insert godot_string_insert [[const godot_string * p_self] [godot_int p_at_pos] [godot_string p_string]]

// IsNumeric godot_string_is_numeric [[const godot_string * p_self]]

// IsSubsequenceOf godot_string_is_subsequence_of [[const godot_string * p_self] [const godot_string * p_string]]

// IsSubsequenceOfi godot_string_is_subsequence_ofi [[const godot_string * p_self] [const godot_string * p_string]]

// Lpad godot_string_lpad [[const godot_string * p_self] [godot_int p_min_length]]

// LpadWithCustomCharacter godot_string_lpad_with_custom_character [[const godot_string * p_self] [godot_int p_min_length] [const godot_string * p_character]]

// Match godot_string_match [[const godot_string * p_self] [const godot_string * p_wildcard]]

// Matchn godot_string_matchn [[const godot_string * p_self] [const godot_string * p_wildcard]]

// Md5 godot_string_md5 [[const uint8_t * p_md5]]

// Num godot_string_num [[double p_num]]

// NumInt64 godot_string_num_int64 [[int64_t p_num] [godot_int p_base]]

// NumInt64Capitalized godot_string_num_int64_capitalized [[int64_t p_num] [godot_int p_base] [godot_bool p_capitalize_hex]]

// NumReal godot_string_num_real [[double p_num]]

// NumScientific godot_string_num_scientific [[double p_num]]

// NumWithDecimals godot_string_num_with_decimals [[double p_num] [godot_int p_decimals]]

// PadDecimals godot_string_pad_decimals [[const godot_string * p_self] [godot_int p_digits]]

// PadZeros godot_string_pad_zeros [[const godot_string * p_self] [godot_int p_digits]]

// ReplaceFirst godot_string_replace_first [[const godot_string * p_self] [godot_string p_key] [godot_string p_with]]

// Replace godot_string_replace [[const godot_string * p_self] [godot_string p_key] [godot_string p_with]]

// Replacen godot_string_replacen [[const godot_string * p_self] [godot_string p_key] [godot_string p_with]]

// Rfind godot_string_rfind [[const godot_string * p_self] [godot_string p_what]]

// Rfindn godot_string_rfindn [[const godot_string * p_self] [godot_string p_what]]

// RfindFrom godot_string_rfind_from [[const godot_string * p_self] [godot_string p_what] [godot_int p_from]]

// RfindnFrom godot_string_rfindn_from [[const godot_string * p_self] [godot_string p_what] [godot_int p_from]]

// Rpad godot_string_rpad [[const godot_string * p_self] [godot_int p_min_length]]

// RpadWithCustomCharacter godot_string_rpad_with_custom_character [[const godot_string * p_self] [godot_int p_min_length] [const godot_string * p_character]]

// Similarity godot_string_similarity [[const godot_string * p_self] [const godot_string * p_string]]

// Sprintf godot_string_sprintf [[const godot_string * p_self] [const godot_array * p_values] [godot_bool * p_error]]

// Substr godot_string_substr [[const godot_string * p_self] [godot_int p_from] [godot_int p_chars]]

// ToDouble godot_string_to_double [[const godot_string * p_self]]

// ToFloat godot_string_to_float [[const godot_string * p_self]]

// ToInt godot_string_to_int [[const godot_string * p_self]]

// CamelcaseToUnderscore godot_string_camelcase_to_underscore [[const godot_string * p_self]]

// CamelcaseToUnderscoreLowercased godot_string_camelcase_to_underscore_lowercased [[const godot_string * p_self]]

// Capitalize godot_string_capitalize [[const godot_string * p_self]]

// CharToDouble godot_string_char_to_double [[const char * p_what]]

// CharToInt godot_string_char_to_int [[const char * p_what]]

// WcharToInt godot_string_wchar_to_int [[const wchar_t * p_str]]

// CharToIntWithLen godot_string_char_to_int_with_len [[const char * p_what] [godot_int p_len]]

// CharToInt64WithLen godot_string_char_to_int64_with_len [[const wchar_t * p_str] [int p_len]]

// HexToInt64 godot_string_hex_to_int64 [[const godot_string * p_self]]

// HexToInt64WithPrefix godot_string_hex_to_int64_with_prefix [[const godot_string * p_self]]

// ToInt64 godot_string_to_int64 [[const godot_string * p_self]]

// UnicodeCharToDouble godot_string_unicode_char_to_double [[const wchar_t * p_str] [const wchar_t ** r_end]]

// GetSliceCount godot_string_get_slice_count [[const godot_string * p_self] [godot_string p_splitter]]

// GetSlice godot_string_get_slice [[const godot_string * p_self] [godot_string p_splitter] [godot_int p_slice]]

// GetSlicec godot_string_get_slicec [[const godot_string * p_self] [wchar_t p_splitter] [godot_int p_slice]]

// Split godot_string_split [[const godot_string * p_self] [const godot_string * p_splitter]]

// SplitAllowEmpty godot_string_split_allow_empty [[const godot_string * p_self] [const godot_string * p_splitter]]

// SplitFloats godot_string_split_floats [[const godot_string * p_self] [const godot_string * p_splitter]]

// SplitFloatsAllowsEmpty godot_string_split_floats_allows_empty [[const godot_string * p_self] [const godot_string * p_splitter]]

// SplitFloatsMk godot_string_split_floats_mk [[const godot_string * p_self] [const godot_array * p_splitters]]

// SplitFloatsMkAllowsEmpty godot_string_split_floats_mk_allows_empty [[const godot_string * p_self] [const godot_array * p_splitters]]

// SplitInts godot_string_split_ints [[const godot_string * p_self] [const godot_string * p_splitter]]

// SplitIntsAllowsEmpty godot_string_split_ints_allows_empty [[const godot_string * p_self] [const godot_string * p_splitter]]

// SplitIntsMk godot_string_split_ints_mk [[const godot_string * p_self] [const godot_array * p_splitters]]

// SplitIntsMkAllowsEmpty godot_string_split_ints_mk_allows_empty [[const godot_string * p_self] [const godot_array * p_splitters]]

// SplitSpaces godot_string_split_spaces [[const godot_string * p_self]]

// CharLowercase godot_string_char_lowercase [[wchar_t p_char]]

// CharUppercase godot_string_char_uppercase [[wchar_t p_char]]

// ToLower godot_string_to_lower [[const godot_string * p_self]]

// ToUpper godot_string_to_upper [[const godot_string * p_self]]

// GetBasename godot_string_get_basename [[const godot_string * p_self]]

// GetExtension godot_string_get_extension [[const godot_string * p_self]]

// Left godot_string_left [[const godot_string * p_self] [godot_int p_pos]]

// OrdAt godot_string_ord_at [[const godot_string * p_self] [godot_int p_idx]]

// PlusFile godot_string_plus_file [[const godot_string * p_self] [const godot_string * p_file]]

// Right godot_string_right [[const godot_string * p_self] [godot_int p_pos]]

// StripEdges godot_string_strip_edges [[const godot_string * p_self] [godot_bool p_left] [godot_bool p_right]]

// StripEscapes godot_string_strip_escapes [[const godot_string * p_self]]

// Erase godot_string_erase [[godot_string * p_self] [godot_int p_pos] [godot_int p_chars]]

// Ascii godot_string_ascii [[const godot_string * p_self]]

// AsciiExtended godot_string_ascii_extended [[const godot_string * p_self]]

// Utf8 godot_string_utf8 [[const godot_string * p_self]]

// ParseUtf8 godot_string_parse_utf8 [[godot_string * p_self] [const char * p_utf8]]

// ParseUtf8WithLen godot_string_parse_utf8_with_len [[godot_string * p_self] [const char * p_utf8] [godot_int p_len]]

// CharsToUtf8 godot_string_chars_to_utf8 [[const char * p_utf8]]

// CharsToUtf8WithLen godot_string_chars_to_utf8_with_len [[const char * p_utf8] [godot_int p_len]]

// Hash godot_string_hash [[const godot_string * p_self]]

// Hash64 godot_string_hash64 [[const godot_string * p_self]]

// HashChars godot_string_hash_chars [[const char * p_cstr]]

// HashCharsWithLen godot_string_hash_chars_with_len [[const char * p_cstr] [godot_int p_len]]

// HashUtf8Chars godot_string_hash_utf8_chars [[const wchar_t * p_str]]

// HashUtf8CharsWithLen godot_string_hash_utf8_chars_with_len [[const wchar_t * p_str] [godot_int p_len]]

// Md5Buffer godot_string_md5_buffer [[const godot_string * p_self]]

// Md5Text godot_string_md5_text [[const godot_string * p_self]]

// Sha256Buffer godot_string_sha256_buffer [[const godot_string * p_self]]

// Sha256Text godot_string_sha256_text [[const godot_string * p_self]]

// Empty godot_string_empty [[const godot_string * p_self]]

// GetBaseDir godot_string_get_base_dir [[const godot_string * p_self]]

// GetFile godot_string_get_file [[const godot_string * p_self]]

// HumanizeSize godot_string_humanize_size [[size_t p_size]]

// IsAbsPath godot_string_is_abs_path [[const godot_string * p_self]]

// IsRelPath godot_string_is_rel_path [[const godot_string * p_self]]

// IsResourceFile godot_string_is_resource_file [[const godot_string * p_self]]

// PathTo godot_string_path_to [[const godot_string * p_self] [const godot_string * p_path]]

// PathToFile godot_string_path_to_file [[const godot_string * p_self] [const godot_string * p_path]]

// SimplifyPath godot_string_simplify_path [[const godot_string * p_self]]

// CEscape godot_string_c_escape [[const godot_string * p_self]]

// CEscapeMultiline godot_string_c_escape_multiline [[const godot_string * p_self]]

// CUnescape godot_string_c_unescape [[const godot_string * p_self]]

// HttpEscape godot_string_http_escape [[const godot_string * p_self]]

// HttpUnescape godot_string_http_unescape [[const godot_string * p_self]]

// JsonEscape godot_string_json_escape [[const godot_string * p_self]]

// WordWrap godot_string_word_wrap [[const godot_string * p_self] [godot_int p_chars_per_line]]

// XmlEscape godot_string_xml_escape [[const godot_string * p_self]]

// XmlEscapeWithQuotes godot_string_xml_escape_with_quotes [[const godot_string * p_self]]

// XmlUnescape godot_string_xml_unescape [[const godot_string * p_self]]

// PercentDecode godot_string_percent_decode [[const godot_string * p_self]]

// PercentEncode godot_string_percent_encode [[const godot_string * p_self]]

// IsValidFloat godot_string_is_valid_float [[const godot_string * p_self]]

// IsValidHexNumber godot_string_is_valid_hex_number [[const godot_string * p_self] [godot_bool p_with_prefix]]

// IsValidHtmlColor godot_string_is_valid_html_color [[const godot_string * p_self]]

// IsValidIdentifier godot_string_is_valid_identifier [[const godot_string * p_self]]

// IsValidInteger godot_string_is_valid_integer [[const godot_string * p_self]]

// IsValidIpAddress godot_string_is_valid_ip_address [[const godot_string * p_self]]

// Destroy godot_string_destroy [[godot_string * p_self]]

// NewStringName godot_string_name_new [[godot_string_name * r_dest] [const godot_string * p_name]]

//func NewStringName(dest StringName, name ConstString, ) *String {
//	return &String{}
//}

// NewStringNameData godot_string_name_new_data [[godot_string_name * r_dest] [const char * p_name]]

//func NewStringNameData(dest StringName, name ConstChar, ) *String {
//	return &String{}
//}

// NameGetName godot_string_name_get_name [[const godot_string_name * p_self]]

// NameGetHash godot_string_name_get_hash [[const godot_string_name * p_self]]

// NameGetDataUniquePointer godot_string_name_get_data_unique_pointer [[const godot_string_name * p_self]]

// NameOperatorEqual godot_string_name_operator_equal [[const godot_string_name * p_self] [const godot_string_name * p_other]]

// NameOperatorLess godot_string_name_operator_less [[const godot_string_name * p_self] [const godot_string_name * p_other]]

// NameDestroy godot_string_name_destroy [[godot_string_name * p_self]]

type CharString struct {
	base *C.godot_char_string
}

func (t *CharString) getBase() *C.godot_char_string {
	return t.base
}

// Length godot_char_string_length [[const godot_char_string * p_cs]]

// GetData godot_char_string_get_data [[const godot_char_string * p_cs]]

// Destroy godot_char_string_destroy [[godot_char_string * p_cs]]