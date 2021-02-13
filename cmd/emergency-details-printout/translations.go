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
	TrMadeWith:                       "Create & customize your own: joonas.fi/emergency-details-printout",
	TrTranslatorCredit:               "Joonas Loppi", // put your (the translatosr's) name, Twitter handle or website here
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
		TrMadeWith:                       "Tee & muokkaa itsellesi oma: joonas.fi/emergency-details-printout",
		TrTranslatorCredit:               "Joonas Loppi",
	},
}
