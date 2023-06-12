package new_method02_oneway

type TransactionGate struct {
	transactionBundle TransactionBundle
}

func (t *TransactionGate) postEntries(entries []Entry) {
	entriesToAdd := make([]Entry, 0)
	for _, entry := range entries {
		if !t.transactionBundle.hasEntry(entry) {
			entry.postDate()
			entriesToAdd = append(entriesToAdd, entry)
		}

	}

	t.transactionBundle.add(entriesToAdd)
}
