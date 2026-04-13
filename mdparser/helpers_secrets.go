package mdparser

// AddSecret adds a new secret variable to the mdparser *globally*.
// from now on, the library itself will automatically censor all of the
// "value"s with their name.
// an example of usage would be:
//
//	mdparser.AddSecret(bot.Token, "$TOKEN")
func AddSecret(value, name string) {
	secretMu.Lock()
	defer secretMu.Unlock()

	index := getSecretIndexByValue(value)
	if index != -1 {
		secrets[index].name = name
		return
	}

	secrets = append(secrets, secretContainer{
		value: value,
		name:  name,
	})
}

func RemoveSecretByValue(value string) {
	secretMu.Lock()
	defer secretMu.Unlock()

	index := getSecretIndexByValue(value)
	if index != -1 {
		secrets = append(secrets[:index], secrets[index+1:]...)
	}
}

func RemoveSecretByName(name string) {
	secretMu.Lock()
	defer secretMu.Unlock()

	var newSecrets []secretContainer
	for _, current := range secrets {
		if current.name != name {
			newSecrets = append(newSecrets, current)
		}
	}

	secrets = newSecrets
}

func GetSecretIndexByValue(value string) int {
	secretMu.RLock()
	defer secretMu.RUnlock()

	return getSecretIndexByValue(value)
}

func getSecretIndexByValue(value string) int {
	for index, current := range secrets {
		if current.value == value {
			return index
		}
	}

	return -1
}

func SecretValueExists(value string) bool {
	secretMu.RLock()
	defer secretMu.RUnlock()

	return getSecretIndexByValue(value) != -1
}
