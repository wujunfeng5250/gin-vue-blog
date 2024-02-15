package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001
	ParamsError             = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ParamsError:   "参数错误",
	}
)
