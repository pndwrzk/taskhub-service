package error

func ErrMapping(err error) bool {

	allErrors := [][]error{GeneralError, UserErrors, TaskErrors}

	for _, group := range allErrors {
		for _, item := range group {
			if item != nil && item.Error() == err.Error() {
				return true
			}
		}
	}

	return false
}
