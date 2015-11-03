package messageparser

import (
	"fmt"
	"github.com/krolim/shoppinglist/dbmanager"
	"strings"
)

var bgAmountStrings = map[string]int{
	"едно":   1,
	"една":   1,
	"един":   1,
	"две":    2,
	"два":    2,
	"три":    3,
	"четири": 4,
	"пет":    5,
	"шест":   6,
	"седем":  7,
	"осем":   8,
	"девет":  9,
	"десет":  10,
}

var measureStrings = map[string]string{
	"кг":        "kg",
	"к":         "kg",
	"гр":        "gr",
	"килограма": "kg",
	"килограм":  "kg",
	"грам":      "gr",
	"г":         "gr",
	"gr":        "gr",
	"kg":        "kg",
	"kilogram":  "kg",
	"gram":      "gr",
	"грама":     "gr",
}

func init() {

}

func ParseMsg(msg string) *dbmanager.Order {
	fmt.Printf("%v\n", msg)
	parts := strings.Split(msg, " ")
	var amount int
	var measure string
	for _, i := range parts {
		amnt, ok := bgAmountStrings[i]
		if ok {
			// fmt.Printf("numeric value found: %v\n", val)
			amount = amnt
		} else {
			msr, isOk := measureStrings[i]
			if isOk {
				measure = msr
			} else {
				if amount > 0 {
					fmt.Printf("%v %v x %v\n", amount, measure, i)
				} else {
					fmt.Printf("%v\n", i)
				}
				amount = -1
				measure = ""
			}
		}
	}
	return nil
}
