package errors

type GMOPGError interface {
	Error() string
	Message() string
	Code() string
	CanRetry() bool
}

// NewErrors converts GMO-PG errors to Struct
// Examples:
//		["E61010001", "E61030001"]
func NewErrors(errInfos ...string) []GMOPGError {
	errs := make([]GMOPGError, 0, len(errInfos))
	for _, info := range errInfos {
		switch info {
		case "E00000000":
			errs = append(errs, &PG_E00000000{})
		case "E00000001":
			errs = append(errs, &PG_E00000001{})
		case "E00000002":
			errs = append(errs, &PG_E00000002{})
		case "E00000003":
			errs = append(errs, &PG_E00000003{})
		case "E00000010":
			errs = append(errs, &PG_E00000010{})
		case "E01010001":
			errs = append(errs, &PG_E01010001{})
		case "E01010008":
			errs = append(errs, &PG_E01010008{})
		case "E01010010":
			errs = append(errs, &PG_E01010010{})
		case "E01020001":
			errs = append(errs, &PG_E01020001{})
		case "E01020008":
			errs = append(errs, &PG_E01020008{})
		case "E01030002":
			errs = append(errs, &PG_E01030002{})
		case "E01030061":
			errs = append(errs, &PG_E01030061{})
		case "E01040001":
			errs = append(errs, &PG_E01040001{})
		case "E01040003":
			errs = append(errs, &PG_E01040003{})
		case "E01040010":
			errs = append(errs, &PG_E01040010{})
		case "E01040013":
			errs = append(errs, &PG_E01040013{})
		case "E01050001":
			errs = append(errs, &PG_E01050001{})
		case "E01050002":
			errs = append(errs, &PG_E01050002{})
		case "E01050004":
			errs = append(errs, &PG_E01050004{})
		case "E01060001":
			errs = append(errs, &PG_E01060001{})
		case "E01060005":
			errs = append(errs, &PG_E01060005{})
		case "E01060006":
			errs = append(errs, &PG_E01060006{})
		case "E01060010":
			errs = append(errs, &PG_E01060010{})
		case "E01060011":
			errs = append(errs, &PG_E01060011{})
		case "E01060021":
			errs = append(errs, &PG_E01060021{})
		case "E01070005":
			errs = append(errs, &PG_E01070005{})
		case "E01070006":
			errs = append(errs, &PG_E01070006{})
		case "E01080007":
			errs = append(errs, &PG_E01080007{})
		case "E01080010":
			errs = append(errs, &PG_E01080010{})
		case "E01080101":
			errs = append(errs, &PG_E01080101{})
		case "E01090001":
			errs = append(errs, &PG_E01090001{})
		case "E01090008":
			errs = append(errs, &PG_E01090008{})
		case "E01100001":
			errs = append(errs, &PG_E01100001{})
		case "E01100008":
			errs = append(errs, &PG_E01100008{})
		case "E01110002":
			errs = append(errs, &PG_E01110002{})
		case "E01110010":
			errs = append(errs, &PG_E01110010{})
		case "E01130012":
			errs = append(errs, &PG_E01130012{})
		case "E01160001":
			errs = append(errs, &PG_E01160001{})
		case "E01160007":
			errs = append(errs, &PG_E01160007{})
		case "E01160010":
			errs = append(errs, &PG_E01160010{})
		case "E01170001":
			errs = append(errs, &PG_E01170001{})
		case "E01170003":
			errs = append(errs, &PG_E01170003{})
		case "E01170006":
			errs = append(errs, &PG_E01170006{})
		case "E01170011":
			errs = append(errs, &PG_E01170011{})
		case "E01180001":
			errs = append(errs, &PG_E01180001{})
		case "E01180003":
			errs = append(errs, &PG_E01180003{})
		case "E01180006":
			errs = append(errs, &PG_E01180006{})
		case "E01180008":
			errs = append(errs, &PG_E01180008{})
		case "E01180011":
			errs = append(errs, &PG_E01180011{})
		case "E01190001":
			errs = append(errs, &PG_E01190001{})
		case "E01190008":
			errs = append(errs, &PG_E01190008{})
		case "E01200001":
			errs = append(errs, &PG_E01200001{})
		case "E01200007":
			errs = append(errs, &PG_E01200007{})
		case "E01200008":
			errs = append(errs, &PG_E01200008{})
		case "E01210002":
			errs = append(errs, &PG_E01210002{})
		case "E01220001":
			errs = append(errs, &PG_E01220001{})
		case "E01220005":
			errs = append(errs, &PG_E01220005{})
		case "E01220008":
			errs = append(errs, &PG_E01220008{})
		case "E01220010":
			errs = append(errs, &PG_E01220010{})
		case "E01220012":
			errs = append(errs, &PG_E01220012{})
		case "E01230006":
			errs = append(errs, &PG_E01230006{})
		case "E01230009":
			errs = append(errs, &PG_E01230009{})
		case "E01240002":
			errs = append(errs, &PG_E01240002{})
		case "E01240012":
			errs = append(errs, &PG_E01240012{})
		case "E01250008":
			errs = append(errs, &PG_E01250008{})
		case "E01250010":
			errs = append(errs, &PG_E01250010{})
		case "E01260001":
			errs = append(errs, &PG_E01260001{})
		case "E01260002":
			errs = append(errs, &PG_E01260002{})
		case "E01260010":
			errs = append(errs, &PG_E01260010{})
		case "E01270001":
			errs = append(errs, &PG_E01270001{})
		case "E01270005":
			errs = append(errs, &PG_E01270005{})
		case "E01270006":
			errs = append(errs, &PG_E01270006{})
		case "E01270010":
			errs = append(errs, &PG_E01270010{})
		case "E01290001":
			errs = append(errs, &PG_E01290001{})
		case "E01300001":
			errs = append(errs, &PG_E01300001{})
		case "E01310001":
			errs = append(errs, &PG_E01310001{})
		case "E01320012":
			errs = append(errs, &PG_E01320012{})
		case "E01320013":
			errs = append(errs, &PG_E01320013{})
		case "E01330012":
			errs = append(errs, &PG_E01330012{})
		case "E01330013":
			errs = append(errs, &PG_E01330013{})
		case "E01340012":
			errs = append(errs, &PG_E01340012{})
		case "E01340013":
			errs = append(errs, &PG_E01340013{})
		case "E01350001":
			errs = append(errs, &PG_E01350001{})
		case "E01350008":
			errs = append(errs, &PG_E01350008{})
		case "E01360001":
			errs = append(errs, &PG_E01360001{})
		case "E01370008":
			errs = append(errs, &PG_E01370008{})
		case "E01370012":
			errs = append(errs, &PG_E01370012{})
		case "E01390002":
			errs = append(errs, &PG_E01390002{})
		case "E01390010":
			errs = append(errs, &PG_E01390010{})
		case "E01400007":
			errs = append(errs, &PG_E01400007{})
		case "E01410010":
			errs = append(errs, &PG_E01410010{})
		case "E01420010":
			errs = append(errs, &PG_E01420010{})
		case "E01430012":
			errs = append(errs, &PG_E01430012{})
		case "E01440008":
			errs = append(errs, &PG_E01440008{})
		case "E01450008":
			errs = append(errs, &PG_E01450008{})
		case "E01460008":
			errs = append(errs, &PG_E01460008{})
		case "E01470008":
			errs = append(errs, &PG_E01470008{})
		case "E01480008":
			errs = append(errs, &PG_E01480008{})
		case "E01480011":
			errs = append(errs, &PG_E01480011{})
		case "E01490005":
			errs = append(errs, &PG_E01490005{})
		case "E01500001":
			errs = append(errs, &PG_E01500001{})
		case "E01500005":
			errs = append(errs, &PG_E01500005{})
		case "E01500012":
			errs = append(errs, &PG_E01500012{})
		case "E01510001":
			errs = append(errs, &PG_E01510001{})
		case "E01510005":
			errs = append(errs, &PG_E01510005{})
		case "E01510010":
			errs = append(errs, &PG_E01510010{})
		case "E01510011":
			errs = append(errs, &PG_E01510011{})
		case "E01510012":
			errs = append(errs, &PG_E01510012{})
		case "E01520002":
			errs = append(errs, &PG_E01520002{})
		case "E01530001":
			errs = append(errs, &PG_E01530001{})
		case "E01530005":
			errs = append(errs, &PG_E01530005{})
		case "E01540005":
			errs = append(errs, &PG_E01540005{})
		case "E01550001":
			errs = append(errs, &PG_E01550001{})
		case "E01550005":
			errs = append(errs, &PG_E01550005{})
		case "E01550006":
			errs = append(errs, &PG_E01550006{})
		case "E01590005":
			errs = append(errs, &PG_E01590005{})
		case "E01590006":
			errs = append(errs, &PG_E01590006{})
		case "E01600001":
			errs = append(errs, &PG_E01600001{})
		case "E01600005":
			errs = append(errs, &PG_E01600005{})
		case "E01600012":
			errs = append(errs, &PG_E01600012{})
		case "E01610005":
			errs = append(errs, &PG_E01610005{})
		case "E01610006":
			errs = append(errs, &PG_E01610006{})
		case "E01620005":
			errs = append(errs, &PG_E01620005{})
		case "E01620006":
			errs = append(errs, &PG_E01620006{})
		case "E01630010":
			errs = append(errs, &PG_E01630010{})
		case "E01640010":
			errs = append(errs, &PG_E01640010{})
		case "E01650012":
			errs = append(errs, &PG_E01650012{})
		case "E01660013":
			errs = append(errs, &PG_E01660013{})
		case "E01670013":
			errs = append(errs, &PG_E01670013{})
		case "E01700001":
			errs = append(errs, &PG_E01700001{})
		case "E01710001":
			errs = append(errs, &PG_E01710001{})
		case "E01710002":
			errs = append(errs, &PG_E01710002{})
		case "E01730001":
			errs = append(errs, &PG_E01730001{})
		case "E01730005":
			errs = append(errs, &PG_E01730005{})
		case "E01730006":
			errs = append(errs, &PG_E01730006{})
		case "E01730007":
			errs = append(errs, &PG_E01730007{})
		case "E01740001":
			errs = append(errs, &PG_E01740001{})
		case "E01740005":
			errs = append(errs, &PG_E01740005{})
		case "E01740007":
			errs = append(errs, &PG_E01740007{})
		case "E01750001":
			errs = append(errs, &PG_E01750001{})
		case "E01750008":
			errs = append(errs, &PG_E01750008{})
		case "E01770002":
			errs = append(errs, &PG_E01770002{})
		case "E01780002":
			errs = append(errs, &PG_E01780002{})
		case "E01790007":
			errs = append(errs, &PG_E01790007{})
		case "E01790011":
			errs = append(errs, &PG_E01790011{})
		case "E01800001":
			errs = append(errs, &PG_E01800001{})
		case "E01800008":
			errs = append(errs, &PG_E01800008{})
		case "E01800010":
			errs = append(errs, &PG_E01800010{})
		case "E01800050":
			errs = append(errs, &PG_E01800050{})
		case "E01810001":
			errs = append(errs, &PG_E01810001{})
		case "E01810008":
			errs = append(errs, &PG_E01810008{})
		case "E01820001":
			errs = append(errs, &PG_E01820001{})
		case "E01820003":
			errs = append(errs, &PG_E01820003{})
		case "E01820008":
			errs = append(errs, &PG_E01820008{})
		case "E01840010":
			errs = append(errs, &PG_E01840010{})
		case "E01850008":
			errs = append(errs, &PG_E01850008{})
		case "E01860008":
			errs = append(errs, &PG_E01860008{})
		case "E01870008":
			errs = append(errs, &PG_E01870008{})
		case "E01880001":
			errs = append(errs, &PG_E01880001{})
		case "E01880002":
			errs = append(errs, &PG_E01880002{})
		case "E01880008":
			errs = append(errs, &PG_E01880008{})
		case "E01890001":
			errs = append(errs, &PG_E01890001{})
		case "E01890002":
			errs = append(errs, &PG_E01890002{})
		case "E01890006":
			errs = append(errs, &PG_E01890006{})
		case "E01890009":
			errs = append(errs, &PG_E01890009{})
		case "E01910008":
			errs = append(errs, &PG_E01910008{})
		case "E01950008":
			errs = append(errs, &PG_E01950008{})
		case "E01960008":
			errs = append(errs, &PG_E01960008{})
		case "E01970008":
			errs = append(errs, &PG_E01970008{})
		case "E01980008":
			errs = append(errs, &PG_E01980008{})
		case "E01990005":
			errs = append(errs, &PG_E01990005{})
		case "E01990006":
			errs = append(errs, &PG_E01990006{})
		case "E01999998":
			errs = append(errs, &PG_E01999998{})
		case "E01A00008":
			errs = append(errs, &PG_E01A00008{})
		case "E01A10005":
			errs = append(errs, &PG_E01A10005{})
		case "E01A10006":
			errs = append(errs, &PG_E01A10006{})
		case "E01A20008":
			errs = append(errs, &PG_E01A20008{})
		case "E01A30008":
			errs = append(errs, &PG_E01A30008{})
		case "E01A40008":
			errs = append(errs, &PG_E01A40008{})
		case "E01A50005":
			errs = append(errs, &PG_E01A50005{})
		case "E01A50006":
			errs = append(errs, &PG_E01A50006{})
		case "E01A60005":
			errs = append(errs, &PG_E01A60005{})
		case "E01A60006":
			errs = append(errs, &PG_E01A60006{})
		case "E01A70012":
			errs = append(errs, &PG_E01A70012{})
		case "E01A80008":
			errs = append(errs, &PG_E01A80008{})
		case "E01A90008":
			errs = append(errs, &PG_E01A90008{})
		case "E01B00008":
			errs = append(errs, &PG_E01B00008{})
		case "E01B10005":
			errs = append(errs, &PG_E01B10005{})
		case "E01B20002":
			errs = append(errs, &PG_E01B20002{})
		case "E01B20005":
			errs = append(errs, &PG_E01B20005{})
		case "E01B30005":
			errs = append(errs, &PG_E01B30005{})
		case "E01B40005":
			errs = append(errs, &PG_E01B40005{})
		case "E01B50005":
			errs = append(errs, &PG_E01B50005{})
		case "E01B60005":
			errs = append(errs, &PG_E01B60005{})
		case "E01B70005":
			errs = append(errs, &PG_E01B70005{})
		case "E01B70010":
			errs = append(errs, &PG_E01B70010{})
		case "E01B80005":
			errs = append(errs, &PG_E01B80005{})
		case "E01B80008":
			errs = append(errs, &PG_E01B80008{})
		case "E01B90005":
			errs = append(errs, &PG_E01B90005{})
		case "E01C00005":
			errs = append(errs, &PG_E01C00005{})
		case "E01C00006":
			errs = append(errs, &PG_E01C00006{})
		case "E01C10005":
			errs = append(errs, &PG_E01C10005{})
		case "E01C20005":
			errs = append(errs, &PG_E01C20005{})
		case "E01C20006":
			errs = append(errs, &PG_E01C20006{})
		case "E01C30005":
			errs = append(errs, &PG_E01C30005{})
		case "E01C40005":
			errs = append(errs, &PG_E01C40005{})
		case "E01C40006":
			errs = append(errs, &PG_E01C40006{})
		case "E01C50005":
			errs = append(errs, &PG_E01C50005{})
		case "E01C60002":
			errs = append(errs, &PG_E01C60002{})
		case "E01C60005":
			errs = append(errs, &PG_E01C60005{})
		case "E01C70005":
			errs = append(errs, &PG_E01C70005{})
		case "E01C80005":
			errs = append(errs, &PG_E01C80005{})
		case "E01C90005":
			errs = append(errs, &PG_E01C90005{})
		case "E01D00005":
			errs = append(errs, &PG_E01D00005{})
		case "E01D10005":
			errs = append(errs, &PG_E01D10005{})
		case "E01D10010":
			errs = append(errs, &PG_E01D10010{})
		case "E01D20005":
			errs = append(errs, &PG_E01D20005{})
		case "E01D20008":
			errs = append(errs, &PG_E01D20008{})
		case "E01D30008":
			errs = append(errs, &PG_E01D30008{})
		case "E01D40005":
			errs = append(errs, &PG_E01D40005{})
		case "E01D40006":
			errs = append(errs, &PG_E01D40006{})
		case "E01D50005":
			errs = append(errs, &PG_E01D50005{})
		case "E01D50006":
			errs = append(errs, &PG_E01D50006{})
		case "E01D60005":
			errs = append(errs, &PG_E01D60005{})
		case "E01D70008":
			errs = append(errs, &PG_E01D70008{})
		case "E01D80008":
			errs = append(errs, &PG_E01D80008{})
		case "E01D90008":
			errs = append(errs, &PG_E01D90008{})
		case "E01E00008":
			errs = append(errs, &PG_E01E00008{})
		case "E01E10008":
			errs = append(errs, &PG_E01E10008{})
		case "E01E20005":
			errs = append(errs, &PG_E01E20005{})
		case "E01E20006":
			errs = append(errs, &PG_E01E20006{})
		case "E01E30001":
			errs = append(errs, &PG_E01E30001{})
		case "E01E30005":
			errs = append(errs, &PG_E01E30005{})
		case "E01E40001":
			errs = append(errs, &PG_E01E40001{})
		case "E01E40008":
			errs = append(errs, &PG_E01E40008{})
		case "E01E60001":
			errs = append(errs, &PG_E01E60001{})
		case "E01E70001":
			errs = append(errs, &PG_E01E70001{})
		case "E01E80001":
			errs = append(errs, &PG_E01E80001{})
		case "E01E90001":
			errs = append(errs, &PG_E01E90001{})
		case "E11010001":
			errs = append(errs, &PG_E11010001{})
		case "E11010002":
			errs = append(errs, &PG_E11010002{})
		case "E11010003":
			errs = append(errs, &PG_E11010003{})
		case "E11010010":
			errs = append(errs, &PG_E11010010{})
		case "E11010011":
			errs = append(errs, &PG_E11010011{})
		case "E11010012":
			errs = append(errs, &PG_E11010012{})
		case "E11010013":
			errs = append(errs, &PG_E11010013{})
		case "E11010014":
			errs = append(errs, &PG_E11010014{})
		case "E11010015":
			errs = append(errs, &PG_E11010015{})
		case "E11010016":
			errs = append(errs, &PG_E11010016{})
		case "E11010017":
			errs = append(errs, &PG_E11010017{})
		case "E11010099":
			errs = append(errs, &PG_E11010099{})
		case "E11010999":
			errs = append(errs, &PG_E11010999{})
		case "E11310001":
			errs = append(errs, &PG_E11310001{})
		case "E11310002":
			errs = append(errs, &PG_E11310002{})
		case "E11310003":
			errs = append(errs, &PG_E11310003{})
		case "E11310004":
			errs = append(errs, &PG_E11310004{})
		case "E11310005":
			errs = append(errs, &PG_E11310005{})
		case "E21010001":
			errs = append(errs, &PG_E21010001{})
		case "E21010007":
			errs = append(errs, &PG_E21010007{})
		case "E21010999":
			errs = append(errs, &PG_E21010999{})
		case "E21020001":
			errs = append(errs, &PG_E21020001{})
		case "E21020002":
			errs = append(errs, &PG_E21020002{})
		case "E21020007":
			errs = append(errs, &PG_E21020007{})
		case "E21020999":
			errs = append(errs, &PG_E21020999{})
		case "E21010201":
			errs = append(errs, &PG_E21010201{})
		case "E21010202":
			errs = append(errs, &PG_E21010202{})
		case "E21030001":
			errs = append(errs, &PG_E21030001{})
		case "E21030007":
			errs = append(errs, &PG_E21030007{})
		case "E21030201":
			errs = append(errs, &PG_E21030201{})
		case "E21030202":
			errs = append(errs, &PG_E21030202{})
		case "E31500014":
			errs = append(errs, &PG_E31500014{})
		case "E41170002":
			errs = append(errs, &PG_E41170002{})
		case "E41170099":
			errs = append(errs, &PG_E41170099{})
		case "E61010001":
			errs = append(errs, &PG_E61010001{})
		case "E61010002":
			errs = append(errs, &PG_E61010002{})
		case "E61010003":
			errs = append(errs, &PG_E61010003{})
		case "E61020001":
			errs = append(errs, &PG_E61020001{})
		case "E61030001":
			errs = append(errs, &PG_E61030001{})
		case "E61040001":
			errs = append(errs, &PG_E61040001{})
		case "E82010001":
			errs = append(errs, &PG_E82010001{})
		case "E90010001":
			errs = append(errs, &PG_E90010001{})
		case "E91099996":
			errs = append(errs, &PG_E91099996{})
		case "E91099997":
			errs = append(errs, &PG_E91099997{})
		case "E91019999":
			errs = append(errs, &PG_E91019999{})
		case "E91020001":
			errs = append(errs, &PG_E91020001{})
		case "E91029998":
			errs = append(errs, &PG_E91029998{})
		case "E91029999":
			errs = append(errs, &PG_E91029999{})
		case "E91050001":
			errs = append(errs, &PG_E91050001{})
		case "E91060001":
			errs = append(errs, &PG_E91060001{})
		case "E91099999":
			errs = append(errs, &PG_E91099999{})
		case "E92000001":
			errs = append(errs, &PG_E92000001{})
		case "E92000002":
			errs = append(errs, &PG_E92000002{})
		case "EX1000301":
			errs = append(errs, &PG_EX1000301{})
		case "EX1000302":
			errs = append(errs, &PG_EX1000302{})
		case "EX1000303":
			errs = append(errs, &PG_EX1000303{})
		case "EX1000304":
			errs = append(errs, &PG_EX1000304{})
		case "42C010000":
			errs = append(errs, &PG_42C010000{})
		case "42C030000":
			errs = append(errs, &PG_42C030000{})
		case "42C120000":
			errs = append(errs, &PG_42C120000{})
		case "42C130000":
			errs = append(errs, &PG_42C130000{})
		case "42C140000":
			errs = append(errs, &PG_42C140000{})
		case "42C150000":
			errs = append(errs, &PG_42C150000{})
		case "42C330000":
			errs = append(errs, &PG_42C330000{})
		case "42C500000":
			errs = append(errs, &PG_42C500000{})
		case "42C510000":
			errs = append(errs, &PG_42C510000{})
		case "42C530000":
			errs = append(errs, &PG_42C530000{})
		case "42C540000":
			errs = append(errs, &PG_42C540000{})
		case "42C550000":
			errs = append(errs, &PG_42C550000{})
		case "42C560000":
			errs = append(errs, &PG_42C560000{})
		case "42C570000":
			errs = append(errs, &PG_42C570000{})
		case "42C580000":
			errs = append(errs, &PG_42C580000{})
		case "42C600000":
			errs = append(errs, &PG_42C600000{})
		case "42C700000":
			errs = append(errs, &PG_42C700000{})
		case "42C710000":
			errs = append(errs, &PG_42C710000{})
		case "42C720000":
			errs = append(errs, &PG_42C720000{})
		case "42C730000":
			errs = append(errs, &PG_42C730000{})
		case "42C740000":
			errs = append(errs, &PG_42C740000{})
		case "42C750000":
			errs = append(errs, &PG_42C750000{})
		case "42C760000":
			errs = append(errs, &PG_42C760000{})
		case "42C770000":
			errs = append(errs, &PG_42C770000{})
		case "42C780000":
			errs = append(errs, &PG_42C780000{})
		case "42G020000":
			errs = append(errs, &PG_42G020000{})
		case "42G030000":
			errs = append(errs, &PG_42G030000{})
		case "42G040000":
			errs = append(errs, &PG_42G040000{})
		case "42G050000":
			errs = append(errs, &PG_42G050000{})
		case "42G060000":
			errs = append(errs, &PG_42G060000{})
		case "42G070000":
			errs = append(errs, &PG_42G070000{})
		case "42G120000":
			errs = append(errs, &PG_42G120000{})
		case "42G220000":
			errs = append(errs, &PG_42G220000{})
		case "42G300000":
			errs = append(errs, &PG_42G300000{})
		case "42G420000":
			errs = append(errs, &PG_42G420000{})
		case "42G440000":
			errs = append(errs, &PG_42G440000{})
		case "42G450000":
			errs = append(errs, &PG_42G450000{})
		case "42G540000":
			errs = append(errs, &PG_42G540000{})
		case "42G550000":
			errs = append(errs, &PG_42G550000{})
		case "42G560000":
			errs = append(errs, &PG_42G560000{})
		case "42G600000":
			errs = append(errs, &PG_42G600000{})
		case "42G610000":
			errs = append(errs, &PG_42G610000{})
		case "42G650000":
			errs = append(errs, &PG_42G650000{})
		case "42G670000":
			errs = append(errs, &PG_42G670000{})
		case "42G680000":
			errs = append(errs, &PG_42G680000{})
		case "42G690000":
			errs = append(errs, &PG_42G690000{})
		case "42G700000":
			errs = append(errs, &PG_42G700000{})
		case "42G710000":
			errs = append(errs, &PG_42G710000{})
		case "42G720000":
			errs = append(errs, &PG_42G720000{})
		case "42G730000":
			errs = append(errs, &PG_42G730000{})
		case "42G740000":
			errs = append(errs, &PG_42G740000{})
		case "42G750000":
			errs = append(errs, &PG_42G750000{})
		case "42G760000":
			errs = append(errs, &PG_42G760000{})
		case "42G770000":
			errs = append(errs, &PG_42G770000{})
		case "42G780000":
			errs = append(errs, &PG_42G780000{})
		case "42G790000":
			errs = append(errs, &PG_42G790000{})
		case "42G800000":
			errs = append(errs, &PG_42G800000{})
		case "42G810000":
			errs = append(errs, &PG_42G810000{})
		case "42G830000":
			errs = append(errs, &PG_42G830000{})
		case "42G910000":
			errs = append(errs, &PG_42G910000{})
		case "42G920000":
			errs = append(errs, &PG_42G920000{})
		case "42G950000":
			errs = append(errs, &PG_42G950000{})
		case "42G960000":
			errs = append(errs, &PG_42G960000{})
		case "42G970000":
			errs = append(errs, &PG_42G970000{})
		case "42G980000":
			errs = append(errs, &PG_42G980000{})
		case "42G990000":
			errs = append(errs, &PG_42G990000{})

		}
	}
	return errs
}
