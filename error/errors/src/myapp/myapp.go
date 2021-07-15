package myapp

type MyApp struct {
}

func New() (t *MyApp) {
	t = new(MyApp)
	return t
}

func (t *MyApp) Hello() string {
	return "Hello World!"
}

func (t *MyApp) Error0() error {
	return &AppError{Code: ERR_zero, Msg: "Error Zero"}
}

func (t *MyApp) Error1() error {
	return &AppError{Code: ERR_one, Msg: "Error One"}
}

func (t *MyApp) Error2() error {
	return &AppError{Code: ERR_two, Msg: "Error Two"}
}

func (t *MyApp) Error3() error {
	return &AppError{Code: ERR_three, Msg: "Error Three"}
}
