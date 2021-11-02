package piscine

func Split(s, sep string) []string {
	var texte []string
	taille := len(sep)
	save := 0
	for i := 0; i+taille < len(s); i++ {
		if s[i:taille+i] == sep {
			texte = append(texte, s[save:i])
			save = i + taille
		}
		if i+taille == len(s)-1 && s[i:taille+i] != sep {
			texte = append(texte, s[save:len(s)])
		}
	}
	return (texte)

}
