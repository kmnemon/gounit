package wrap_method03_exposewrap

type Employee struct {
	payRate   float32
	timecards []Timecard
}

func (e *Employee) makeLoggedPayment() {
	e.logPayment()
	e.pay()
}

func (e *Employee) pay() {
	var amount Money
	for _, card := range e.timecards {
		amount.add(float32(card.getHours()) * e.payRate)
	}

	dispatcherPay(e, amount)
}

func (e *Employee) logPayment() {
	//...
}

func dispatcherPay(e *Employee, amount Money) {
	//...
}

//原pay方法未被客户端调用
