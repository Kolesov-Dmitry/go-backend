package group

import "context"

type GroupRepo interface {
	// AddGroup creates new group
	// Inputs:
	//   g - new group data
	// Output:
	//   returns group with its ID in case of success, otherwise returns error
	AddGroup(ctx context.Context, g *Group) (*Group, error)

	// AppendUserToGroups appends user to the groups
	// Inputs:
	//   userName - name of the user to be appended to the groups
	//   groups - group names
	// Output:
	//   returns error if failed
	AppendUserToGroups(ctx context.Context, userName string, groups []string) error

	// RemoveUserFromGroups removes user from the groups
	// Inputs:
	//   userName - name of the user to be removed from the groups
	//   groups - group names
	// Output:
	//   returns error if failed
	RemoveUserFromGroups(ctx context.Context, userName string, groups []string) error

	// FindGroupByName searches group by provided name
	// Inputs:
	//   groupName - name of the group
	// Output:
	//   - found group if succeeded
	//   - nil if no group was found
	//   - error if failed
	FindGroupByName(ctx context.Context, groupName string) (*Group, error)

	// FindGroupByUsers searches groups by provided user names
	// Inputs:
	//   users - list of user names
	// Output:
	//   - found groups if succeeded
	//   - empty list if no groups were found
	//   - error if failed
	FindGroupByUsers(ctx context.Context, users []string) ([]*Group, error)
}
