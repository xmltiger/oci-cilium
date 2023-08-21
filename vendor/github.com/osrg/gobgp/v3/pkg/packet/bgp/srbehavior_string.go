// Code generated by "stringer -type=SRBehavior"; DO NOT EDIT.

package bgp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RESERVED-0]
	_ = x[END-1]
	_ = x[END_WITH_PSP-2]
	_ = x[END_WITH_USP-3]
	_ = x[END_WITH_PSP_USP-4]
	_ = x[ENDX-5]
	_ = x[ENDX_WITH_PSP-6]
	_ = x[ENDX_WITH_USP-7]
	_ = x[ENDX_WITH_PSP_USP-8]
	_ = x[ENDT-9]
	_ = x[ENDT_WITH_PSP-10]
	_ = x[ENDT_WITH_USP-11]
	_ = x[ENDT_WITH_PSP_USP-12]
	_ = x[END_B6_ENCAPS-14]
	_ = x[END_BM-15]
	_ = x[END_DX6-16]
	_ = x[END_DX4-17]
	_ = x[END_DT6-18]
	_ = x[END_DT4-19]
	_ = x[END_DT46-20]
	_ = x[END_DX2-21]
	_ = x[END_DX2V-22]
	_ = x[END_DT2U-23]
	_ = x[END_DT2M-24]
	_ = x[END_B6_ENCAPS_Red-27]
	_ = x[END_WITH_USD-28]
	_ = x[END_WITH_PSP_USD-29]
	_ = x[END_WITH_USP_USD-30]
	_ = x[END_WITH_PSP_USP_USD-31]
	_ = x[ENDX_WITH_USD-32]
	_ = x[ENDX_WITH_PSP_USD-33]
	_ = x[ENDX_WITH_USP_USD-34]
	_ = x[ENDX_WITH_PSP_USP_USD-35]
	_ = x[ENDT_WITH_USD-36]
	_ = x[ENDT_WITH_PSP_USD-37]
	_ = x[ENDT_WITH_USP_USD-38]
	_ = x[ENDT_WITH_PSP_USP_USD-39]
	_ = x[ENDM_GTP6D-69]
	_ = x[ENDM_GTP6DI-70]
	_ = x[ENDM_GTP6E-71]
	_ = x[ENDM_GTP4E-72]
}

const (
	_SRBehavior_name_0 = "RESERVEDENDEND_WITH_PSPEND_WITH_USPEND_WITH_PSP_USPENDXENDX_WITH_PSPENDX_WITH_USPENDX_WITH_PSP_USPENDTENDT_WITH_PSPENDT_WITH_USPENDT_WITH_PSP_USP"
	_SRBehavior_name_1 = "END_B6_ENCAPSEND_BMEND_DX6END_DX4END_DT6END_DT4END_DT46END_DX2END_DX2VEND_DT2UEND_DT2M"
	_SRBehavior_name_2 = "END_B6_ENCAPS_RedEND_WITH_USDEND_WITH_PSP_USDEND_WITH_USP_USDEND_WITH_PSP_USP_USDENDX_WITH_USDENDX_WITH_PSP_USDENDX_WITH_USP_USDENDX_WITH_PSP_USP_USDENDT_WITH_USDENDT_WITH_PSP_USDENDT_WITH_USP_USDENDT_WITH_PSP_USP_USD"
	_SRBehavior_name_3 = "ENDM_GTP6DENDM_GTP6DIENDM_GTP6EENDM_GTP4E"
)

var (
	_SRBehavior_index_0 = [...]uint8{0, 8, 11, 23, 35, 51, 55, 68, 81, 98, 102, 115, 128, 145}
	_SRBehavior_index_1 = [...]uint8{0, 13, 19, 26, 33, 40, 47, 55, 62, 70, 78, 86}
	_SRBehavior_index_2 = [...]uint8{0, 17, 29, 45, 61, 81, 94, 111, 128, 149, 162, 179, 196, 217}
	_SRBehavior_index_3 = [...]uint8{0, 10, 21, 31, 41}
)

func (i SRBehavior) String() string {
	switch {
	case 0 <= i && i <= 12:
		return _SRBehavior_name_0[_SRBehavior_index_0[i]:_SRBehavior_index_0[i+1]]
	case 14 <= i && i <= 24:
		i -= 14
		return _SRBehavior_name_1[_SRBehavior_index_1[i]:_SRBehavior_index_1[i+1]]
	case 27 <= i && i <= 39:
		i -= 27
		return _SRBehavior_name_2[_SRBehavior_index_2[i]:_SRBehavior_index_2[i+1]]
	case 69 <= i && i <= 72:
		i -= 69
		return _SRBehavior_name_3[_SRBehavior_index_3[i]:_SRBehavior_index_3[i+1]]
	default:
		return "SRBehavior(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}