package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"m-macdonald/gog-backup/gog/api"
)

//
// type Flag int
//
// const (
//     Code Flag = iota
// )
//
// var MapStringToFlag = func() map[string]Flag {
//     m := make(map[string]Flag)
//     for i := Code; i <= Code; i++ {
//         m[i.String()] = i
//     }
//
//     return m
// }()
//
// var MapFlagToString = func() map[Flag]string {
//
// }

const code = "code"

// var (
//     flagsMap
// )

// const flags = {
//     "code": Code
// }


var flags = []cli.Flag {
    &cli.StringFlag{Name: code, Aliases: []string{"c"}},
}


func main() {
    app := &cli.App {
        Name: "GOG Backup",
        Usage: "",
        Flags: flags,
        Action: cliEntryPoint,
    }
    
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

// https://auth.gog.com/auth?client_id=46899977096215655&redirect_uri=https%3A%2F%2Fembed.gog.com%2Fon_login_success%3Forigin%3Dclient&response_type=code&layout=client2

func cliEntryPoint(cCtx *cli.Context) error {
    authCode := cCtx.String(code)
    if (len(authCode) > 0) {
        token, err := api.GetToken("authorization_code", authCode, "https://embed.gog.com/on_login_success?origin=client", "")
        fmt.Printf("%+v", token)

        api.GetUsersGames(token.AccessToken)
    } else {
        println("'code' flag was not set. This is not recommended. Attempting browser auth flow")
    }

    return nil
}
