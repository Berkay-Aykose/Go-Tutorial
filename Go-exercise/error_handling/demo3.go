package errorhandling

import "fmt"

type borderException struct {
	parametre int
	message   string
}

func (b borderException) Error() string {
	return fmt.Sprintf("%d-->%s", b.parametre, b.message)
}

func TahminEt(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{tahmin, "sayı 1 ve 100 arasında değil"}
	}

	return "bildiniz", nil
}
