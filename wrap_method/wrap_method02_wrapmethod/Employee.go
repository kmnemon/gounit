package wrapmethod01_wrapmethod

type Employee struct {
	payRate   float32
	timecards []Timecard
}

func (e *Employee) dispatchPayment() {
	var amount Money
	for _, card := range e.timecards {
		amount.add(float32(card.getHours()) * e.payRate)
	}

	dispatcherPay(e, amount)
}

func (e *Employee) pay() {
	e.logPayment()
	e.dispatchPayment()
}

func (e *Employee) logPayment() {
	//...
}

func dispatcherPay(e *Employee, amount Money) {
	//...
}
