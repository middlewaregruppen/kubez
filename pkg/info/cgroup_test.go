package info

import "testing"

func TestGetIntFromFile_1(t *testing.T) {
	v, err := getIntFromFile("./testdata/cpu.cfs_quota_us")
	if err != nil {
		t.Logf("error returned: %s ", err)
		t.Fail()
	}
	if v != 1337 {
		t.Log("value of integer is not 1337")
		t.Fail()
	}
}

func TestGetIntFromFile_2(t *testing.T) {
	v, err := getIntFromFile("./testdata/not_an_integer")
	if err == nil {
		t.Logf("Expected an error")
		t.Fail()
	}
	if v != 0 {
		t.Log("value of integer is not 0")
		t.Fail()
	}
}

func TestGetIntMapFromFile(t *testing.T) {
	m, _ := getIntMapFromFile("./testdata/cpu.stat")
	if m["nr_throttled"] != 6 {
		t.Log("value of map value is not 6")
		t.Fail()
	}

}

func TestGetIntMapFromFile_2(t *testing.T) {
	_, err := getIntMapFromFile("./testdata/not_exists")
	if err == nil {
		t.Logf("Expected an error")
		t.Fail()
	}

}

func TestGetIntArrayFromFile(t *testing.T) {
	a, err := getIntArrayFromFile("./testdata/cpuacct.usage_percpu")
	if err != nil {
		t.Logf("error returned: %s ", err)
		t.Fail()
	}

	if a[1] != 1533472003 {
		t.Log("value is not 1533472003")
		t.Fail()
	}

}
