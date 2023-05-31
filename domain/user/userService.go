package user

// Get finds a simple by ID
func Get(name string) (User, error) {
	return GetOne(name)
}

func Getrows(offset int, limit int) ([]User, error) {
	return Getpage(offset, limit)
}

func Post(user *User) (User, error) {
	return PostOne(user)
}
