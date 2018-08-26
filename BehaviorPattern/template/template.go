package template

import "strings"

type MessageRetriever interface {
	Message() string
}

// Templater interface define three step functions.
type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

//-------------------------------------------------------------------------
// Template will implement the Templater interface and provide some struct
// which want to use the template.Template
// For example, create another struct and wrap the Teplate struct

type Template struct{}

// first step
func (t *Template) first() string {
	return "hello"
}

// ExecuteAlgorithm is the second step.
func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

// third step
func (t *Template) third() string {
	return "template"
}

//-------------------------------------------------------------------------

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

//---------------------------------------------------------------------------------

type adapter struct {
	myFunc func() string
}

// 檢查adatper是否有被塞入匿名函式，沒的話回傳空字串
// 有的話執行該匿名函式，並且預期回傳字串
func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

// 定義一個MessageRetriver的adapter
// 負責先定義好一個adapter struct，包著myFunc
// 然後接收一個匿名函式f並myFunc = f
func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}
