// generated by "go run gen.go". DO NOT EDIT.

package ccitt

// Each decodeTable is represented by an array of [2]int16's: a binary tree.
// Each array element (other than element 0, which means invalid) is a branch
// node in that tree. The root node is always element 1 (the second element).
//
// To walk the tree, look at the next bit in the bit stream, using it to select
// the first or second element of the [2]int16. If that int16 is 0, we have an
// invalid code. If it is positive, go to that branch node. If it is negative,
// then we have a leaf node, whose value is the bitwise complement (the ^
// operator) of that int16.
//
// Comments above each decodeTable also show the same structure visually. The
// "b123" lines show the 123'rd branch node. The "=XXXXX" lines show an invalid
// code. The "=v1234" lines show a leaf node with value 1234. When reading the
// bit stream, a 0 or 1 bit means to go up or down, as you move left to right.
//
// For example, in modeDecodeTable, branch node b005 is three steps up from the
// root node, meaning that we have already seen "000". If the next bit is "0"
// then we move to branch node b006. Otherwise, the next bit is "1", and we
// move to the leaf node v0000 (also known as the modePass constant). Indeed,
// the bits that encode modePass are "0001".
//
// Tables 1, 2 and 3 come from the "ITU-T Recommendation T.6: FACSIMILE CODING
// SCHEMES AND CODING CONTROL FUNCTIONS FOR GROUP 4 FACSIMILE APPARATUS"
// specification:
//
// https://www.itu.int/rec/dologin_pub.asp?lang=e&id=T-REC-T.6-198811-I!!PDF-E&type=items

// modeDecodeTable represents Table 1 and the End-of-Line code.
//
//	+=XXXXX
//
// b015                       +-+
//
//	| +=v0010
//
// b014                     +-+
//
//	| +=XXXXX
//
// b013                   +-+
//
//	| +=XXXXX
//
// b012                 +-+
//
//	| +=XXXXX
//
// b011               +-+
//
//	| +=XXXXX
//
// b009             +-+
//
//	| +=v0009
//
// b007           +-+
//
//	| | +=v0008
//
// b010           | +-+
//
//	|   +=v0005
//
// b006         +-+
//
//	| | +=v0007
//
// b008         | +-+
//
//	|   +=v0004
//
// b005       +-+
//
//	| +=v0000
//
// b003     +-+
//
//	| +=v0001
//
// b002   +-+
//
//	| | +=v0006
//
// b004   | +-+
//
//	|   +=v0003
//
// b001 +-+
//
//	+=v0002
var modeDecodeTable = [...][2]int16{
	0:  {0, 0},
	1:  {2, ^2},
	2:  {3, 4},
	3:  {5, ^1},
	4:  {^6, ^3},
	5:  {6, ^0},
	6:  {7, 8},
	7:  {9, 10},
	8:  {^7, ^4},
	9:  {11, ^9},
	10: {^8, ^5},
	11: {12, 0},
	12: {13, 0},
	13: {14, 0},
	14: {15, 0},
	15: {0, ^10},
}

// whiteDecodeTable represents Tables 2 and 3 for a white run.
//
//	+=XXXXX
//
// b059               +-+
//
//	| |     +=v1792
//
// b096               | |   +-+
//
//	| |   | | +=v1984
//
// b100               | |   | +-+
//
//	| |   |   +=v2048
//
// b094               | | +-+
//
//	| | | |   +=v2112
//
// b101               | | | | +-+
//
//	| | | | | +=v2176
//
// b097               | | | +-+
//
//	| | |   | +=v2240
//
// b102               | | |   +-+
//
//	| | |     +=v2304
//
// b085               | +-+
//
//	|   |   +=v1856
//
// b098               |   | +-+
//
//	|   | | +=v1920
//
// b095               |   +-+
//
//	|     |   +=v2368
//
// b103               |     | +-+
//
//	|     | | +=v2432
//
// b099               |     +-+
//
//	|       | +=v2496
//
// b104               |       +-+
//
//	|         +=v2560
//
// b040             +-+
//
//	| | +=v0029
//
// b060             | +-+
//
//	|   +=v0030
//
// b026           +-+
//
//	| |   +=v0045
//
// b061           | | +-+
//
//	| | | +=v0046
//
// b041           | +-+
//
//	|   +=v0022
//
// b016         +-+
//
//	| |   +=v0023
//
// b042         | | +-+
//
//	| | | | +=v0047
//
// b062         | | | +-+
//
//	| | |   +=v0048
//
// b027         | +-+
//
//	|   +=v0013
//
// b008       +-+
//
//	| |     +=v0020
//
// b043       | |   +-+
//
//	| |   | | +=v0033
//
// b063       | |   | +-+
//
//	| |   |   +=v0034
//
// b028       | | +-+
//
//	| | | |   +=v0035
//
// b064       | | | | +-+
//
//	| | | | | +=v0036
//
// b044       | | | +-+
//
//	| | |   | +=v0037
//
// b065       | | |   +-+
//
//	| | |     +=v0038
//
// b017       | +-+
//
//	|   |   +=v0019
//
// b045       |   | +-+
//
//	|   | | | +=v0031
//
// b066       |   | | +-+
//
//	|   | |   +=v0032
//
// b029       |   +-+
//
//	|     +=v0001
//
// b004     +-+
//
//	| |     +=v0012
//
// b030     | |   +-+
//
//	| |   | |   +=v0053
//
// b067     | |   | | +-+
//
//	| |   | | | +=v0054
//
// b046     | |   | +-+
//
//	| |   |   +=v0026
//
// b018     | | +-+
//
//	| | | |     +=v0039
//
// b068     | | | |   +-+
//
//	| | | |   | +=v0040
//
// b047     | | | | +-+
//
//	| | | | | | +=v0041
//
// b069     | | | | | +-+
//
//	| | | | |   +=v0042
//
// b031     | | | +-+
//
//	| | |   |   +=v0043
//
// b070     | | |   | +-+
//
//	| | |   | | +=v0044
//
// b048     | | |   +-+
//
//	| | |     +=v0021
//
// b009     | +-+
//
//	|   |     +=v0028
//
// b049     |   |   +-+
//
//	|   |   | | +=v0061
//
// b071     |   |   | +-+
//
//	|   |   |   +=v0062
//
// b032     |   | +-+
//
//	|   | | |   +=v0063
//
// b072     |   | | | +-+
//
//	|   | | | | +=v0000
//
// b050     |   | | +-+
//
//	|   | |   | +=v0320
//
// b073     |   | |   +-+
//
//	|   | |     +=v0384
//
// b019     |   +-+
//
//	|     +=v0010
//
// b002   +-+
//
//	| |     +=v0011
//
// b020   | |   +-+
//
//	| |   | |   +=v0027
//
// b051   | |   | | +-+
//
//	| |   | | | | +=v0059
//
// b074   | |   | | | +-+
//
//	| |   | | |   +=v0060
//
// b033   | |   | +-+
//
//	| |   |   |     +=v1472
//
// b086   | |   |   |   +-+
//
//	| |   |   |   | +=v1536
//
// b075   | |   |   | +-+
//
//	| |   |   | | | +=v1600
//
// b087   | |   |   | | +-+
//
//	| |   |   | |   +=v1728
//
// b052   | |   |   +-+
//
//	| |   |     +=v0018
//
// b010   | | +-+
//
//	| | | |     +=v0024
//
// b053   | | | |   +-+
//
//	| | | |   | | +=v0049
//
// b076   | | | |   | +-+
//
//	| | | |   |   +=v0050
//
// b034   | | | | +-+
//
//	| | | | | |   +=v0051
//
// b077   | | | | | | +-+
//
//	| | | | | | | +=v0052
//
// b054   | | | | | +-+
//
//	| | | | |   +=v0025
//
// b021   | | | +-+
//
//	| | |   |     +=v0055
//
// b078   | | |   |   +-+
//
//	| | |   |   | +=v0056
//
// b055   | | |   | +-+
//
//	| | |   | | | +=v0057
//
// b079   | | |   | | +-+
//
//	| | |   | |   +=v0058
//
// b035   | | |   +-+
//
//	| | |     +=v0192
//
// b005   | +-+
//
//	|   |     +=v1664
//
// b036   |   |   +-+
//
//	|   |   | |   +=v0448
//
// b080   |   |   | | +-+
//
//	|   |   | | | +=v0512
//
// b056   |   |   | +-+
//
//	|   |   |   |   +=v0704
//
// b088   |   |   |   | +-+
//
//	|   |   |   | | +=v0768
//
// b081   |   |   |   +-+
//
//	|   |   |     +=v0640
//
// b022   |   | +-+
//
//	|   | | |     +=v0576
//
// b082   |   | | |   +-+
//
//	|   | | |   | | +=v0832
//
// b089   |   | | |   | +-+
//
//	|   | | |   |   +=v0896
//
// b057   |   | | | +-+
//
//	|   | | | | |   +=v0960
//
// b090   |   | | | | | +-+
//
//	|   | | | | | | +=v1024
//
// b083   |   | | | | +-+
//
//	|   | | | |   | +=v1088
//
// b091   |   | | | |   +-+
//
//	|   | | | |     +=v1152
//
// b037   |   | | +-+
//
//	|   | |   |     +=v1216
//
// b092   |   | |   |   +-+
//
//	|   | |   |   | +=v1280
//
// b084   |   | |   | +-+
//
//	|   | |   | | | +=v1344
//
// b093   |   | |   | | +-+
//
//	|   | |   | |   +=v1408
//
// b058   |   | |   +-+
//
//	|   | |     +=v0256
//
// b011   |   +-+
//
//	|     +=v0002
//
// b001 +-+
//
//	|     +=v0003
//
// b012   |   +-+
//
//	|   | | +=v0128
//
// b023   |   | +-+
//
//	|   |   +=v0008
//
// b006   | +-+
//
//	| | |   +=v0009
//
// b024   | | | +-+
//
//	| | | | | +=v0016
//
// b038   | | | | +-+
//
//	| | | |   +=v0017
//
// b013   | | +-+
//
//	| |   +=v0004
//
// b003   +-+
//
//	|   +=v0005
//
// b014     | +-+
//
//	| | |   +=v0014
//
// b039     | | | +-+
//
//	| | | | +=v0015
//
// b025     | | +-+
//
//	| |   +=v0064
//
// b007     +-+
//
//	| +=v0006
//
// b015       +-+
//
//	+=v0007
var whiteDecodeTable = [...][2]int16{
	0:   {0, 0},
	1:   {2, 3},
	2:   {4, 5},
	3:   {6, 7},
	4:   {8, 9},
	5:   {10, 11},
	6:   {12, 13},
	7:   {14, 15},
	8:   {16, 17},
	9:   {18, 19},
	10:  {20, 21},
	11:  {22, ^2},
	12:  {^3, 23},
	13:  {24, ^4},
	14:  {^5, 25},
	15:  {^6, ^7},
	16:  {26, 27},
	17:  {28, 29},
	18:  {30, 31},
	19:  {32, ^10},
	20:  {^11, 33},
	21:  {34, 35},
	22:  {36, 37},
	23:  {^128, ^8},
	24:  {^9, 38},
	25:  {39, ^64},
	26:  {40, 41},
	27:  {42, ^13},
	28:  {43, 44},
	29:  {45, ^1},
	30:  {^12, 46},
	31:  {47, 48},
	32:  {49, 50},
	33:  {51, 52},
	34:  {53, 54},
	35:  {55, ^192},
	36:  {^1664, 56},
	37:  {57, 58},
	38:  {^16, ^17},
	39:  {^14, ^15},
	40:  {59, 60},
	41:  {61, ^22},
	42:  {^23, 62},
	43:  {^20, 63},
	44:  {64, 65},
	45:  {^19, 66},
	46:  {67, ^26},
	47:  {68, 69},
	48:  {70, ^21},
	49:  {^28, 71},
	50:  {72, 73},
	51:  {^27, 74},
	52:  {75, ^18},
	53:  {^24, 76},
	54:  {77, ^25},
	55:  {78, 79},
	56:  {80, 81},
	57:  {82, 83},
	58:  {84, ^256},
	59:  {0, 85},
	60:  {^29, ^30},
	61:  {^45, ^46},
	62:  {^47, ^48},
	63:  {^33, ^34},
	64:  {^35, ^36},
	65:  {^37, ^38},
	66:  {^31, ^32},
	67:  {^53, ^54},
	68:  {^39, ^40},
	69:  {^41, ^42},
	70:  {^43, ^44},
	71:  {^61, ^62},
	72:  {^63, ^0},
	73:  {^320, ^384},
	74:  {^59, ^60},
	75:  {86, 87},
	76:  {^49, ^50},
	77:  {^51, ^52},
	78:  {^55, ^56},
	79:  {^57, ^58},
	80:  {^448, ^512},
	81:  {88, ^640},
	82:  {^576, 89},
	83:  {90, 91},
	84:  {92, 93},
	85:  {94, 95},
	86:  {^1472, ^1536},
	87:  {^1600, ^1728},
	88:  {^704, ^768},
	89:  {^832, ^896},
	90:  {^960, ^1024},
	91:  {^1088, ^1152},
	92:  {^1216, ^1280},
	93:  {^1344, ^1408},
	94:  {96, 97},
	95:  {98, 99},
	96:  {^1792, 100},
	97:  {101, 102},
	98:  {^1856, ^1920},
	99:  {103, 104},
	100: {^1984, ^2048},
	101: {^2112, ^2176},
	102: {^2240, ^2304},
	103: {^2368, ^2432},
	104: {^2496, ^2560},
}

// blackDecodeTable represents Tables 2 and 3 for a black run.
//
//	+=XXXXX
//
// b017               +-+
//
//	| |     +=v1792
//
// b042               | |   +-+
//
//	| |   | | +=v1984
//
// b063               | |   | +-+
//
//	| |   |   +=v2048
//
// b029               | | +-+
//
//	| | | |   +=v2112
//
// b064               | | | | +-+
//
//	| | | | | +=v2176
//
// b043               | | | +-+
//
//	| | |   | +=v2240
//
// b065               | | |   +-+
//
//	| | |     +=v2304
//
// b022               | +-+
//
//	|   |   +=v1856
//
// b044               |   | +-+
//
//	|   | | +=v1920
//
// b030               |   +-+
//
//	|     |   +=v2368
//
// b066               |     | +-+
//
//	|     | | +=v2432
//
// b045               |     +-+
//
//	|       | +=v2496
//
// b067               |       +-+
//
//	|         +=v2560
//
// b013             +-+
//
//	| |     +=v0018
//
// b031             | |   +-+
//
//	| |   | |   +=v0052
//
// b068             | |   | | +-+
//
//	| |   | | | | +=v0640
//
// b095             | |   | | | +-+
//
//	| |   | | |   +=v0704
//
// b046             | |   | +-+
//
//	| |   |   |   +=v0768
//
// b096             | |   |   | +-+
//
//	| |   |   | | +=v0832
//
// b069             | |   |   +-+
//
//	| |   |     +=v0055
//
// b023             | | +-+
//
//	| | | |     +=v0056
//
// b070             | | | |   +-+
//
//	| | | |   | | +=v1280
//
// b097             | | | |   | +-+
//
//	| | | |   |   +=v1344
//
// b047             | | | | +-+
//
//	| | | | | |   +=v1408
//
// b098             | | | | | | +-+
//
//	| | | | | | | +=v1472
//
// b071             | | | | | +-+
//
//	| | | | |   +=v0059
//
// b032             | | | +-+
//
//	| | |   |   +=v0060
//
// b072             | | |   | +-+
//
//	| | |   | | | +=v1536
//
// b099             | | |   | | +-+
//
//	| | |   | |   +=v1600
//
// b048             | | |   +-+
//
//	| | |     +=v0024
//
// b018             | +-+
//
//	|   |     +=v0025
//
// b049             |   |   +-+
//
//	|   |   | |   +=v1664
//
// b100             |   |   | | +-+
//
//	|   |   | | | +=v1728
//
// b073             |   |   | +-+
//
//	|   |   |   +=v0320
//
// b033             |   | +-+
//
//	|   | | |   +=v0384
//
// b074             |   | | | +-+
//
//	|   | | | | +=v0448
//
// b050             |   | | +-+
//
//	|   | |   |   +=v0512
//
// b101             |   | |   | +-+
//
//	|   | |   | | +=v0576
//
// b075             |   | |   +-+
//
//	|   | |     +=v0053
//
// b024             |   +-+
//
//	|     |     +=v0054
//
// b076             |     |   +-+
//
//	|     |   | | +=v0896
//
// b102             |     |   | +-+
//
//	|     |   |   +=v0960
//
// b051             |     | +-+
//
//	|     | | |   +=v1024
//
// b103             |     | | | +-+
//
//	|     | | | | +=v1088
//
// b077             |     | | +-+
//
//	|     | |   | +=v1152
//
// b104             |     | |   +-+
//
//	|     | |     +=v1216
//
// b034             |     +-+
//
//	|       +=v0064
//
// b010           +-+
//
//	| |   +=v0013
//
// b019           | | +-+
//
//	| | | |     +=v0023
//
// b052           | | | |   +-+
//
//	| | | |   | | +=v0050
//
// b078           | | | |   | +-+
//
//	| | | |   |   +=v0051
//
// b035           | | | | +-+
//
//	| | | | | |   +=v0044
//
// b079           | | | | | | +-+
//
//	| | | | | | | +=v0045
//
// b053           | | | | | +-+
//
//	| | | | |   | +=v0046
//
// b080           | | | | |   +-+
//
//	| | | | |     +=v0047
//
// b025           | | | +-+
//
//	| | |   |     +=v0057
//
// b081           | | |   |   +-+
//
//	| | |   |   | +=v0058
//
// b054           | | |   | +-+
//
//	| | |   | | | +=v0061
//
// b082           | | |   | | +-+
//
//	| | |   | |   +=v0256
//
// b036           | | |   +-+
//
//	| | |     +=v0016
//
// b014           | +-+
//
//	|   |     +=v0017
//
// b037           |   |   +-+
//
//	|   |   | |   +=v0048
//
// b083           |   |   | | +-+
//
//	|   |   | | | +=v0049
//
// b055           |   |   | +-+
//
//	|   |   |   | +=v0062
//
// b084           |   |   |   +-+
//
//	|   |   |     +=v0063
//
// b026           |   | +-+
//
//	|   | | |     +=v0030
//
// b085           |   | | |   +-+
//
//	|   | | |   | +=v0031
//
// b056           |   | | | +-+
//
//	|   | | | | | +=v0032
//
// b086           |   | | | | +-+
//
//	|   | | | |   +=v0033
//
// b038           |   | | +-+
//
//	|   | |   |   +=v0040
//
// b087           |   | |   | +-+
//
//	|   | |   | | +=v0041
//
// b057           |   | |   +-+
//
//	|   | |     +=v0022
//
// b020           |   +-+
//
//	|     +=v0014
//
// b008         +-+
//
//	| |   +=v0010
//
// b015         | | +-+
//
//	| | | +=v0011
//
// b011         | +-+
//
//	|   |     +=v0015
//
// b027         |   |   +-+
//
//	|   |   | |     +=v0128
//
// b088         |   |   | |   +-+
//
//	|   |   | |   | +=v0192
//
// b058         |   |   | | +-+
//
//	|   |   | | | | +=v0026
//
// b089         |   |   | | | +-+
//
//	|   |   | | |   +=v0027
//
// b039         |   |   | +-+
//
//	|   |   |   |   +=v0028
//
// b090         |   |   |   | +-+
//
//	|   |   |   | | +=v0029
//
// b059         |   |   |   +-+
//
//	|   |   |     +=v0019
//
// b021         |   | +-+
//
//	|   | | |     +=v0020
//
// b060         |   | | |   +-+
//
//	|   | | |   | | +=v0034
//
// b091         |   | | |   | +-+
//
//	|   | | |   |   +=v0035
//
// b040         |   | | | +-+
//
//	|   | | | | |   +=v0036
//
// b092         |   | | | | | +-+
//
//	|   | | | | | | +=v0037
//
// b061         |   | | | | +-+
//
//	|   | | | |   | +=v0038
//
// b093         |   | | | |   +-+
//
//	|   | | | |     +=v0039
//
// b028         |   | | +-+
//
//	|   | |   |   +=v0021
//
// b062         |   | |   | +-+
//
//	|   | |   | | | +=v0042
//
// b094         |   | |   | | +-+
//
//	|   | |   | |   +=v0043
//
// b041         |   | |   +-+
//
//	|   | |     +=v0000
//
// b016         |   +-+
//
//	|     +=v0012
//
// b006       +-+
//
//	| |   +=v0009
//
// b012       | | +-+
//
//	| | | +=v0008
//
// b009       | +-+
//
//	|   +=v0007
//
// b004     +-+
//
//	| | +=v0006
//
// b007     | +-+
//
//	|   +=v0005
//
// b002   +-+
//
//	| | +=v0001
//
// b005   | +-+
//
//	|   +=v0004
//
// b001 +-+
//
//	| +=v0003
//
// b003   +-+
//
//	+=v0002
var blackDecodeTable = [...][2]int16{
	0:   {0, 0},
	1:   {2, 3},
	2:   {4, 5},
	3:   {^3, ^2},
	4:   {6, 7},
	5:   {^1, ^4},
	6:   {8, 9},
	7:   {^6, ^5},
	8:   {10, 11},
	9:   {12, ^7},
	10:  {13, 14},
	11:  {15, 16},
	12:  {^9, ^8},
	13:  {17, 18},
	14:  {19, 20},
	15:  {^10, ^11},
	16:  {21, ^12},
	17:  {0, 22},
	18:  {23, 24},
	19:  {^13, 25},
	20:  {26, ^14},
	21:  {27, 28},
	22:  {29, 30},
	23:  {31, 32},
	24:  {33, 34},
	25:  {35, 36},
	26:  {37, 38},
	27:  {^15, 39},
	28:  {40, 41},
	29:  {42, 43},
	30:  {44, 45},
	31:  {^18, 46},
	32:  {47, 48},
	33:  {49, 50},
	34:  {51, ^64},
	35:  {52, 53},
	36:  {54, ^16},
	37:  {^17, 55},
	38:  {56, 57},
	39:  {58, 59},
	40:  {60, 61},
	41:  {62, ^0},
	42:  {^1792, 63},
	43:  {64, 65},
	44:  {^1856, ^1920},
	45:  {66, 67},
	46:  {68, 69},
	47:  {70, 71},
	48:  {72, ^24},
	49:  {^25, 73},
	50:  {74, 75},
	51:  {76, 77},
	52:  {^23, 78},
	53:  {79, 80},
	54:  {81, 82},
	55:  {83, 84},
	56:  {85, 86},
	57:  {87, ^22},
	58:  {88, 89},
	59:  {90, ^19},
	60:  {^20, 91},
	61:  {92, 93},
	62:  {^21, 94},
	63:  {^1984, ^2048},
	64:  {^2112, ^2176},
	65:  {^2240, ^2304},
	66:  {^2368, ^2432},
	67:  {^2496, ^2560},
	68:  {^52, 95},
	69:  {96, ^55},
	70:  {^56, 97},
	71:  {98, ^59},
	72:  {^60, 99},
	73:  {100, ^320},
	74:  {^384, ^448},
	75:  {101, ^53},
	76:  {^54, 102},
	77:  {103, 104},
	78:  {^50, ^51},
	79:  {^44, ^45},
	80:  {^46, ^47},
	81:  {^57, ^58},
	82:  {^61, ^256},
	83:  {^48, ^49},
	84:  {^62, ^63},
	85:  {^30, ^31},
	86:  {^32, ^33},
	87:  {^40, ^41},
	88:  {^128, ^192},
	89:  {^26, ^27},
	90:  {^28, ^29},
	91:  {^34, ^35},
	92:  {^36, ^37},
	93:  {^38, ^39},
	94:  {^42, ^43},
	95:  {^640, ^704},
	96:  {^768, ^832},
	97:  {^1280, ^1344},
	98:  {^1408, ^1472},
	99:  {^1536, ^1600},
	100: {^1664, ^1728},
	101: {^512, ^576},
	102: {^896, ^960},
	103: {^1024, ^1088},
	104: {^1152, ^1216},
}

const maxCodeLength = 13

// Each encodeTable is represented by an array of bitStrings.

// bitString is a pair of uint32 values representing a bit code.
// The nBits low bits of bits make up the actual bit code.
// Eg. bitString{0x0004, 8} represents the bitcode "00000100".
type bitString struct {
	bits  uint32
	nBits uint32
}

// modeEncodeTable represents Table 1 and the End-of-Line code.
var modeEncodeTable = [...]bitString{
	0:  {0x0001, 4},  // "0001"
	1:  {0x0001, 3},  // "001"
	2:  {0x0001, 1},  // "1"
	3:  {0x0003, 3},  // "011"
	4:  {0x0003, 6},  // "000011"
	5:  {0x0003, 7},  // "0000011"
	6:  {0x0002, 3},  // "010"
	7:  {0x0002, 6},  // "000010"
	8:  {0x0002, 7},  // "0000010"
	9:  {0x0001, 7},  // "0000001"
	10: {0x0001, 12}, // "000000000001"
}

// whiteEncodeTable2 represents Table 2 for a white run.
var whiteEncodeTable2 = [...]bitString{
	0:  {0x0035, 8}, // "00110101"
	1:  {0x0007, 6}, // "000111"
	2:  {0x0007, 4}, // "0111"
	3:  {0x0008, 4}, // "1000"
	4:  {0x000b, 4}, // "1011"
	5:  {0x000c, 4}, // "1100"
	6:  {0x000e, 4}, // "1110"
	7:  {0x000f, 4}, // "1111"
	8:  {0x0013, 5}, // "10011"
	9:  {0x0014, 5}, // "10100"
	10: {0x0007, 5}, // "00111"
	11: {0x0008, 5}, // "01000"
	12: {0x0008, 6}, // "001000"
	13: {0x0003, 6}, // "000011"
	14: {0x0034, 6}, // "110100"
	15: {0x0035, 6}, // "110101"
	16: {0x002a, 6}, // "101010"
	17: {0x002b, 6}, // "101011"
	18: {0x0027, 7}, // "0100111"
	19: {0x000c, 7}, // "0001100"
	20: {0x0008, 7}, // "0001000"
	21: {0x0017, 7}, // "0010111"
	22: {0x0003, 7}, // "0000011"
	23: {0x0004, 7}, // "0000100"
	24: {0x0028, 7}, // "0101000"
	25: {0x002b, 7}, // "0101011"
	26: {0x0013, 7}, // "0010011"
	27: {0x0024, 7}, // "0100100"
	28: {0x0018, 7}, // "0011000"
	29: {0x0002, 8}, // "00000010"
	30: {0x0003, 8}, // "00000011"
	31: {0x001a, 8}, // "00011010"
	32: {0x001b, 8}, // "00011011"
	33: {0x0012, 8}, // "00010010"
	34: {0x0013, 8}, // "00010011"
	35: {0x0014, 8}, // "00010100"
	36: {0x0015, 8}, // "00010101"
	37: {0x0016, 8}, // "00010110"
	38: {0x0017, 8}, // "00010111"
	39: {0x0028, 8}, // "00101000"
	40: {0x0029, 8}, // "00101001"
	41: {0x002a, 8}, // "00101010"
	42: {0x002b, 8}, // "00101011"
	43: {0x002c, 8}, // "00101100"
	44: {0x002d, 8}, // "00101101"
	45: {0x0004, 8}, // "00000100"
	46: {0x0005, 8}, // "00000101"
	47: {0x000a, 8}, // "00001010"
	48: {0x000b, 8}, // "00001011"
	49: {0x0052, 8}, // "01010010"
	50: {0x0053, 8}, // "01010011"
	51: {0x0054, 8}, // "01010100"
	52: {0x0055, 8}, // "01010101"
	53: {0x0024, 8}, // "00100100"
	54: {0x0025, 8}, // "00100101"
	55: {0x0058, 8}, // "01011000"
	56: {0x0059, 8}, // "01011001"
	57: {0x005a, 8}, // "01011010"
	58: {0x005b, 8}, // "01011011"
	59: {0x004a, 8}, // "01001010"
	60: {0x004b, 8}, // "01001011"
	61: {0x0032, 8}, // "00110010"
	62: {0x0033, 8}, // "00110011"
	63: {0x0034, 8}, // "00110100"
}

// whiteEncodeTable3 represents Table 3 for a white run.
var whiteEncodeTable3 = [...]bitString{
	0:  {0x001b, 5},  // "11011"
	1:  {0x0012, 5},  // "10010"
	2:  {0x0017, 6},  // "010111"
	3:  {0x0037, 7},  // "0110111"
	4:  {0x0036, 8},  // "00110110"
	5:  {0x0037, 8},  // "00110111"
	6:  {0x0064, 8},  // "01100100"
	7:  {0x0065, 8},  // "01100101"
	8:  {0x0068, 8},  // "01101000"
	9:  {0x0067, 8},  // "01100111"
	10: {0x00cc, 9},  // "011001100"
	11: {0x00cd, 9},  // "011001101"
	12: {0x00d2, 9},  // "011010010"
	13: {0x00d3, 9},  // "011010011"
	14: {0x00d4, 9},  // "011010100"
	15: {0x00d5, 9},  // "011010101"
	16: {0x00d6, 9},  // "011010110"
	17: {0x00d7, 9},  // "011010111"
	18: {0x00d8, 9},  // "011011000"
	19: {0x00d9, 9},  // "011011001"
	20: {0x00da, 9},  // "011011010"
	21: {0x00db, 9},  // "011011011"
	22: {0x0098, 9},  // "010011000"
	23: {0x0099, 9},  // "010011001"
	24: {0x009a, 9},  // "010011010"
	25: {0x0018, 6},  // "011000"
	26: {0x009b, 9},  // "010011011"
	27: {0x0008, 11}, // "00000001000"
	28: {0x000c, 11}, // "00000001100"
	29: {0x000d, 11}, // "00000001101"
	30: {0x0012, 12}, // "000000010010"
	31: {0x0013, 12}, // "000000010011"
	32: {0x0014, 12}, // "000000010100"
	33: {0x0015, 12}, // "000000010101"
	34: {0x0016, 12}, // "000000010110"
	35: {0x0017, 12}, // "000000010111"
	36: {0x001c, 12}, // "000000011100"
	37: {0x001d, 12}, // "000000011101"
	38: {0x001e, 12}, // "000000011110"
	39: {0x001f, 12}, // "000000011111"
}

// blackEncodeTable2 represents Table 2 for a black run.
var blackEncodeTable2 = [...]bitString{
	0:  {0x0037, 10}, // "0000110111"
	1:  {0x0002, 3},  // "010"
	2:  {0x0003, 2},  // "11"
	3:  {0x0002, 2},  // "10"
	4:  {0x0003, 3},  // "011"
	5:  {0x0003, 4},  // "0011"
	6:  {0x0002, 4},  // "0010"
	7:  {0x0003, 5},  // "00011"
	8:  {0x0005, 6},  // "000101"
	9:  {0x0004, 6},  // "000100"
	10: {0x0004, 7},  // "0000100"
	11: {0x0005, 7},  // "0000101"
	12: {0x0007, 7},  // "0000111"
	13: {0x0004, 8},  // "00000100"
	14: {0x0007, 8},  // "00000111"
	15: {0x0018, 9},  // "000011000"
	16: {0x0017, 10}, // "0000010111"
	17: {0x0018, 10}, // "0000011000"
	18: {0x0008, 10}, // "0000001000"
	19: {0x0067, 11}, // "00001100111"
	20: {0x0068, 11}, // "00001101000"
	21: {0x006c, 11}, // "00001101100"
	22: {0x0037, 11}, // "00000110111"
	23: {0x0028, 11}, // "00000101000"
	24: {0x0017, 11}, // "00000010111"
	25: {0x0018, 11}, // "00000011000"
	26: {0x00ca, 12}, // "000011001010"
	27: {0x00cb, 12}, // "000011001011"
	28: {0x00cc, 12}, // "000011001100"
	29: {0x00cd, 12}, // "000011001101"
	30: {0x0068, 12}, // "000001101000"
	31: {0x0069, 12}, // "000001101001"
	32: {0x006a, 12}, // "000001101010"
	33: {0x006b, 12}, // "000001101011"
	34: {0x00d2, 12}, // "000011010010"
	35: {0x00d3, 12}, // "000011010011"
	36: {0x00d4, 12}, // "000011010100"
	37: {0x00d5, 12}, // "000011010101"
	38: {0x00d6, 12}, // "000011010110"
	39: {0x00d7, 12}, // "000011010111"
	40: {0x006c, 12}, // "000001101100"
	41: {0x006d, 12}, // "000001101101"
	42: {0x00da, 12}, // "000011011010"
	43: {0x00db, 12}, // "000011011011"
	44: {0x0054, 12}, // "000001010100"
	45: {0x0055, 12}, // "000001010101"
	46: {0x0056, 12}, // "000001010110"
	47: {0x0057, 12}, // "000001010111"
	48: {0x0064, 12}, // "000001100100"
	49: {0x0065, 12}, // "000001100101"
	50: {0x0052, 12}, // "000001010010"
	51: {0x0053, 12}, // "000001010011"
	52: {0x0024, 12}, // "000000100100"
	53: {0x0037, 12}, // "000000110111"
	54: {0x0038, 12}, // "000000111000"
	55: {0x0027, 12}, // "000000100111"
	56: {0x0028, 12}, // "000000101000"
	57: {0x0058, 12}, // "000001011000"
	58: {0x0059, 12}, // "000001011001"
	59: {0x002b, 12}, // "000000101011"
	60: {0x002c, 12}, // "000000101100"
	61: {0x005a, 12}, // "000001011010"
	62: {0x0066, 12}, // "000001100110"
	63: {0x0067, 12}, // "000001100111"
}

// blackEncodeTable3 represents Table 3 for a black run.
var blackEncodeTable3 = [...]bitString{
	0:  {0x000f, 10}, // "0000001111"
	1:  {0x00c8, 12}, // "000011001000"
	2:  {0x00c9, 12}, // "000011001001"
	3:  {0x005b, 12}, // "000001011011"
	4:  {0x0033, 12}, // "000000110011"
	5:  {0x0034, 12}, // "000000110100"
	6:  {0x0035, 12}, // "000000110101"
	7:  {0x006c, 13}, // "0000001101100"
	8:  {0x006d, 13}, // "0000001101101"
	9:  {0x004a, 13}, // "0000001001010"
	10: {0x004b, 13}, // "0000001001011"
	11: {0x004c, 13}, // "0000001001100"
	12: {0x004d, 13}, // "0000001001101"
	13: {0x0072, 13}, // "0000001110010"
	14: {0x0073, 13}, // "0000001110011"
	15: {0x0074, 13}, // "0000001110100"
	16: {0x0075, 13}, // "0000001110101"
	17: {0x0076, 13}, // "0000001110110"
	18: {0x0077, 13}, // "0000001110111"
	19: {0x0052, 13}, // "0000001010010"
	20: {0x0053, 13}, // "0000001010011"
	21: {0x0054, 13}, // "0000001010100"
	22: {0x0055, 13}, // "0000001010101"
	23: {0x005a, 13}, // "0000001011010"
	24: {0x005b, 13}, // "0000001011011"
	25: {0x0064, 13}, // "0000001100100"
	26: {0x0065, 13}, // "0000001100101"
	27: {0x0008, 11}, // "00000001000"
	28: {0x000c, 11}, // "00000001100"
	29: {0x000d, 11}, // "00000001101"
	30: {0x0012, 12}, // "000000010010"
	31: {0x0013, 12}, // "000000010011"
	32: {0x0014, 12}, // "000000010100"
	33: {0x0015, 12}, // "000000010101"
	34: {0x0016, 12}, // "000000010110"
	35: {0x0017, 12}, // "000000010111"
	36: {0x001c, 12}, // "000000011100"
	37: {0x001d, 12}, // "000000011101"
	38: {0x001e, 12}, // "000000011110"
	39: {0x001f, 12}, // "000000011111"
}

// COPY PASTE table.go BEGIN

const (
	modePass = iota // Pass
	modeH           // Horizontal
	modeV0          // Vertical-0
	modeVR1         // Vertical-Right-1
	modeVR2         // Vertical-Right-2
	modeVR3         // Vertical-Right-3
	modeVL1         // Vertical-Left-1
	modeVL2         // Vertical-Left-2
	modeVL3         // Vertical-Left-3
	modeExt         // Extension
	modeEOL         // End-of-Line
)

// COPY PASTE table.go END
