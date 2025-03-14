package logger

//"DEV" or "PROD"
var logLevel string

func SetLogLevel(level string) {
	logLevel = level
}

func DebugLog() {
	if logLevel != "PROD" {

	}
}

func ErrorLog() {

}

func InfoLog() {
	if logLevel != "PROD" {

	}
}

func CriticalLog() {

}
