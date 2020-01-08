//  Copyright (c) 2017 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vellum

const maxCommon = 1<<6 - 1

func encodeCommon(in byte) byte {
	val := byte((int(commonInputs[in]) + 1) % 256)
	if val > maxCommon {
		return 0
	}
	return val
}

func decodeCommon(in byte) byte {
	return commonInputsInv[in-1]
}

var commonInputs = []byte{
	84,  // '\x00'
	85,  // '\x01'
	86,  // '\x02'
	87,  // '\x03'
	88,  // '\x04'
	89,  // '\x05'
	90,  // '\x06'
	91,  // '\x07'
	92,  // '\x08'
	93,  // '\t'
	94,  // '\n'
	95,  // '\x0b'
	96,  // '\x0c'
	97,  // '\r'
	98,  // '\x0e'
	99,  // '\x0f'
	100, // '\x10'
	101, // '\x11'
	102, // '\x12'
	103, // '\x13'
	104, // '\x14'
	105, // '\x15'
	106, // '\x16'
	107, // '\x17'
	108, // '\x18'
	109, // '\x19'
	110, // '\x1a'
	111, // '\x1b'
	112, // '\x1c'
	113, // '\x1d'
	114, // '\x1e'
	115, // '\x1f'
	116, // ' '
	80,  // '!'
	117, // '"'
	118, // '#'
	79,  // '$'
	39,  // '%'
	30,  // '&'
	81,  // "'"
	75,  // '('
	74,  // ')'
	82,  // '*'
	57,  // '+'
	66,  // ','
	16,  // '-'
	12,  // '.'
	2,   // '/'
	19,  // '0'
	20,  // '1'
	21,  // '2'
	27,  // '3'
	32,  // '4'
	29,  // '5'
	35,  // '6'
	36,  // '7'
	37,  // '8'
	34,  // '9'
	24,  // ':'
	73,  // ';'
	119, // '<'
	23,  // '='
	120, // '>'
	40,  // '?'
	83,  // '@'
	44,  // 'A'
	48,  // 'B'
	42,  // 'C'
	43,  // 'D'
	49,  // 'E'
	46,  // 'F'
	62,  // 'G'
	61,  // 'H'
	47,  // 'I'
	69,  // 'J'
	68,  // 'K'
	58,  // 'L'
	56,  // 'M'
	55,  // 'N'
	59,  // 'O'
	51,  // 'P'
	72,  // 'Q'
	54,  // 'R'
	45,  // 'S'
	52,  // 'T'
	64,  // 'U'
	65,  // 'V'
	63,  // 'W'
	71,  // 'X'
	67,  // 'Y'
	70,  // 'Z'
	77,  // '['
	121, // '\\'
	78,  // ']'
	122, // '^'
	31,  // '_'
	123, // '`'
	4,   // 'a'
	25,  // 'b'
	9,   // 'c'
	17,  // 'd'
	1,   // 'e'
	26,  // 'f'
	22,  // 'g'
	13,  // 'h'
	7,   // 'i'
	50,  // 'j'
	38,  // 'k'
	14,  // 'l'
	15,  // 'm'
	10,  // 'n'
	3,   // 'o'
	8,   // 'p'
	60,  // 'q'
	6,   // 'r'
	5,   // 's'
	0,   // 't'
	18,  // 'u'
	33,  // 'v'
	11,  // 'w'
	41,  // 'x'
	28,  // 'y'
	53,  // 'z'
	124, // '{'
	125, // '|'
	126, // '}'
	76,  // '~'
	127, // '\x7f'
	128, // '\x80'
	129, // '\x81'
	130, // '\x82'
	131, // '\x83'
	132, // '\x84'
	133, // '\x85'
	134, // '\x86'
	135, // '\x87'
	136, // '\x88'
	137, // '\x89'
	138, // '\x8a'
	139, // '\x8b'
	140, // '\x8c'
	141, // '\x8d'
	142, // '\x8e'
	143, // '\x8f'
	144, // '\x90'
	145, // '\x91'
	146, // '\x92'
	147, // '\x93'
	148, // '\x94'
	149, // '\x95'
	150, // '\x96'
	151, // '\x97'
	152, // '\x98'
	153, // '\x99'
	154, // '\x9a'
	155, // '\x9b'
	156, // '\x9c'
	157, // '\x9d'
	158, // '\x9e'
	159, // '\x9f'
	160, // '\xa0'
	161, // '¡'
	162, // '¢'
	163, // '£'
	164, // '¤'
	165, // '¥'
	166, // '¦'
	167, // '§'
	168, // '¨'
	169, // '©'
	170, // 'ª'
	171, // '«'
	172, // '¬'
	173, // '\xad'
	174, // '®'
	175, // '¯'
	176, // '°'
	177, // '±'
	178, // '²'
	179, // '³'
	180, // '´'
	181, // 'µ'
	182, // '¶'
	183, // '·'
	184, // '¸'
	185, // '¹'
	186, // 'º'
	187, // '»'
	188, // '¼'
	189, // '½'
	190, // '¾'
	191, // '¿'
	192, // 'À'
	193, // 'Á'
	194, // 'Â'
	195, // 'Ã'
	196, // 'Ä'
	197, // 'Å'
	198, // 'Æ'
	199, // 'Ç'
	200, // 'È'
	201, // 'É'
	202, // 'Ê'
	203, // 'Ë'
	204, // 'Ì'
	205, // 'Í'
	206, // 'Î'
	207, // 'Ï'
	208, // 'Ð'
	209, // 'Ñ'
	210, // 'Ò'
	211, // 'Ó'
	212, // 'Ô'
	213, // 'Õ'
	214, // 'Ö'
	215, // '×'
	216, // 'Ø'
	217, // 'Ù'
	218, // 'Ú'
	219, // 'Û'
	220, // 'Ü'
	221, // 'Ý'
	222, // 'Þ'
	223, // 'ß'
	224, // 'à'
	225, // 'á'
	226, // 'â'
	227, // 'ã'
	228, // 'ä'
	229, // 'å'
	230, // 'æ'
	231, // 'ç'
	232, // 'è'
	233, // 'é'
	234, // 'ê'
	235, // 'ë'
	236, // 'ì'
	237, // 'í'
	238, // 'î'
	239, // 'ï'
	240, // 'ð'
	241, // 'ñ'
	242, // 'ò'
	243, // 'ó'
	244, // 'ô'
	245, // 'õ'
	246, // 'ö'
	247, // '÷'
	248, // 'ø'
	249, // 'ù'
	250, // 'ú'
	251, // 'û'
	252, // 'ü'
	253, // 'ý'
	254, // 'þ'
	255, // 'ÿ'
}

var commonInputsInv = []byte{
	't',
	'e',
	'/',
	'o',
	'a',
	's',
	'r',
	'i',
	'p',
	'c',
	'n',
	'w',
	'.',
	'h',
	'l',
	'm',
	'-',
	'd',
	'u',
	'0',
	'1',
	'2',
	'g',
	'=',
	':',
	'b',
	'f',
	'3',
	'y',
	'5',
	'&',
	'_',
	'4',
	'v',
	'9',
	'6',
	'7',
	'8',
	'k',
	'%',
	'?',
	'x',
	'C',
	'D',
	'A',
	'S',
	'F',
	'I',
	'B',
	'E',
	'j',
	'P',
	'T',
	'z',
	'R',
	'N',
	'M',
	'+',
	'L',
	'O',
	'q',
	'H',
	'G',
	'W',
	'U',
	'V',
	',',
	'Y',
	'K',
	'J',
	'Z',
	'X',
	'Q',
	';',
	')',
	'(',
	'~',
	'[',
	']',
	'$',
	'!',
	'\'',
	'*',
	'@',
	'\x00',
	'\x01',
	'\x02',
	'\x03',
	'\x04',
	'\x05',
	'\x06',
	'\x07',
	'\x08',
	'\t',
	'\n',
	'\x0b',
	'\x0c',
	'\r',
	'\x0e',
	'\x0f',
	'\x10',
	'\x11',
	'\x12',
	'\x13',
	'\x14',
	'\x15',
	'\x16',
	'\x17',
	'\x18',
	'\x19',
	'\x1a',
	'\x1b',
	'\x1c',
	'\x1d',
	'\x1e',
	'\x1f',
	' ',
	'"',
	'#',
	'<',
	'>',
	'\\',
	'^',
	'`',
	'{',
	'|',
	'}',
	'\x7f',
	'\x80',
	'\x81',
	'\x82',
	'\x83',
	'\x84',
	'\x85',
	'\x86',
	'\x87',
	'\x88',
	'\x89',
	'\x8a',
	'\x8b',
	'\x8c',
	'\x8d',
	'\x8e',
	'\x8f',
	'\x90',
	'\x91',
	'\x92',
	'\x93',
	'\x94',
	'\x95',
	'\x96',
	'\x97',
	'\x98',
	'\x99',
	'\x9a',
	'\x9b',
	'\x9c',
	'\x9d',
	'\x9e',
	'\x9f',
	'\xa0',
	'\xa1',
	'\xa2',
	'\xa3',
	'\xa4',
	'\xa5',
	'\xa6',
	'\xa7',
	'\xa8',
	'\xa9',
	'\xaa',
	'\xab',
	'\xac',
	'\xad',
	'\xae',
	'\xaf',
	'\xb0',
	'\xb1',
	'\xb2',
	'\xb3',
	'\xb4',
	'\xb5',
	'\xb6',
	'\xb7',
	'\xb8',
	'\xb9',
	'\xba',
	'\xbb',
	'\xbc',
	'\xbd',
	'\xbe',
	'\xbf',
	'\xc0',
	'\xc1',
	'\xc2',
	'\xc3',
	'\xc4',
	'\xc5',
	'\xc6',
	'\xc7',
	'\xc8',
	'\xc9',
	'\xca',
	'\xcb',
	'\xcc',
	'\xcd',
	'\xce',
	'\xcf',
	'\xd0',
	'\xd1',
	'\xd2',
	'\xd3',
	'\xd4',
	'\xd5',
	'\xd6',
	'\xd7',
	'\xd8',
	'\xd9',
	'\xda',
	'\xdb',
	'\xdc',
	'\xdd',
	'\xde',
	'\xdf',
	'\xe0',
	'\xe1',
	'\xe2',
	'\xe3',
	'\xe4',
	'\xe5',
	'\xe6',
	'\xe7',
	'\xe8',
	'\xe9',
	'\xea',
	'\xeb',
	'\xec',
	'\xed',
	'\xee',
	'\xef',
	'\xf0',
	'\xf1',
	'\xf2',
	'\xf3',
	'\xf4',
	'\xf5',
	'\xf6',
	'\xf7',
	'\xf8',
	'\xf9',
	'\xfa',
	'\xfb',
	'\xfc',
	'\xfd',
	'\xfe',
	'\xff',
}