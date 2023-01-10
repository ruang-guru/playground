package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	total := vp.GetSalary()
	for _, val := range vp.Subordinate {
		total += val.TotalDivisonSalary()
	}
	return total // TODO: replace this
}
