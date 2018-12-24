package mysql

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_Mysql(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	Demo := MysqlAgent{User: "root", Password: "SMdemT2Pm",
		Host: "172.18.8.88", Port: "60333", DBName: "Demo2"}

	// db, err := Demo.InitClient()
	// if err != nil {
	// 	logrus.Errorln(err)
	// }

	s, _ := Demo.Get("asdf")
	logrus.Debugln(s)

	Demo.Set("md5sum", "https://tangxin-test-02.oss-cn-hangzhou.aliyuncs.com/v6/naruto.png")
	s, _ = Demo.Get("md5sum")
	logrus.Debugln(s)
}
