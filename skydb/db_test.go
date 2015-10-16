package skydb

import (
	"testing"
)

type fakeConn struct {
	Conn
	AppName      string
	OptionString string
}

type fakeDriver struct {
	Driver
}

func (driver fakeDriver) Open(appName, optionString string) (Conn, error) {
	return fakeConn{
		AppName:      appName,
		OptionString: optionString,
	}, nil
}

func TestOpen(t *testing.T) {
	defer unregisterAllDrivers()

	Register("fakeImpl", fakeDriver{})

	if driver, err := Open("fakeImpl", "com.example.app.test", "fakeOption"); err != nil {
		t.Fatalf("got err: %v, want a driver", err.Error())
	} else {
		if driver, ok := driver.(fakeConn); !ok {
			t.Fatalf("got driver = %v, want a driver of type fakeDriver", driver)
		} else {
			if driver.AppName != "com.example.app.test" {
				t.Fatalf("got driver.AppName = %v, want \"com.example.app.test\"", driver.AppName)
			}
			if driver.OptionString != "fakeOption" {
				t.Fatalf("got driver.OptionString = %v, want \"fakeOption\"", driver.OptionString)
			}
		}
	}
}