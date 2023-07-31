package model

type ObjectProcess struct {
	TableName string
	EndPoint  string
	DBName    string
	Value     []TargetObject
}

type SourceObject struct {
	QCD    string
	TIME   string
	TKQKBN string
	SNDC   string
	ZXD    string
	NOPV   string
	ZSS    string
	DPS    string
	DPS_T  string
	DPP    string
	DPP_T  string
	DPR    string
	DPR_T  string
	XV     string
	XV_T   string
	NOQ    string
	QAS    string
	QAP    string
	QAP_T  string
	QAR    string
	QAR_T  string
	AV     string
	AV_T   string
	QBS    string
	QBP    string
	QBP_T  string
	QBR    string
	QBR_T  string
	BV     string
	BV_T   string
	J_SNDC string
	XJ     string
	XJ_T   string
	DPCF   string
	JQCS   string
	JQJS   string
	DPCY   string
	DPCY_T string
	TLNM   string
	XV_S   string
	AV_S   string
	BV_S   string
	DPP_S  string
	J_ZXD  string
	MIDP   string
	MIDP_T string
	DYWP   string
	DYWP_T string
	DYWR   string
	VWAP   string
	VWAP_T string
	ABV    string
	ABV_T  string
	AAV    string
	AAV_T  string
	QOV    string
	QOV_T  string
	QUV    string
	QUV_T  string
	INAV   string
	INAV_T string
	IYRP   string
	IYRP_T string
	IQRP   string
	IQRP_T string
	QACY   string
	QACY_T string
	QBCY   string
	QBCY_T string
	TSTF   string
	NO     string
	NOT    string
	K1KF   string
	K2KF   string
	DKF    string
	DSYRP  string
	DSYWP  string
}

type TargetObject struct {
	QCD            string
	TIME           string
	TKZXD          string
	TKTIM          string
	TKSERIALNUMBER int
	TKQKBN         string
	SNDC           string
	ZXD            string
	NOPV           string
	ZSS            string
	DPS            string
	DPS_T          string
	DPP            string
	DPP_T          string
	DPR            string
	DPR_T          string
	XV             string
	XV_T           string
	NOQ            string
	QAS            string
	QAP            string
	QAP_T          string
	QAR            string
	QAR_T          string
	AV             string
	AV_T           string
	QBS            string
	QBP            string
	QBP_T          string
	QBR            string
	QBR_T          string
	BV             string
	BV_T           string
	J_SNDC         string
	XJ             string
	XJ_T           string
	DPCF           string
	JQCS           string
	JQJS           string
	DPCY           string
	DPCY_T         string
	TLNM           string
	XV_S           string
	AV_S           string
	BV_S           string
	DPP_S          string
	J_ZXD          string
	MIDP           string
	MIDP_T         string
	DYWP           string
	DYWP_T         string
	DYWR           string
	VWAP           string
	VWAP_T         string
	ABV            string
	ABV_T          string
	AAV            string
	AAV_T          string
	QOV            string
	QOV_T          string
	QUV            string
	QUV_T          string
	INAV           string
	INAV_T         string
	IYRP           string
	IYRP_T         string
	IQRP           string
	IQRP_T         string
	QACY           string
	QACY_T         string
	QBCY           string
	QBCY_T         string
	TSTF           string
	NO             string
	NOT            string
	K1KF           string
	K2KF           string
	DKF            string
	DSYRP          string
	DSYWP          string
}
