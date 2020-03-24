package Palindrom

func CreatePalindrom(txt string) string {
	var txtPalindrom string
	for i := len(txt) - 1; i >= 0; i-- {
		txtPalindrom += string(txt[i])
	}
	return txtPalindrom
}
