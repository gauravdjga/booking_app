package models

type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	Flush string
	CSRFToken string
	Warning string
	Error string
}