package hcv1a

import (
	ap "github.com/hivdb/nucamino/alignmentprofile"
	a "github.com/hivdb/nucamino/types/amino"
)

var hcv1aPositionalIndelScores = ap.GenePositionalIndelScores{}

// HCV1a Reference Sequences; Genebank accession ID: NC_004102.1

var (
	HCV1ASEQ_NS3 = a.ReadString(`
		APITAYAQQTRGLLGCIITSLTGRDKNQVEGEVQIVSTATQTFLATCINGVCWTVYHGAGTRTIASPKGP
		VIQMYTNVDQDLVGWPAPQGSRSLTPCTCGSSDLYLVTRHADVIPVRRRGDSRGSLLSPRPISYLKGSSG
		GPLLCPAGHAVGLFRAAVCTRGVAKAVDFIPVENLETTMRSPVFTDNSSPPAVPQSFQVAHLHAPTGSGK
		STKVPAAYAAQGYKVLVLNPSVAATLGFGAYMSKAHGVDPNIRTGVRTITTGSPITYSTYGKFLADGGCS
		GGAYDIIICDECHSTDATSILGIGTVLDQAETAGARLVVLATATPPGSVTVSHPNIEEVALSTTGEIPFY
		GKAIPLEVIKGGRHLIFCHSKKKCDELAAKLVALGINAVAYYRGLDVSVIPTSGDVVVVSTDALMTGFTG
		DFDSVIDCNTCVTQTVDFSLDPTFTIETTTLPQDAVSRTQRRGRTGRGKPGIYRFVAPGERPSGMFDSSV
		LCECYDAGCAWYELTPAETTVRLRAYMNTPGLPVCQDHLEFWEGVFTGLTHIDAHFLSQTKQSGENFPYL
		VAYQATVCARAQAPPPSWDQMWKCLIRLKPTLHGPTPLLYRLGAVQNEVTLTHPITKYIMTCMSADLEVVT
	`)
	HCV1ASEQ_NS5A = a.ReadString(`
		SGSWLRDIWDWICEVLSDFKTWLKAKLMPQLPGIPFVSCQRGYRGVWRGDGIMHTRCHCGAEITGHVKNG
		TMRIVGPRTCRNMWSGTFPINAYTTGPCTPLPAPNYKFALWRVSAEEYVEIRRVGDFHYVSGMTTDNLKC
		PCQIPSPEFFTELDGVRLHRFAPPCKPLLREEVSFRVGLHEYPVGSQLPCEPEPDVAVLTSMLTDPSHIT
		AEAAGRRLARGSPPSMASSSASQLSAPSLKATCTANHDSPDAELIEANLLWRQEMGGNITRVESENKVVI
		LDSFDPLVAEEDEREVSVPAEILRKSRRFARALPVWARPDYNPPLVETWKKPDYEPPVVHGCPLPPPRSP
		PVPPPRKKRTVVLTESTLSTALAELATKSFGSSSTSGITGDNTTTSSEPAPSGCPPDSDVESYSSMPPLE
		GEPGDPDLSDGSWSTVSSGADTEDVVCC
	`)
	HCV1ASEQ_NS5B = a.ReadString(`
		SMSYSWTGALVTPCAAEEQKLPINALSNSLLRHHNLVYSTTSRSACQRQKKVTFDRLQVLDSHYQDVLKE
		VKAAASKVKANLLSVEEACSLTPPHSAKSKFGYGAKDVRCHARKAVAHINSVWKDLLEDSVTPIDTTIMA
		KNEVFCVQPEKGGRKPARLIVFPDLGVRVCEKMALYDVVSKLPLAVMGSSYGFQYSPGQRVEFLVQAWKS
		KKTPMGFSYDTRCFDSTVTESDIRTEEAIYQCCDLDPQARVAIKSLTERLYVGGPLTNSRGENCGYRRCR
		ASGVLTTSCGNTLTCYIKARAACRAAGLQDCTMLVCGDDLVVICESAGVQEDAASLRAFTEAMTRYSAPP
		GDPPQPEYDLELITSCSSNVSVAHDGAGKRVYYLTRDPTTPLARAAWETARHTPVNSWLGNIIMFAPTLW
		ARMILMTHFFSVLIARDQLEQALNCEIYGACYSIEPLDLPPIIQRLHGLSAFSLHSYSPGEINRVAACLR
		KLGVPPLRAWRHRARSVRARLLSRGGRAAICGKYLFNWAVRTKLKLTPIAAAGRLDLSGWFTAGYSGGDI
		YHSVSHARPRWFWFCLLLLAAGVGIYLLPNR
	`)
)

var HCV1ARefLookup = ap.ReferenceSeqs{
	"NS3":  HCV1ASEQ_NS3,
	"NS5A": HCV1ASEQ_NS5A,
	"NS5B": HCV1ASEQ_NS5B,
}

var Profile = ap.AlignmentProfile{
	StopCodonPenalty:         4,
	GapOpeningPenalty:        10,
	GapExtensionPenalty:      2,
	IndelCodonOpeningBonus:   0,
	IndelCodonExtensionBonus: 2,
	GeneIndelScores:          hcv1aPositionalIndelScores,
	ReferenceSequences:       HCV1ARefLookup,
}