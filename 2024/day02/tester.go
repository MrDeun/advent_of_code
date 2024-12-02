type RecordChecker struct {
	safe       bool
	extra_life bool
	increase   bool
	decrease   bool
	records    [][]int
}

func (checker *RecordChecker) reset() {
	checker.safe = true
	checker.extra_life = true
	checker.increase = false
	checker.decrease = false
}

func (checker *RecordChecker) CountAll() int {
	count := 0
}
