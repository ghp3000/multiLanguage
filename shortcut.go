package multiLanguage

var v = NewValidator(LocaleZh)

func SetDefaultLocale(locale Locale) error {
	return v.SetDefaultLocale(locale)
}
func Register(locale Locale, fieldFilename string) error {
	return v.Register(locale, fieldFilename)
}
func Locales() []string {
	return v.Locales()
}
func Validate(data interface{}, lang string) *ValidatError {
	return v.Validate(data, lang)
}
func Validates(data interface{}, lang string) []ValidatError {
	return v.Validates(data, lang)
}
