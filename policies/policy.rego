package authz

default allow = false

allow {
	some i, j, k
	groups := get_groups(input.user)
	roles := get_group_roles(input.resource, groups[i])
	actions := get_role_actions(roles[j])
	actions[k] == input.action
}

get_groups(user) = groups {
	groups := data.users[user]
}

get_group_roles(resource, group) = roles {
	roles := data.resources[resource][group]
}

get_role_actions(role) = actions {
	actions := data.roles[role]
}
