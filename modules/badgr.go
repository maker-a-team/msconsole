package modules

import (
	"fmt"
	// "log"

	"github.com/tempor1s/msconsole/graphql"

	"github.com/imroc/req"

	"github.com/spf13/cobra"
)


// BadgrModule is the source code that allows a user to interact with the Badgr API
func BadgrModule(cmdCtx *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter an API Option")
		return
	}

	// Create a new session
	session := req.New()

	// Log the user in and set cookies :)
	loginUser(session, false)

	// Get the logged in users name and email from Graph QL
	name, email := graphql.UserInfo(session)

	fmt.Printf("\nName: %s\nMS Email: %s\n\n", name, email)

	description := req.Param{
		"username": email,
		"password": args[1],
	}

	// Try to get AccessToken for logged in user
	TokenResp, TokenErr := session.Post(fmt.Sprintf("https://api.badgr.io/o/token"), description)
	if TokenErr != nil {
		// log.Fatal(TokenErr)
		fmt.Print("Error Referencing Badgr Access Token\n", TokenErr)
	} else {
		fmt.Print("Badgr Access Token\n", TokenResp.String())
		// EXAMPLE OUTPUT : Access Badgr Token{"access_token": "oTykfhuTbgTt0uOuj7MVEphpIKRtT4", "expires_in": 86400, "token_type": "Bearer", "scope": "rw:profile rw:issuer rw:backpack", "refresh_token": "bvLV2U4a5fkN5VDR441cHRg9rn2DPe"}
	}
	fmt.Print("\n------\n")

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer Ed0CWwJK1B198OfiI0wE1HAby1v0t5",
	}

	// Make School Badgr Issuer ID:  M4OFTtt7QqiJ87tCGt0a0A
	// Chris Barnes Badgr ID:        Z-s8Mj4pTCa3AVraQxc-4A

	switch args[0] {
	// case "Access-Token":
	// 	description := req.Param{
	// 		"username": email,
	// 		"password": args[1],
	// 	}
	// 	// Try to get AccessToken for logged in user
	// 	TokenResp, TokenErr := session.Post(fmt.Sprintf("https://api.badgr.io/o/token"), description)
	// 	if TokenErr != nil {
	// 		// log.Fatal(TokenErr)
	// 		fmt.Print("Error Referencing Access Token\n", TokenErr)
	// 	} else {
	// 		fmt.Print("Access Token\n", TokenResp.String())
	// 	}
	case "User":
		// Get User Info
		UserResp, UserErr := session.Get(fmt.Sprintf("https://api.badgr.io/v2/users/self"), header)
		if UserErr != nil {
			// log.Fatal(UserErr)
			fmt.Print("Error Referencing USER DATA\n", UserErr)
		} else {
			fmt.Print("USER DATA\n", UserResp.String())
		}
	case "Issuer":
		// GET ISSUER Entity ID    I.E Make School - M4OFTtt7QqiJ87tCGt0a0A
		IssuerResp, IssuerErr := session.Get(fmt.Sprintf("https://api.badgr.io/v2/issuers/M4OFTtt7QqiJ87tCGt0a0A"), header)
		if IssuerErr != nil {
			// log.Fatal(IssuerErr)
			fmt.Print("Error Referencing ISSUER DATA\n", IssuerErr)
		} else {
			fmt.Print("ISSUER DATA\n", IssuerResp.String())
		}
	default:
		fmt.Print("Unknow Option\n")
	}


// {
//   "entityType": "string",
//   "entityId": "string",
//   "firstName": "string",
//   "lastName": "string",
//   "emails": [
//     {
//       "entityType": "string",
//       "entityId": "string",
//       "email": "user@example.com",
//       "verified": true,
//       "primary": true,
//       "caseVariants": "string"
//     }
//   ],
//   "url": "string",
//   "telephone": "string",
//   "agreedTermsVersion": "string",
//   "hasAgreedToLatestTermsVersion": "string",
//   "marketingOptIn": "string",
//   "badgrDomain": "string",
//   "hasPasswordSet": "string",
//   "recipient": "string"
// }


// {
// 	"status":{"success":true,"description":"ok"},
// 	"result":[{
// 		"entityType":"BadgeUser",
// 		"entityId":"Z-s8Mj4pTCa3AVraQxc-4A",
// 		"firstName":"Christopher",
// 		"lastName":"Barnes",
// 		"emails":[
// 			{
// 				"email":"chris.barnes.2000@me.com",
// 				"verified":true,
// 				"primary":true,
// 				"caseVariants":[]
// 			},
// 			{
// 				"email":"christopher.barnes@students.makeschool.com",
// 				"verified":true,
// 				"primary":false,
// 				"caseVariants":[]
// 			}
// 		],
// 		"url":[],
// 		"telephone":[],
// 		"agreedTermsVersion":2,
// 		"hasAgreedToLatestTermsVersion":true,
// 		"marketingOptIn":false,
// 		"badgrDomain":"badgr.com",
// 		"hasPasswordSet":true,
// 		"recipient":{
// 			"type":"email",
// 			"identity":"chris.barnes.2000@me.com"
// 		}
// 	}],
// 	"latestTermsVersion":2
// }


	// Print the new banner message.
	// bannerMessage := getBannerMessage(resp.String())
	// fmt.Print(colorBannerMessage(bannerMessage))
	// fmt.Print(colorBannerMessage(resp.String()))
}

