package piscine

func Capitalize(s string) string {
	tab := []rune(s)
	var result = make([]rune, len(tab))

	for i := 0; i < len(tab); i++ { //Tout mettre en minuscule
		if tab[i] > 64 && tab[i] < 91 {
			result[i] = tab[i] + 32
		} else {
			result[i] = tab[i]
		}
	}

	if result[0] > 96 && result[0] < 123 {
		result[0] -= 32
	}
	for j := 0; j < len(result); j++ {
		if (result[j] > 96 && result[j] < 123) && ((result[j-1] < 65 || result[j-1] > 90) && (result[j-1] < 97 || result[j-1] > 122) && !(result[j-1] > 47 && result[j-1] < 58)) { //Si j est une lettre et que j-1 est un caractère spécial
			result[j] -= 32
		}
	}
	return string(result)
}
