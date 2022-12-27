 package iteration


 func Repeat(char string) string {
	var newCharacters string
	for i := 0; i < 5; i++ {
		newCharacters = newCharacters + char
	}
	return newCharacters
 }