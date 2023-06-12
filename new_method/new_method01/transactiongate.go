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
