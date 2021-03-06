package openweather

import "errors"

func ValidCoordinates(lat, lon float64) error {
	if !ValidLatitude(lat) {
		return errors.New("invalid latitude")
	}
	if !ValidLongitude(lon) {
		return errors.New("invalid longitude")
	}
	return nil
}

func ValidLongitude(lon float64) bool {
	return lon >= -180 && lon <= 180
}

func ValidLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

var langs = map[string]string{
	"af":    "Afrikaans",
	"al":    "Albanian",
	"ar":    "Arabic",
	"az":    "Azerbaijani",
	"bg":    "Bulgarian",
	"ca":    "Catalan",
	"cz":    "Czech",
	"da":    "Danish",
	"de":    "German",
	"el":    "Greek",
	"en":    "English",
	"eu":    "Basque",
	"fa":    "Persian (Farsi)",
	"fi":    "Finnish",
	"fr":    "French",
	"gl":    "Galician",
	"he":    "Hebrew",
	"hi":    "Hindi",
	"hr":    "Croatian",
	"hu":    "Hungarian",
	"id":    "Indonesian",
	"it":    "Italian",
	"ja":    "Japanese",
	"kr":    "Korean",
	"la":    "Latvian",
	"lt":    "Lithuanian",
	"mk":    "Macedonian",
	"no":    "Norwegian",
	"nl":    "Dutch",
	"pl":    "Polish",
	"pt":    "Portuguese",
	"pr_br": "Português Brasil",
	"ro":    "Romanian",
	"ru":    "Russian",
	"sv":    "Swedish",
	"se":    "Swedish",
	"sk":    "Slovak",
	"sl":    "Slovenian",
	"sp":    "Spanish",
	"es":    "Spanish",
	"sr":    "Serbian",
	"th":    "Thai",
	"tr":    "Turkish",
	"ua":    "Ukrainian",
	"uk":    "Ukrainian",
	"vi":    "Vietnamese",
	"zh_cn": "Chinese Simplified",
	"zh_tw": "Chinese Traditional",
	"zu":    "Zulu",
}

func ValidLang(lang string) bool {
	if _, ok := langs[lang]; ok {
		return true
	}

	return false
}

var units = map[string]string{
	"imperial": "Fahrenheit",
	"metric":   "Celsius",
	"standard": "Kelvin ",
}

func ValidUnit(unit string) bool {
	if _, ok := units[unit]; ok {
		return true
	}

	return false
}
