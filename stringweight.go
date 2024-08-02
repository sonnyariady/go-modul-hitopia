package hitopia

import (
	"strconv"
	"strings"
)

type StringWeight struct {
	InputString    string
	ArrayQuery     string
	ListInvalid    []string
	ArrQueryResult []string
}

func (sw *StringWeight) IsValid() bool {
	return len(sw.ListInvalid) == 0
}

func NewStringWeight(inputString, arrayQuery string) *StringWeight {
	return &StringWeight{
		InputString: inputString,
		ArrayQuery:  arrayQuery,
		ListInvalid: []string{},
	}
}

func (sw *StringWeight) GenerateResult() {
	if sw.InputString == "" {
		sw.ListInvalid = append(sw.ListInvalid, "InputString cannot be blank")
	}
	if sw.ArrayQuery == "" {
		sw.ListInvalid = append(sw.ListInvalid, "ArrayQuery cannot be blank")
	} else {
		query, err := sw.parseArrayQuery()
		if err != nil {
			sw.ListInvalid = append(sw.ListInvalid, "The string query input is not valid!")
		} else {
			arrSub := sw.GenerateArrayQuery(sw.InputString)
			sw.ArrQueryResult = GenerateQueryResult(query, arrSub)
		}
	}
}

func (sw *StringWeight) parseArrayQuery() ([]int, error) {
	stringsArray := strings.Split(sw.ArrayQuery, ",")
	query := make([]int, len(stringsArray))
	for i, str := range stringsArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		query[i] = num
	}
	return query, nil
}

func GenerateQueryResult(arrQry []int, arrSubQuery []int) []string {
	qryResult := make([]string, len(arrQry))
	for i, intItem := range arrQry {
		if contains(arrSubQuery, intItem) {
			qryResult[i] = "Yes"
		} else {
			qryResult[i] = "No"
		}
	}
	return qryResult
}

func contains(arr []int, item int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func (sw *StringWeight) GenerateArrayQuery(str string) []int {
	lstSubquery := []int{}
	for _, ch := range str {
		iCh := 1
		isHasSubs := true
		for isHasSubs {
			strSubCheck := strings.Repeat(string(ch), iCh)
			if strings.Contains(str, strSubCheck) {
				weightCombination := sw.CalculateStringWeight(strSubCheck)
				if !contains(lstSubquery, weightCombination) {
					lstSubquery = append(lstSubquery, weightCombination)
				}
			} else {
				isHasSubs = false
			}
			iCh++
		}
	}
	return lstSubquery
}

func (sw *StringWeight) CalculateCharacterWeight(c rune) int {
	c = rune(strings.ToLower(string(c))[0])
	return int(c - 'a' + 1)
}

func (sw *StringWeight) CalculateStringWeight(input string) int {
	totalWeight := 0
	for _, c := range input {
		totalWeight += sw.CalculateCharacterWeight(c)
	}
	return totalWeight
}
