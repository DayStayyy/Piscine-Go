package piscine

func IterativePower(nb int, power int) int {
	if power > 0 && nb < 21 {
		stockage := nb
		for i := 0; i < power-1; i++ {
			nb = nb * stockage
		}
		return (nb)

	} else if power == 0 {
		return (1)

	}
	return (0)
}
