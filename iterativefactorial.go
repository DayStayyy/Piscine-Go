package piscine

func IterativeFactorial(nb int) int {
	if nb > -1 && nb < 21 {
		resultat := 1
		for i := 1; i < nb+1; i++ {
			resultat = resultat * i
		}
		return (resultat)

	}
	return (0)
}
