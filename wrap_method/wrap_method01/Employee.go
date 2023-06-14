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

//新增需求：每次我们给一个雇员支付薪水时都做一下日志记录

func dispatcherPay(e *Employee, amount Money) {

}
