package logging

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/cloud-barista/cb-apigw/restapigw/pkg/core"
	cblog "github.com/cloud-barista/cb-log"
	"github.com/sirupsen/logrus"
)

// ===== [ Constants and Variables ] =====

// ===== [ Types ] =====

// Logger - CB-LOG에서 사용하는 "logrus" Logger를 위한 Wrapper 구조
type Logger struct {
	*logrus.Logger
}

// ===== [ Implementations ] =====

// SetOutput - 로그 출력기 설정
func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.Out = w
}

// DisableOutput - 로그 출력 비활성화
func (l *Logger) DisableOutput() {
	l.SetOutput(ioutil.Discard)
}

// SetFormatter - 로그 포맷터 설정
func (l *Logger) SetFormatter(f logrus.Formatter) {
	l.Logger.Formatter = f
}

// SetLogLevel - 로그 레벨 설정
func (l *Logger) SetLogLevel(lv logrus.Level) {
	l.Logger.SetLevel(lv)
}

// ===== [ Private Functions ] =====

// ===== [ Public Functions ] =====

// NewLogger - 초기화된 Logger의 인스턴스 생성
func NewLogger() *Logger {
	// FIXME: CBLOG 경로관련 문제 (현재 경로로 환경변수 설정)
	if dir, err := os.Getwd(); err == nil {
		os.Setenv("CBLOG_ROOT", dir)
	}

	return &Logger{
		Logger: cblog.GetLogger(core.AppName),
	}
}
