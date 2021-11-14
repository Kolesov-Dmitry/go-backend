package user

import "context"

type UserRepo interface {
	// AddUser creates new user
	// Inputs:
	//   u - new user data
	// Output:
	//   returns user with its ID in case of success, otherwise returns error
	AddUser(ctx context.Context, u *User) (*User, error)

	// FindUserByName searches user by provided name
	// Inputs:
	//   userName - name of the user
	// Output:
	//   - found user if succeeded
	//   - nil if no user was found
	//   - error if failed
	FindUserByName(ctx context.Context, userName string) (*User, error)

	// FindUserByGroup searches users by provided group name
	// Inputs:
	//   groupName - name of the group
	// Output:
	//   - found users if succeeded
	//   - empty list if no users were found
	//   - error if failed
	FindUserByGroup(ctx context.Context, groupName string) ([]*User, error)
}
