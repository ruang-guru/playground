package perusahaan

type Junior struct {
}

func (j Junior) GetSalary() int {
	return 10
}

func (j Junior) TotalDivisonSalary() int {
	return j.GetSalary() // TODO: replace this
}
