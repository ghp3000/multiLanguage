package multiLanguage

import (
	"errors"
	"fmt"

	"github.com/ghp3000/multiLanguage/translator"
)

type MultiLanguage struct {
	store       map[string]*translator.Translator
	defaultLang string
}

func NewMultiLanguage() *MultiLanguage {
	return &MultiLanguage{
		store: make(map[string]*translator.Translator),
	}
}
func (m *MultiLanguage) Register(name, filename string, defaultLang bool) error {
	m.store[name] = translator.NewTranslator(name, filename)
	if defaultLang {
		m.defaultLang = name
		return m.Load(name)
	}
	return nil
}

// SetDefaultLocale 设置默认语言,线程不安全,最佳时机是初始化阶段执行
func (m *MultiLanguage) SetDefaultLocale(locale Locale) error {
	if !IsSupportedLocale(locale) {
		return fmt.Errorf("%s not supported", locale)
	}
	m.defaultLang = string(locale)
	return nil
}

func (m *MultiLanguage) Load(name string) error {
	trans, ok := m.store[name]
	if !ok {
		return errors.New("translator not found")
	}
	return trans.Load()
}
func (m *MultiLanguage) Translate(key, lang string) string {
	trans, ok := m.store[lang]
	if !ok {
		trans = m.store[m.defaultLang]
	}
	if trans == nil {
		return key
	}
	return trans.Translate(key)
}
