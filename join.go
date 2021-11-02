package piscine

func Join(strs []string, sep string) string {
	texte := ""
	for i := range strs {
		texte += strs[i]
		if !(i == len(strs)-1) {
			texte += sep
		}

	}
	return (texte)
}
