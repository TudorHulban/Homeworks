package main

func contains(source, what []string) bool {
	for _, elementWhat := range what {
		var elementIsContained bool

		for _, elementSource := range source {
			if elementSource == elementWhat {
				elementIsContained = true

				break
			}
		}

		if !elementIsContained {
			return false
		}
	}

	return true
}
