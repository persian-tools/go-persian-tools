package bank

import "strconv"

type ShebaProcess struct {
	normal    string
	formatted string
}

func parsian(strInput string) ShebaProcess {
	str := strInput[14:]
	formatted := "0" + str[0:2] + "-0" + str[2:9] + "-" + str[9:12]

	return ShebaProcess{normal: str, formatted: formatted}
}
func pasargad(strInput string) ShebaProcess {
	str := strInput[7:]
	for strconv.Itoa(int(str[0])) == "0" {
		str = str[1:]
	}
	str = str[0 : len(str)-2]
	formatted :=
		str[0:3] + "-" + str[3:6] + "-" + str[6:14] + "-" + str[14:15]

	return ShebaProcess{normal: str, formatted: formatted}
}

func bankshahr(strInput string) ShebaProcess {
	str := strInput[7:]
	for strconv.Itoa(int(str[0])) == "0" {
		str = str[1:]
	}

	return ShebaProcess{normal: str, formatted: str}
}

type shebaResultHash struct {
	Name                   string
	Code                   string
	NickName               string
	PersianName            string
	AccountNumber          string
	AccountNumberAvailable bool
	FormattedAccountNumber string
	Process                func(str string) ShebaProcess
}

func shebaHashTable(index string) shebaResultHash {
	bankCodes := []shebaResultHash{
		{
			NickName:               "central-bank",
			Name:                   "Central Bank of Iran",
			PersianName:            "بانک مرکزی جمهوری اسلامی ایران",
			Code:                   "010",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "sanat-o-madan",
			Name:                   "Sanat O Madan Bank",
			PersianName:            "بانک صنعت و معدن",
			Code:                   "011",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "mellat",
			Name:                   "Mellat Bank",
			PersianName:            "بانک ملت",
			Code:                   "012",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "refah",
			Name:                   "Refah Bank",
			PersianName:            "بانک رفاه کارگران",
			Code:                   "013",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "maskan",
			Name:                   "Maskan Bank",
			PersianName:            "بانک مسکن",
			Code:                   "014",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "sepah",
			Name:                   "Sepah Bank",
			PersianName:            "بانک سپه",
			Code:                   "015",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "keshavarzi",
			Name:                   "Keshavarzi",
			PersianName:            "بانک کشاورزی",
			Code:                   "016",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "melli",
			Name:                   "Melli",
			PersianName:            "بانک ملی ایران",
			Code:                   "017",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "tejarat",
			Name:                   "Tejarat Bank",
			PersianName:            "بانک تجارت",
			Code:                   "018",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "saderat",
			Name:                   "Saderat Bank",
			PersianName:            "بانک صادرات ایران",
			Code:                   "019",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "tosee-saderat",
			Name:                   "Tose Saderat Bank",
			PersianName:            "بانک توسعه صادرات",
			Code:                   "020",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "post",
			Name:                   "Post Bank",
			PersianName:            "پست بانک ایران",
			Code:                   "021",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "toose-taavon",
			Name:                   "Tosee Taavon Bank",
			PersianName:            "بانک توسعه تعاون",
			Code:                   "022",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "tosee",
			Name:                   "Tosee Bank",
			PersianName:            "موسسه اعتباری توسعه",
			Code:                   "051",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "ghavamin",
			Name:                   "Ghavamin Bank",
			PersianName:            "بانک قوامین",
			Code:                   "052",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "karafarin",
			Name:                   "Karafarin Bank",
			PersianName:            "بانک کارآفرین",
			Code:                   "053",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "parsian",
			Name:                   "Parsian Bank",
			PersianName:            "بانک پارسیان",
			Code:                   "054",
			AccountNumberAvailable: true,
			Process:                parsian,
		},
		{
			NickName:               "eghtesad-novin",
			Name:                   "Eghtesad Novin Bank",
			PersianName:            "بانک اقتصاد نوین",
			Code:                   "055",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "saman",
			Name:                   "Saman Bank",
			PersianName:            "بانک سامان",
			Code:                   "056",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "pasargad",
			Name:                   "Pasargad Bank",
			PersianName:            "بانک پاسارگاد",
			Code:                   "057",
			AccountNumberAvailable: true,
			Process:                pasargad,
		},
		{
			NickName:               "sarmayeh",
			Name:                   "Sarmayeh Bank",
			PersianName:            "بانک سرمایه",
			Code:                   "058",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "sina",
			Name:                   "Sina Bank",
			PersianName:            "بانک سینا",
			Code:                   "059",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "mehr-iran",
			Name:                   "Mehr Iran Bank",
			PersianName:            "بانک مهر ایران",
			Code:                   "060",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "shahr",
			Name:                   "City Bank",
			PersianName:            "بانک شهر",
			Code:                   "061",
			AccountNumberAvailable: true,
			Process:                bankshahr,
		},
		{
			NickName:               "ayandeh",
			Name:                   "Ayandeh Bank",
			PersianName:            "بانک آینده",
			Code:                   "062",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "ansar",
			Name:                   "Ansar Bank",
			PersianName:            "بانک انصار",
			Code:                   "063",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "gardeshgari",
			Name:                   "Gardeshgari Bank",
			PersianName:            "بانک گردشگری",
			Code:                   "064",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "hekmat-iranian",
			Name:                   "Hekmat Iranian Bank",
			PersianName:            "بانک حکمت ایرانیان",
			Code:                   "065",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "dey",
			Name:                   "Dey Bank",
			PersianName:            "بانک دی",
			Code:                   "066",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "iran-zamin",
			Name:                   "Iran Zamin Bank",
			PersianName:            "بانک ایران زمین",
			Code:                   "069",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "resalat",
			Name:                   "Resalat Bank",
			PersianName:            "بانک قرض الحسنه رسالت",
			Code:                   "070",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "kosar",
			Name:                   "Kosar Credit Institute",
			PersianName:            "موسسه اعتباری کوثر",
			Code:                   "073",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "melal",
			Name:                   "Melal Credit Institute",
			PersianName:            "موسسه اعتباری ملل",
			Code:                   "075",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "middle-east-bank",
			Name:                   "Middle East Bank",
			PersianName:            "بانک خاورمیانه",
			Code:                   "078",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "noor-bank",
			Name:                   "Noor Credit Institution",
			PersianName:            "موسسه اعتباری نور",
			Code:                   "080",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "mehr-eqtesad",
			Name:                   "Mehr Eqtesad Bank",
			PersianName:            "بانک مهر اقتصاد",
			Code:                   "079",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "mehr-iran",
			Name:                   "Mehr Iran Bank",
			PersianName:            "بانک مهر ایران",
			Code:                   "090",
			AccountNumberAvailable: false,
		},
		{
			NickName:               "iran-venezuela",
			Name:                   "Iran and Venezuela Bank",
			PersianName:            "بانک ایران و ونزوئلا",
			Code:                   "095",
			AccountNumberAvailable: false,
		},
	}
	bank := shebaResultHash{}
	for _, n := range bankCodes {
		if n.Code == index {
			return n
		}
	}
	return bank
}
