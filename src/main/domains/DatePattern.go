package main_domains

type DatePattern string

const PATTERN_YYYYMMDD = "2006-01-02"
const PATTERN_DDMMYYYY = "02-01-2006"

const (
	YYYYMMDD DatePattern = PATTERN_YYYYMMDD
	DDMMYYYY DatePattern = PATTERN_DDMMYYYY
)

func (s DatePattern) GetDescription() string {
	switch s {
	case YYYYMMDD:
		return PATTERN_YYYYMMDD
	case DDMMYYYY:
		return PATTERN_DDMMYYYY
	}
	return PATTERN_YYYYMMDD
}
