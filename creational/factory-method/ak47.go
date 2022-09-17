package main

type Ak47 struct {
	Gun
}

func newAk47() *Ak47 {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

var _ IGun = (*Ak47)(nil)
