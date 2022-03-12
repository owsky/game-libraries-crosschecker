package main

import "errors"

// binary search within the whole steam library
func Contains(apps []App, name string) (App, error) {
	left := 0
	right := len(apps)

	for left <= right {
		mid := (left + right) / 2

		if apps[mid].Name == name {
			return apps[mid], nil
		} else if apps[mid].Name < name {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	err := errors.New("App not found")
	return App{}, err
}
