package ddp

import (
	"errors"
	"strings"
)

type sectionType int

const (
	typeUnknown sectionType = iota
	typeConstants
	typeStruct
	typeExample
	typeParams
)

// nextSubSection defines the range of a sub-title+content:
//
//  ###### Default Message Notification Level
//
//  | Key           | Value |
//  | ------------- | ----- |
//  | ALL_MESSAGES  | 0     |
//  | ONLY_MENTIONS | 1     |
func nextSubSection(offset uint, data []byte) (t sectionType, start, length uint, err error) {
	if beginning := strings.Index(string(data[offset:]), "######"); beginning == -1 {
		return typeUnknown, 0, 0, errors.New("no more sub sections")
	} else {
		start = offset + uint(beginning)
	}

	var end int
	if end = strings.Index(string(data[start+7:]), "###"); end == -1 {
		length = uint(len(data)) - start
	} else {
		length = uint(end) + 7
	}

	titleAndBody := strings.SplitN(string(data[start:start+length]), "\n", 2)
	if len(titleAndBody) != 2 {
		t = typeUnknown
	} else if strings.Contains(strings.ToLower(titleAndBody[0]), "structure") {
		t = typeStruct
	} else if strings.Contains(strings.ToLower(titleAndBody[0]), "fields") {
		t = typeStruct
	} else if strings.Contains(strings.ToLower(titleAndBody[0]), "response") {
		t = typeStruct
	} else if strings.Contains(strings.ToLower(titleAndBody[0]), "example") {
		t = typeExample
	} else if strings.Contains(strings.ToLower(titleAndBody[0]), "params") {
		t = typeParams
	} else if !strings.Contains(strings.ToLower(titleAndBody[0]), "structure") {
		t = typeConstants
	} else {
		t = typeUnknown
	}

	return t, start, length, nil
}

func ParseConstants(data []byte) {

}
