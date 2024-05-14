package api

import "golang.org/x/crypto/bcrypt"

var usersPasswords = map[string][]byte{
	"joe":  []byte("$2a$12$aMfFQpGSiPiYkekov7LOsu63pZFaWzmlfm1T8lvG6JFj2Bh4SZPWS"),
	"mary": []byte("$2a$12$l398tX477zeEBP6Se0mAv.ZLR8.LZZehuDgbtw2yoQeMjIyCNCsRW"),
}

func LoginU(l string, p string) bool {
	wantPass, hasUser := usersPasswords[l]
	if !hasUser {
		return false
	}
	if cmperr := bcrypt.CompareHashAndPassword(wantPass, []byte(p)); cmperr == nil {
		return true
	}
	return false
}
