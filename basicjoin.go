package piscine

func BasicJoin(elems []string) string {
	texte := ""
	for i := range elems {
		texte += elems[i]
	}
	return (texte)
}
