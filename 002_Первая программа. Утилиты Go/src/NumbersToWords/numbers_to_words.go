package NTW

var underTwenty map[int]string = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "eight",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}
var tenners map[int]string = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func GetWord(num int) string {
	var word string
	if num > 99 {
		return "Number too big"
	}
	if num < 1 {
		return "Number too small"
	}
	if num < 20 {
		return underTwenty[num]
	} else {
		result := tenners[(num/10)%10]
		if num%10 != 0 {
			result += "-" + underTwenty[num%10]
		}
		return result
	}
	return word
}
