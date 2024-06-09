package for_range

import "fmt"

func Range() {
	sehirler := []string{"Ankara", "İzmir", "İstanbul", "Eskişehir"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for i /* '_' kullanmak i değişkeni kullnmaycağımzda i yerine yazılarak kullanılır */, sehir := range sehirler {
		fmt.Print(i, " : ")
		fmt.Println(sehir)
	}
}
