package getHint

import "testing"

func TestGetHint(t *testing.T)  {
	secret := "1807"
	guess := "7810"
	info := getHint(secret,guess)
	t.Log(info)
}