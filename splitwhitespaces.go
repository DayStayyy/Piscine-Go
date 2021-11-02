package piscine

func SplitWhiteSpaces(s string) []string {
	tab := []rune(s)
	var texte []string
	save := 0
	truc := ""
	for j := range tab {
		truc += string(tab[j])
	}
	for i := range tab {
		if tab[i] == 32 || tab[i] == 9 || tab[i] == '\n' || tab[i] == 0 {
			if tab[i-1] != 32 && tab[i-1] != 9 && tab[i-1] != '\n' && tab[i-1] != 0 {
				texte = append(texte, truc[save:i])
				save = i + 1
			} else {
				save = i + 1
			}
		}
		if i == len(tab)-1 && (tab[i] != 32 && tab[i] != 9 && tab[i] != '\n' && tab[i] != 0) {
			texte = append(texte, truc[save:i+1])
		}
	}
	return (texte)

}
