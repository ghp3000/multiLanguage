package translator

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type Translator struct {
	name     string
	filename string
	once     sync.Once
	err      error
	store    map[string]string
}

// NewTranslator i18n翻译器.懒加载或主动加载
// name: 语言名
// filename: key=value文件路径
func NewTranslator(name, filename string) *Translator {
	return &Translator{name: name, filename: filename, store: nil}
}

// Name 返回语言名
func (t *Translator) Name() string {
	return t.name
}

// Load 未加载立即加载,已加载返回加载时的错误
func (t *Translator) Load() error {
	t.once.Do(func() {
		t.store, t.err = loadFieldsFile(t.filename) //将key=value文本文件逐行解析成map
		if t.store == nil {
			t.store = make(map[string]string)
		}
	})
	return t.err
}

// Translate 翻译,含懒加载
func (t *Translator) Translate(key string) string {
	_ = t.Load()
	str, ok := t.store[key]
	if ok {
		return str
	}
	return key
}

func loadFieldsFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // 跳过空行和注释
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key != "" {
				result[key] = value
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
