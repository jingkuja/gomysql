package funs

type Woman struct {
	Dame string
	Dge  int
}

func (w *Woman) Name() string {
	w.Dame = "yu"
	return w.Dame
}
func (w Woman) Age() int {
	w.Dge = 25
	return w.Dge
}
