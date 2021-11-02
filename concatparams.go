package piscine

func ConcatParams(args []string) string {
	texte := ""
	for i := range args {
		texte += args[i]
		if !(i == len(args)-1) {
			texte += "\n"
		}

	}
	return texte
}
