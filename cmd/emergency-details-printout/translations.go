package main

type Translated struct {
	TrInCaseOfEmergency              string
	TrEmergencyNumber                string
	TrAddress                        string
	TrFire                           string
	TrBleeding                       string
	TrResuscitationAdult             string
	TrResuscitationAdultInstructions string
	TrWaterLeak                      string
	TrMaintenanceCompany             string
	TrVeterinarian                   string
	TrInformationUpToDateChecked     string
	TrContactDetails                 string
	Tr5BackBlows                     string
	Tr5AbdominalThrusts              string
	TrMadeWith                       string
	TrTranslatorCredit               string // currently not shown
}

// this is an reference (also an example) translation bundle
//
// if you want to translate but don't know how to send a "pull request" (or it seems like too much work),
// it's ok! you can take this block of code, translate it (only the things inside quotes!) and send it
// to me and I'll add the translated texts back!
var englishPhrases = Translated{
	TrInCaseOfEmergency:              "In case of emergency",
	TrEmergencyNumber:                "Emergency number",
	TrAddress:                        "Address",
	TrFire:                           "Fire",
	TrBleeding:                       "Bleeding",
	TrResuscitationAdult:             "Resuscitation, adult",
	TrResuscitationAdultInstructions: "30 presses, 2 blows (hold nose closed). Rhythm: 2 presses/second",
	TrWaterLeak:                      "Water leak",
	TrMaintenanceCompany:             "Maintenance company",
	TrVeterinarian:                   "Veterinarian",
	TrInformationUpToDateChecked:     "Information up-to-date checked",
	TrContactDetails:                 "Contact details",
	Tr5BackBlows:                     "Give 5 back blows",
	Tr5AbdominalThrusts:              "Give 5 abdominal thrusts",
	TrMadeWith:                       "Create & customize your own: joonas.fi/emergency-details-printout",
	TrTranslatorCredit:               "https://joonas.fi/", // put your (the translator's) name, Twitter handle or website here
}

var translations = map[string]Translated{
	"en": englishPhrases,
	"fi": {
		TrInCaseOfEmergency:              "Hädän sattuessa",
		TrEmergencyNumber:                "Hätänumero",
		TrAddress:                        "Osoite",
		TrFire:                           "Tulipalo",
		TrBleeding:                       "Verenvuoto",
		TrResuscitationAdult:             "Elvytys, aikuisen",
		TrResuscitationAdultInstructions: "30 painelua, 2 puhallusta (paina sieraimet kiinni). Rytmi: 2 painelua/sek",
		TrWaterLeak:                      "Vesivuoto",
		TrMaintenanceCompany:             "Huoltoyhtiö",
		TrVeterinarian:                   "Eläinlääkäri",
		TrInformationUpToDateChecked:     "Ajantasaisuus tarkistettu",
		TrContactDetails:                 "Yhteystiedot",
		Tr5BackBlows:                     "Kämmenellä 5 kertaa lapaluiden väliin",
		Tr5AbdominalThrusts:              "5 vetoa ylävatsan kohdalta",
		TrMadeWith:                       "Tee & muokkaa itsellesi oma: joonas.fi/emergency-details-printout",
		TrTranslatorCredit:               "https://joonas.fi/",
	},
	"zh-CN": {
		TrInCaseOfEmergency:              "如果发生紧急情况",
		TrEmergencyNumber:                "紧急号码",
		TrAddress:                        "地址",
		TrFire:                           "火灾",
		TrBleeding:                       "受伤",
		TrResuscitationAdult:             "心肺复苏（成年人）",
		TrResuscitationAdultInstructions: "按压 30 次，吹气 2 次（夹住对方鼻子）。按压频率：2 次/秒",
		TrWaterLeak:                      "漏水",
		TrMaintenanceCompany:             "物业",
		TrVeterinarian:                   "兽医",
		TrInformationUpToDateChecked:     "信息更新时间",
		TrContactDetails:                 "联系方式",
		Tr5BackBlows:                     "拍击背部 5 次",
		Tr5AbdominalThrusts:              "冲击腹部 5 次",
		TrMadeWith:                       "创建并定制一份属于你的：joonas.fi/emergency-details-printout",
		TrTranslatorCredit:               "RainSlide",
	},
}
