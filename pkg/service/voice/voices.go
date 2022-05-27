package voice

var supportedLanguage = []string{"en",
	"en-UK",
	"en-AU",
	"ja",
	"de",
	"es",
	"ru",
	"ar",
	"bn",
	"cs",
	"da",
	"nl",
	"fi",
	"el",
	"hi",
	"hu",
	"id",
	"km",
	"la",
	"it",
	"no",
	"pl",
	"sk",
	"sv",
	"th",
	"tr",
	"uk",
	"vi",
	"af",
	"bg",
	"ca",
	"cy",
	"et",
	"fr",
	"gu",
	"is",
	"jv",
	"kn",
	"ko",
	"lv",
	"ml",
	"mr",
	"ms",
	"ne",
	"pt",
	"ro",
	"si",
	"sr",
	"su",
	"ta",
	"te",
	"tl",
	"ur",
	"zh",
	"sw",
	"sq",
	"my",
	"mk",
	"hy",
	"hr",
	"eo",
	"bs"}

func IsSupported(language string) bool {
	for _, item := range supportedLanguage {
		if item == language {
			return true
		}
	}
	return false
}
