package helpers

func CheckExists(units []Unit, unit_id string) int {
	exists := -1
	for i, unit := range units {
		if unit.ID == unit_id {
			exists = i
		}
	}
	return exists
}
