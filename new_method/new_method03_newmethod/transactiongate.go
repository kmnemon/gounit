package new_method01_newmethod

type TransactionGate struct {
	transactionBundle TransactionBundle
}

func (t *TransactionGate) postEntries(entries []Entry) {
	entriesToAdd := t.uniqueEntries(entries)
	for _, entry := range entries {
		entry.postDate()
	}

	t.transactionBundle.add(entriesToAdd)
}

func (t *TransactionGate) uniqueEntries(entries []Entry) []Entry {
	//filter entries
	return entries
}
