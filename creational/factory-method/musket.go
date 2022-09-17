package main

type musket struct {
	Gun
}

func newMusket() *musket {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

var _ IGun = (*musket)(nil)
