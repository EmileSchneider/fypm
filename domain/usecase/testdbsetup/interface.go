package testdbsetup

type usecase interface{
	Setup()
	Reset()
}
