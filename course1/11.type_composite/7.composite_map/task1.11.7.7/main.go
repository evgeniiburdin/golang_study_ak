package main

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	uniqueUsers := make([]User, 0, len(users))
	nicknamesUnique := make(map[string]struct{})
	for _, user := range users {
		if _, ok := nicknamesUnique[user.Nickname]; !ok {
			nicknamesUnique[user.Nickname] = struct{}{}
			uniqueUsers = append(uniqueUsers, user)
		}
	}
	return uniqueUsers
}
