package authz

data = {
	"users": {
		"u1": ["g1", "g2", "g3"],
		"u3": ["g3"],
	},
	"resources": {
		"r1": {
			"g1": ["reader"],
			"g2": ["writer"],
		},
		"r2": {
			"g1": ["reader"],
			"g2": ["reader"],
		},
	},
	"roles": {
		"writer": ["read", "write"],
		"reader": ["read"],
	},
}

test_user_not_found_denied {
	not allow with input as {"user": "u404"}
}

test_resource_not_found_denied {
	not allow with input as {"user": "u1", "resource": "r404"} with data as data
}

test_action_not_found_denied {
	not allow with input as {"user": "u1", "resource": "r1", "action": "a404"} with data as data
}

test_user_not_member_of_group_denied {
	not allow with input as {"user": "u3", "resource": "r1", "action": "read"} with data as data
}

test_group_doesnt_have_role_denied {
	not allow with input as {"user": "u1", "resource": "r2", "action": "write"} with data as data
}

test_group_have_read_permission_allowed {
	allow with input as {"user": "u1", "resource": "r2", "action": "read"} with data as data
}

test_group_have_read_permission_allowed {
	allow with input as {"user": "u1", "resource": "r1", "action": "read"} with data as data
}

test_group_have_write_permission_allowed {
	allow with input as {"user": "u1", "resource": "r1", "action": "write"} with data as data
}
