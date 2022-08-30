package model

// MODELS
type ALBUM struct {
	MEDIA    []FILE
	NAME     string
	CATEGORY string
	YEAR     int
	RUNTIME  float32
}

type FILE struct {
	DATE_MODIFIED string
	NAME          string
	PATH          string
	ABLUM_NAME    string
	FILE_TYPE     string
	SIZE          int64
	LENGTH        float64
}
