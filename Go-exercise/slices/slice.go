package slices

import "fmt"

func Slice() {
	isimler := make([]string, 3)

	fmt.Println(isimler) //boş yazdırır

	isimler[0] = "Berkay"
	isimler[1] = "Zeynep"
	isimler[2] = "Egemen"

	fmt.Println(isimler) //yukarıdaki 3 ismi yazdırır

	isimler = append(isimler, "Büşra")

	fmt.Println(isimler) //yeni eklenen isimle bir yazdırır

	/***************************SLİCE KOPYALAMA*********************/

	sehirler := []string{"Ankara", "İzmir", "İstanbul"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)

	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirler)

	fmt.Println(sehirler[1:3])
}
