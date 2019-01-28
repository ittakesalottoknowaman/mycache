package mycache

func cronExpirationCheck() {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	for dbName := range cache.table {
		table := cache.getTable(dbName)
		table.mutex.RLock()
		for key, item := range table.items {
			if item.isExpire() {
				table.mutex.RUnlock()
				table.Delete(key)
				table.mutex.RLock()
			}
		}
	}
}
