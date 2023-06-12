package wrapmethod01

type Employee struct {
	payRate   float32
	timecards []Timecard
}

func (e *Employee) pay() {
	var amount Money
	for _, card := range e.timecards {
		amount.add(float32(card.getHours()) * e.payRate)
	}

	dispatcherPay(e, amount)
}

func dispatcherPay(e *Employee, amount Money) {

}
