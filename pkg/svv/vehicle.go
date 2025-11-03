package svv

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const BASE_URL = "https://akfell-datautlevering.atlas.vegvesen.no/enkeltoppslag/kjoretoydata"

func GetVehicle(apiKey string, regNumber string) (KjoretoyData, error) {
	url := BASE_URL + "?kjennemerke=" + regNumber

	var kjoretoyData KjoretoyData
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return kjoretoyData, err
	}

	req.Header.Add("SVV-Authorization", fmt.Sprintf("Apikey %s", apiKey))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return kjoretoyData, err
	}

	statuscode := response.StatusCode
	if statuscode == 400 {
		return kjoretoyData, fmt.Errorf("invalid regnumber: %d", statuscode)
	}
	if statuscode == 403 {
		return kjoretoyData, fmt.Errorf("invalid api-key: %d", statuscode)
	}
	if statuscode == 429 {
		return kjoretoyData, fmt.Errorf("quota exceeded: %d", statuscode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return kjoretoyData, err
	}

	err = json.Unmarshal(body, &kjoretoyData)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return kjoretoyData, err
	}

	return kjoretoyData, nil
}

type KjoretoyData struct {
	KjoretoydataListe []KjoretoydataListe `json:"kjoretoydataListe"`
}

type KjoretoyID struct {
	Kjennemerke       string `json:"kjennemerke"`
	Understellsnummer string `json:"understellsnummer"`
}

type Forstegangsregistrering struct {
	RegistrertForstegangNorgeDato string `json:"registrertForstegangNorgeDato"`
}

type Kjennemerke struct {
	FomTidspunkt        time.Time       `json:"fomTidspunkt"`
	Kjennemerke         string          `json:"kjennemerke"`
	Kjennemerkekategori string          `json:"kjennemerkekategori"`
	Kjennemerketype     Kjennemerketype `json:"kjennemerketype"`
}

type Kjennemerketype struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Registrering struct {
	FomTidspunkt        time.Time           `json:"fomTidspunkt"`
	KjoringensArt       KjoringensArt       `json:"kjoringensArt"`
	Registreringsstatus Registreringsstatus `json:"registreringsstatus"`
}

type KjoringensArt struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Registreringsstatus struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type ForstegangsGodkjenning struct {
	ForstegangRegistrertDato string                `json:"forstegangRegistrertDato"`
	GodkjenningsID           string                `json:"godkjenningsId"`
	Godkjenningsundertype    Godkjenningsundertype `json:"godkjenningsundertype"`
	GyldigFraDato            string                `json:"gyldigFraDato"`
	GyldigFraDatoTid         time.Time             `json:"gyldigFraDatoTid"`
	Unntak                   []any                 `json:"unntak"`
}

type Godkjenningsundertype struct {
	KodeNavn           string `json:"kodeNavn"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Kjoretoymerknad struct {
	Merknad         string `json:"merknad"`
	MerknadtypeKode string `json:"merknadtypeKode"`
}

type Registreringsbegrensninger struct {
	Registreringsbegrensning []any `json:"registreringsbegrensning"`
}

type Typegodkjenningnummer struct {
	Direktiv  string `json:"direktiv"`
	Land      string `json:"land"`
	Serie     string `json:"serie"`
	Utvidelse string `json:"utvidelse"`
}

type EfTypegodkjenning struct {
	TypegodkjenningNrTekst string                `json:"typegodkjenningNrTekst"`
	Typegodkjenningnummer  Typegodkjenningnummer `json:"typegodkjenningnummer"`
	Variant                string                `json:"variant"`
	Versjon                string                `json:"versjon"`
}

type KjoretoyAvgiftsKode struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}
type NasjonalGodkjenning struct {
	NasjonaltGodkjenningsAr          string `json:"nasjonaltGodkjenningsAr"`
	NasjonaltGodkjenningsHovednummer string `json:"nasjonaltGodkjenningsHovednummer"`
	NasjonaltGodkjenningsUndernummer string `json:"nasjonaltGodkjenningsUndernummer"`
}

type TekniskKode struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type TekniskUnderkode struct {
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Kjoretoyklassifisering struct {
	Beskrivelse                string              `json:"beskrivelse"`
	EfTypegodkjenning          EfTypegodkjenning   `json:"efTypegodkjenning"`
	KjoretoyAvgiftsKode        KjoretoyAvgiftsKode `json:"kjoretoyAvgiftsKode"`
	NasjonalGodkjenning        NasjonalGodkjenning `json:"nasjonalGodkjenning"`
	SpesielleKjennetegn        string              `json:"spesielleKjennetegn"`
	TekniskKode                TekniskKode         `json:"tekniskKode"`
	TekniskUnderkode           TekniskUnderkode    `json:"tekniskUnderkode"`
	ISamsvarMedTypegodkjenning bool                `json:"iSamsvarMedTypegodkjenning"`
}

type Krav struct {
	Kravomrade      Kravomrade      `json:"kravomrade"`
	Kravoppfyllelse Kravoppfyllelse `json:"kravoppfyllelse"`
}

type Kravoppfyllelse struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Kravomrade struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Aksel struct {
	AvstandTilNesteAksling  int    `json:"avstandTilNesteAksling"`
	DrivAksel               bool   `json:"drivAksel"`
	ID                      int    `json:"id"`
	PlasseringAksel         string `json:"plasseringAksel"`
	Sporvidde               int    `json:"sporvidde"`
	TekniskTillattAkselLast int    `json:"tekniskTillattAkselLast"`
}

type AkselListe struct {
	Aksel []Aksel `json:"aksel"`
}

type AkselGruppe struct {
	AkselListe                    AkselListe `json:"akselListe"`
	ID                            int        `json:"id"`
	PlasseringAkselGruppe         string     `json:"plasseringAkselGruppe"`
	TekniskTillattAkselGruppeLast int        `json:"tekniskTillattAkselGruppeLast"`
}

type Akslinger struct {
	AkselGruppe  []AkselGruppe `json:"akselGruppe"`
	AntallAksler int           `json:"antallAksler"`
}

type Bremser struct {
	TilhengerBremseforbindelse []any `json:"tilhengerBremseforbindelse"`
}

type DekkOgFelg struct {
	AkselDekkOgFelgKombinasjon []AkselDekkOgFelgKombinasjon `json:"akselDekkOgFelgKombinasjon"`
}

type AkselDekkOgFelgKombinasjon struct {
	AkselDekkOgFelg []AkselDekkOgFelg `json:"akselDekkOgFelg"`
}

type AkselDekkOgFelg struct {
	AkselID             int    `json:"akselId"`
	BelastningskodeDekk string `json:"belastningskodeDekk"`
	Dekkdimensjon       string `json:"dekkdimensjon"`
	Felgdimensjon       string `json:"felgdimensjon"`
	HastighetskodeDekk  string `json:"hastighetskodeDekk"`
	Innpress            string `json:"innpress"`
}

type Dimensjoner struct {
	Bredde int `json:"bredde"`
	Lengde int `json:"lengde"`
}

type Generelt struct {
	Fabrikant         []any       `json:"fabrikant"`
	Handelsbetegnelse []string    `json:"handelsbetegnelse"`
	Merke             []Merke     `json:"merke"`
	TekniskKode       TekniskKode `json:"tekniskKode"`
	Typebetegnelse    string      `json:"typebetegnelse"`
}

type Merke struct {
	Merke     string `json:"merke"`
	MerkeKode string `json:"merkeKode"`
}

type KjennemerketypeBak struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type KjennemerkestorrelseBak struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type KjennemerketypeForan struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type KjennemerkestorrelseForan struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type PlasseringFabrikasjonsplate struct {
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type PlasseringUnderstellsnummer struct {
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type RFarge struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type KarosseriOgLasteplan struct {
	AntallDorer                 []any                         `json:"antallDorer"`
	DorUtforming                []any                         `json:"dorUtforming"`
	KjennemerketypeBak          KjennemerketypeBak            `json:"kjennemerketypeBak"`
	KjennemerkestorrelseBak     KjennemerkestorrelseBak       `json:"kjennemerkestorrelseBak"`
	KjennemerketypeForan        KjennemerketypeForan          `json:"kjennemerketypeForan"`
	KjennemerkestorrelseForan   KjennemerkestorrelseForan     `json:"kjennemerkestorrelseForan"`
	KjoringSide                 string                        `json:"kjoringSide"`
	OppbygningUnderstellsnummer string                        `json:"oppbygningUnderstellsnummer"`
	PlasseringFabrikasjonsplate []PlasseringFabrikasjonsplate `json:"plasseringFabrikasjonsplate"`
	PlasseringUnderstellsnummer []PlasseringUnderstellsnummer `json:"plasseringUnderstellsnummer"`
	RFarge                      []RFarge                      `json:"rFarge"`
}

type Miljodata struct {
	EuroKlasse             EuroKlasse               `json:"euroKlasse"`
	MiljoOgdrivstoffGruppe []MiljoOgdrivstoffGruppe `json:"miljoOgdrivstoffGruppe"`
	OkoInnovasjon          bool                     `json:"okoInnovasjon"`
}

type EuroKlasse struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type MiljoOgdrivstoffGruppe struct {
	DrivstoffKodeMiljodata DrivstoffKodeMiljodata `json:"drivstoffKodeMiljodata"`
	Lyd                    Lyd                    `json:"lyd"`
}

type Lyd struct {
	Standstoy             int                 `json:"standstoy"`
	StoyMalingOppgittAv   StoyMalingOppgittAv `json:"stoyMalingOppgittAv"`
	VedAntallOmdreininger int                 `json:"vedAntallOmdreininger"`
}

type StoyMalingOppgittAv struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type DrivstoffKodeMiljodata struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Girkassetype struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type HybridKategori struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Drivstoff struct {
	DrivstoffKode   DrivstoffKode `json:"drivstoffKode"`
	MaksNettoEffekt float64       `json:"maksNettoEffekt"`
}

type DrivstoffKode struct {
	KodeBeskrivelse    string `json:"kodeBeskrivelse"`
	KodeNavn           string `json:"kodeNavn"`
	KodeTypeID         string `json:"kodeTypeId"`
	KodeVerdi          string `json:"kodeVerdi"`
	TidligereKodeVerdi []any  `json:"tidligereKodeVerdi"`
}

type Motor struct {
	AntallSylindre int         `json:"antallSylindre"`
	Drivstoff      []Drivstoff `json:"drivstoff"`
	MotorKode      string      `json:"motorKode"`
	Slagvolum      int         `json:"slagvolum"`
}

type MotorOgDrivverk struct {
	Girkassetype          Girkassetype   `json:"girkassetype"`
	GirutvekslingPrGir    []any          `json:"girutvekslingPrGir"`
	HybridKategori        HybridKategori `json:"hybridKategori"`
	MaksimumHastighet     []int          `json:"maksimumHastighet"`
	MaksimumHastighetMalt []any          `json:"maksimumHastighetMalt"`
	Motor                 []Motor        `json:"motor"`
}

type Persontall struct {
	SitteplasserForan  int `json:"sitteplasserForan"`
	SitteplasserTotalt int `json:"sitteplasserTotalt"`
}
type Tilhengerkopling struct {
	Kopling []any `json:"kopling"`
}
type Vekter struct {
	Egenvekt                      int   `json:"egenvekt"`
	EgenvektMinimum               int   `json:"egenvektMinimum"`
	Nyttelast                     int   `json:"nyttelast"`
	TillattTaklast                int   `json:"tillattTaklast"`
	TillattTilhengervektMedBrems  int   `json:"tillattTilhengervektMedBrems"`
	TillattTilhengervektUtenBrems int   `json:"tillattTilhengervektUtenBrems"`
	TillattTotalvekt              int   `json:"tillattTotalvekt"`
	TillattVertikalKoplingslast   int   `json:"tillattVertikalKoplingslast"`
	TillattVogntogvekt            int   `json:"tillattVogntogvekt"`
	VogntogvektAvhBremsesystem    []any `json:"vogntogvektAvhBremsesystem"`
}

type TekniskeData struct {
	Akslinger            Akslinger            `json:"akslinger"`
	Bremser              Bremser              `json:"bremser"`
	DekkOgFelg           DekkOgFelg           `json:"dekkOgFelg"`
	Dimensjoner          Dimensjoner          `json:"dimensjoner"`
	Generelt             Generelt             `json:"generelt"`
	KarosseriOgLasteplan KarosseriOgLasteplan `json:"karosseriOgLasteplan"`
	Miljodata            Miljodata            `json:"miljodata"`
	MotorOgDrivverk      MotorOgDrivverk      `json:"motorOgDrivverk"`
	OvrigeTekniskeData   []any                `json:"ovrigeTekniskeData"`
	Persontall           Persontall           `json:"persontall"`
	Tilhengerkopling     Tilhengerkopling     `json:"tilhengerkopling"`
	Vekter               Vekter               `json:"vekter"`
}

type TekniskGodkjenning struct {
	GodkjenningsID         string                 `json:"godkjenningsId"`
	Godkjenningsundertype  Godkjenningsundertype  `json:"godkjenningsundertype"`
	GyldigFraDato          string                 `json:"gyldigFraDato"`
	GyldigFraDatoTid       time.Time              `json:"gyldigFraDatoTid"`
	Kjoretoyklassifisering Kjoretoyklassifisering `json:"kjoretoyklassifisering"`
	Krav                   []Krav                 `json:"krav"`
	TekniskeData           TekniskeData           `json:"tekniskeData"`
	Unntak                 []any                  `json:"unntak"`
}

type Godkjenning struct {
	ForstegangsGodkjenning     ForstegangsGodkjenning     `json:"forstegangsGodkjenning"`
	Kjoretoymerknad            []Kjoretoymerknad          `json:"kjoretoymerknad"`
	Registreringsbegrensninger Registreringsbegrensninger `json:"registreringsbegrensninger"`
	TekniskGodkjenning         TekniskGodkjenning         `json:"tekniskGodkjenning"`
	Tilleggsgodkjenninger      []any                      `json:"tilleggsgodkjenninger"`
}

type KjoretoydataListe struct {
	KjoretoyID                KjoretoyID                `json:"kjoretoyId"`
	Forstegangsregistrering   Forstegangsregistrering   `json:"forstegangsregistrering"`
	Kjennemerke               []Kjennemerke             `json:"kjennemerke"`
	Registrering              Registrering              `json:"registrering"`
	Godkjenning               Godkjenning               `json:"godkjenning"`
	PeriodiskKjoretoyKontroll PeriodiskKjoretoyKontroll `json:"periodiskKjoretoyKontroll"`
}

type PeriodiskKjoretoyKontroll struct {
	Kontrollfrist string `json:"kontrollfrist"`
	SistGodkjent  string `json:"sistGodkjent"`
}
