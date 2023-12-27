package tools

type UserMock struct{}

func (u *UserMock) HomeDir() string {
	return "/root/"
}

func (u *UserMock) IsRoot() bool {
	return true
}
