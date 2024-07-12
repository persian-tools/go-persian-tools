package phonenumbers

var MCIMap = map[string]OperatorDetails{
	"910": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"914": {
		base:     "آذربایجان غربی",
		province: []string{"آذربایجان شرقی", "اردبیل", "اصفهان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"911": {
		base:     "مازندران",
		province: []string{"گلستان", "گیلان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"912": {
		base:     "تهران",
		province: []string{"البرز", "زنجان", "سمنان", "قزوین", "قم", "برخی از شهرستان های استان مرکزی"},
		simTypes: []SimType{Permanent},
		operator: MCI,
	},
	"913": {
		base:     "اصفهان",
		province: []string{"یزد", "چهارمحال و بختیاری", "کرمان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"915": {
		base:     "خراسان رضوی",
		province: []string{"خراسان شمالی", "خراسان جنوبی", "سیستان و بلوچستان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"916": {
		base:     "خوزستان",
		province: []string{"لرستان", "فارس", "اصفهان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"917": {
		base:     "فارس",
		province: []string{"بوشهر", "کهگیلویه و بویر احمد", "هرمزگان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"918": {
		base:     "کرمانشاه",
		province: []string{"کردستان", "ایلام", "همدان"},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"919": {
		base:     "تهران",
		province: []string{"البرز", "سمنان", "قم", "قزوین", "زنجان"},
		simTypes: []SimType{Credit},
		operator: MCI,
	},
	"990": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: MCI,
	},
	"991": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"992": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: MCI,
	},
	"993": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: MCI,
	},
	"994": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: MCI,
	},
	"995": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
	"996": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	},
}

var TALIYA = map[string]OperatorDetails{
	"932": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: Taliya,
	},
}

var RIGHTTEL = map[string]OperatorDetails{
	"920": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Permanent},
		operator: RightTel,
	},
	"921": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: RightTel,
	},
	"922": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: RightTel,
	},
	"923": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: RightTel,
	},
}

var IRANCELL = map[string]OperatorDetails{
	"930": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"933": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"935": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"936": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"937": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"938": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"939": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"901": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"902": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"903": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"905": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"900": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit, Permanent},
		operator: Irancell,
	},
	"904": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: Irancell,
		model:    "سیم‌کارت کودک",
	},
	"941": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: Irancell,
		model:    "TD-LTE",
	},
}

var SHATELMOBILE = map[string]OperatorDetails{
	"998": {
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: ShatelMobile,
	},
}
