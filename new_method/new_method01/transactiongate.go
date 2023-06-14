package new_method01

type TransactionGate struct {
	transactionBundle TransactionBundle
}

func (t *TransactionGate) postEntries(entries []Entry) {
	for _, entry := range entries {
		entry.postDate()
	}

	t.transactionBundle.add(entries)
}

//TO DO： 检查entries中的对象在日期被发送并添加到transactionBundle中去之前是否已经存在
