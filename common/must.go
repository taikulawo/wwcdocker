package common


// Must2 panics if the second parameter is not nil, Otherwise returns the first parameter
func Must2(v interface{}, err error ) interface{}{
	Must(err)
	return v
}

// Must panics if err is not nil
func Must(err error) {
	if err !=  nil {
		panic(err)
	}
}