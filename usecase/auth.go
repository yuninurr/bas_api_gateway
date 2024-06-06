// package auth

// // struct login
// type Login struct {
// 	Username string
// 	password string
// }

// // type login interface 
// type LoginInterface interface {
// 	Authentication(username, password string) bool
// }

// // buat fungsi new login dengan definisi return type LoginInterface 
// func (a *Login) Authentication(username, password string) bool {
//     // Contoh sederhana menggunakan hardcoded username dan password
//     if username == "admin" && password == "admin123" {
//         return true
//     }
//     return false
// }


// // // NewLogin mengembalikan instance dari Login yang mengimplementasikan LoginInterface
// func NewLogin() LoginInterface {
//     return &Login{}
// }



package usecase

type login struct {
}

type LoginInterface interface {
	Authentifikasi(username, password string) bool
}

func (a *login) Authentifikasi(username, password string) bool {
	if username == "admin" && password == "admin123" {
		return true
	}
	return false
}

func NewLogin() LoginInterface {
	return &login{}
}