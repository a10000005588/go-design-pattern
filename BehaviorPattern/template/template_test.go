package template

import (
	"strings"
	"testing"
)

// TestStruct will own a template which will have three step we can choose
type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		// because we didn't create an instance of Template.
		// We should use &TestStruct{} to use any attribute in Template
		s := &TestStruct{}
		res := s.ExecuteAlgorithm(s)

		expectedOrError(res, " world ", t)
	})

	t.Run("Using anonymous functions", func(t *testing.T) {
		// equals to : &AnomnymousTemplate
		m := new(AnonymousTemplate)
		// ExecuteAlgorithm accept a anomonous function
		// and will execute its logic and return string variable
		res := m.ExecuteAlgorithm(func() string {
			return "world"
		})

		expectedOrError(res, " world ", t)
	})
	// 透過adapter pattern和匿名函式來擴充template的接口(ExecuteAlgorithm)
	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T) {
		// 傳給adpater一個會回傳字串的匿名函式
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("Can not continue with a nil MessageRetriever")
		}
		// 宣告template
		template := Template{}
		res := template.ExecuteAlgorithm(messageRetriever)

		expectedOrError(res, " world ", t)
	})
}

func expectedOrError(res string, expected string, t *testing.T) {
	if !strings.Contains(res, expected) {
		t.Errorf("Expected string '%s' was not found on returned string\n", expected)
	}
}
