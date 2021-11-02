package piscine

func FindNextPrime(nb int) int {
	test := true
	if nb < 2 {
		return (2)
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			test = (false)
		}
	}
	if test == false {
		for j := 1; 1 == 1; j++ {
			test2 := true
			for i := 2; i < nb+j; i++ {
				if (nb+j)%i == 0 {
					test2 = false
					break
				}
			}
			if test2 == true {
				return (nb + j)
			}
		}

	}

	return (nb)

}
