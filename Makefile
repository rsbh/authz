build:
	go build -o bin/authz

test-policy:
	opa test . -v