package main

import (
	auth "api_gateway/usecase"
	"fmt"
)

func main() {
	
	login := auth.NewLogin()

	username := "admin"
	password := "admin123"

	isAuthenticated := login.Authentication(username, password)

	if isAuthenticated {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Invalid username or password")
	}
}


























// package main

// import (
//     "encoding/json"
//     "fmt"
//     "net/http"
//     "project/usecase" // Pastikan path ini sesuai dengan struktur project Anda
// )

// // LoginRequest struct untuk menyimpan data login
// type LoginRequest struct {
//     Username string `json:"username"`
//     Password string `json:"password"`
// }

// // LoginResponse struct untuk respon login
// type LoginResponse struct {
//     Message string `json:"message"`
//     Success bool   `json:"success"`
// }

// func main() {
//     http.HandleFunc("/login", loginHandler)
//     fmt.Println("Server is running on port 8080")
//     http.ListenAndServe(":8080", nil)
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
//     if r.Method != http.MethodPost {
//         http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//         return
//     }

//     var loginReq LoginRequest
//     err := json.NewDecoder(r.Body).Decode(&loginReq)
//     if err != nil {
//         http.Error(w, "Invalid request body", http.StatusBadRequest)
//         return
//     }

//     loginService := usecase.NewLogin()
//     if loginService.Authenticate(loginReq.Username, loginReq.Password) {
//         response := LoginResponse{
//             Message: "Login successful",
//             Success: true,
//         }
//         w.Header().Set("Content-Type", "application/json")
//         json.NewEncoder(w).Encode(response)
//     } else {
//         response := LoginResponse{
//             Message: "Invalid username or password",
//             Success: false,
//         }
//         w.Header().Set("Content-Type", "application/json")
//         json.NewEncoder(w).Encode(response)
//     }
// }
