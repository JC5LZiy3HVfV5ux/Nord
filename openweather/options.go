package openweather

type options struct {
	key  string
	lang string
	unit string
}

func initOptions(key string) *options {
	return &options{
		key:  key,
		lang: "en",
		unit: "standard",
	}
}

func (o *options) getMap() map[string]string {
	return map[string]string{
		"appid": o.key,
		"units": o.unit,
		"lang":  o.lang,
	}
}
