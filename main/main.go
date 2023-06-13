package main

func main() {

	// NOTE - should not change the order of the calling methods
	BuildEnv()

	StartDb()

	BuildValidator()

	BuildDomain()

	StartApi()
}
