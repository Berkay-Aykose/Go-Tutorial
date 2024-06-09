package functions

func ToplaVariadic(sayilar ...int) int {
	topla := 0

	for i := 0; i < len(sayilar); i++ {
		topla = topla + sayilar[i]
	}

	return topla
}
