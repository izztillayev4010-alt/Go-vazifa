package handler

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"todoapp/cache"
	"todoapp/db"
	"todoapp/helpers"
	"todoapp/src/models"
)

func printEndpoints(e map[int]string) {
	for k, v := range e {
		fmt.Println(k, " | ", v)
	}
}

func Api() {
	var endpoints = map[int]string{
		1: "/singup",
		2: "/login",
		3: "/users",
		4: "/logout",
		5: "/who-am-i",
	}

	printEndpoints(endpoints)

	fmt.Println("App is running on the console")

	for {
		var endpoint int

		fmt.Scanln(&endpoint)

		fmt.Println(endpoints[endpoint])

		switch endpoint {
		case 1:
			ClearTerminal()
			singUp()
		case 2:
			Login()

		case 3:
			ClearTerminal()
			Users()
		case 4:
			Logout()
		case 5:
			if helpers.Token == nil {
				fmt.Println("Siz ro'yxatdan o'tmagansiz!")
				continue
			}
			whoami()

		}
	}

}

func singUp() {
	var newUser models.UserCreate

	fmt.Print("Gmailingizni kiriting: ")
	fmt.Scanln(&newUser.Gmail)

	for {
		var (
			firstCreate  string
			secondCreate string
		)

		fmt.Print("Parol yarating: ")
		fmt.Scanln(&firstCreate)

		if len(firstCreate) < 4 {
			fmt.Println("Parol juda qisqa! Uzunroq parol kiriting!")
			continue
		}

		fmt.Print("Parolni qayta kiriting: ")
		fmt.Scanln(&secondCreate)

		if firstCreate == secondCreate {
			newUser.Password = firstCreate
			break
		} else {
			fmt.Println("Parol mos kelmadi!  Qayta kiriting!")
		}
	}

	fmt.Println(newUser)
	ClearTerminal()

	otp := helpers.GenerateOTP()

	cache.Cache[newUser.Gmail] = cache.CacheData{
		OtpCode: otp,
		NewUser: newUser,
	}

	helpers.SendOtpGmail(otp, newUser.Gmail)

	fmt.Print("Gmailingizga kod yuborildi! Kodni bu yerga kiriting: ")
	var kod string
	fmt.Scanln(&kod)

	if kod != otp {
		fmt.Println("kodni xato kiritdingiz!")
		Api()
		return
	}

	regUser := db.UserTable{Id: uint64(len(db.Users)) + 1, Gmail: newUser.Gmail, Password: newUser.Password}

	db.Users = append(db.Users, regUser)

	delete(cache.Cache, newUser.Gmail)
	ClearTerminal()

	var authInfo helpers.Authourizeid = helpers.Authourizeid{
		Id:    int(regUser.Id),
		Gmail: regUser.Gmail,
	}

	helpers.Token = &authInfo

	fmt.Println("Siz muvofaqqiyatli ro'yxatdan o'tdingiz!")
	Api()
}

func Users() {
	for _, v := range db.Users {
		fmt.Println(v)
	}
	Api()
}

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")

	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()

}

func whoami() {
	for _, v := range db.Users {
		if v.Id == uint64(helpers.Token.Id) {
			fmt.Println(v.Id)

			fmt.Println(v.Gmail)

			fmt.Println(v.Password)
			break
		}
	}
}

func Logout() {
	helpers.Token = nil
}

func Login() {

	var (
		loginUser models.UserCreate
	)
	fmt.Print("Ro'yxatdan o'tgan gmailingizni kiriting: ")
	fmt.Scanln(&loginUser.Gmail)

	fmt.Print("Parolingizni kiriting: ")
	fmt.Scanln(&loginUser.Password)

	for _, v := range db.Users {
		if v.Gmail == loginUser.Gmail && v.Password == loginUser.Password {
			var authInfo helpers.Authourizeid = helpers.Authourizeid{
				Id:    int(v.Id),
				Gmail: v.Gmail,
			}
			helpers.Token = &authInfo

		}
	}
}
