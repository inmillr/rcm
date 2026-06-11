package validate

func Row(row []string) (bool, string) {
	if len(row) < 4 {
		return false, "too few columns"
	}
	if row[0] == "" {
		return false, "missing PATID"
	}
	if row[3] == "" {
		return false, "missing amount"
	}
	return true, ""
}
