// GOLY SHORTENER API.
//
// # This API helps users to shorten their url links
//
// It has various features which include;
// Generation of shortened link
// Customization of shortened link
// Generation of QR CODE
// link analytics  i.e number of times the link has been visited
// User history, etc.
//
//	Schemes: http
//	Host: localhost:8001
//	BasePath: /
//	Version: 1.0.1
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Clinton Mekwunye<Mekwunyeclinton22@gmail.com> https://github.com/clintonMF
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"time"
)

// Goly Model
//
// Each goly contains all the details of the link.
//
// swagger:model Goly
type Goly struct {
	// the link to be shortened
	// required: true
	// min length: 3
	// example: http://go.dev
	Redirect string `json:"redirect"`

	// a unique identifier that is used in place of the original url
	// required: false
	// min length: 10
	// example: http://localhost:8001/r/IPs2yW2p
	Goly string `json:"goly" gorm:"unique;not null"`

	// This is used to determine if the user wants a customized link or a randomly generated shortened link
	// required: true
	// example: true
	Custom bool `json:"custom"`

	// This indicates if the shortened link is private or public i.e only the user can access it.
	// required: false
	// example: false
	Public bool `json:"public" gorm:"default:false"`
}

// User
//
// The author of links/Golies
//
// swagger:model User
type User struct {
	// the name of the user
	// required: true
	// min length: 3
	// example: Yuta
	Name string `json:"name" binding:"required"`

	// the email of the user
	// required: true
	// min length: 5
	// format: email
	// example: TheCursedChild@JJK.com
	Email string `json:"email" gorm:"unique;not null" binding:"required"`

	// the email of the student
	// required: true
	// min length: 8
	// format: string
	// example: son44THz
	Password string `json:"password" gorm:"not null" binding:"required"`
}

// UserLogin
//
// The author of links/Golies
//
// swagger:model UserLogin
type UserLogin struct {
	// the name of the user
	// required: true
	// min length: 3
	// example: Yuta
	Name string `json:"name" binding:"required"`

	// the email of the user
	// required: true
	// min length: 5
	// format: email
	// example: TheCursedChild@JJK.com
	Email string `json:"email" gorm:"unique;not null" binding:"required"`
}

// GolyResponse
//
// # Response gotten from getting a GOly by ID
//
// swagger:model GolyResponse
type GolyResponse struct {
	// Goly ID
	// example: 1
	ID uint `json:"id"`

	// Redirect link
	// example: http://go.dev
	Redirect string `json:"redirect"`

	// Shortened URL
	// example: http://localhost:8001/r/IPs2yW2p
	Goly string `json:"goly" gorm:"unique;not null"`

	// Custom
	// example: true
	Custom bool `json:"custom"`

	// Public
	// example: false
	Public bool `json:"public" gorm:"default:false"`

	// Creators ID
	// example: 1
	UserID uint `json:"userId"`

	// Clicked
	// example: 4
	Clicked uint64 `json:"clicked"`

	// Course creation time
	// example: 2022-06-01T12:30:00Z
	CreatedAt time.Time `json:"created_at"`

	// Course last update time
	// example: 2022-07-12T08:45:00Z
	UpdatedAt time.Time `json:"updated_at"`

	// QR code
	// example: http://localhost:8001/r/IPs2yW2p/generateQRCode
	QRCode string
}

// PublicGolyResponse
//
// response gotten when a public goly is accessed by
// anyone who is not the author
//
// swagger:model PublicGolyResponse
type PublicGolyResponse struct {
	// Redirect link
	// example: http://go.dev
	Redirect string `json:"redirect"`

	// Shortened URL
	// example: http://localhost:8001/r/IPs2yW2p
	Goly string `json:"goly" gorm:"unique;not null"`

	// QR code
	// example: http://localhost:8001/r/IPs2yW2p/generateQRCode
	QRCode string
}

// GoliesResponse
//
// # Response generated when a user accesses his golies history
//
// swagger:model GoliesResponse
type GoliesResponse struct {
	// List of golies
	// example:
	// {
	// 	"Golies": [
	// 		{
	// 			"ID": 1,
	// 			"CreatedAt": "2023-06-26T15:10:07Z",
	// 			"UpdatedAt": "2023-06-26T15:22:04Z",
	// 			"DeletedAt": null,
	// 			"redirect": "http://go.dev",
	// 			"goly": "http://localhost:8001/r/OerZMapl",
	// 			"clicked": 1,
	// 			"custom": false,
	// 			"public": true,
	// 			"userId": 1
	// 		},
	// 		{
	// 			"ID": 2,
	// 			"CreatedAt": "2023-06-26T15:31:59Z",
	// 			"UpdatedAt": "2023-06-26T15:31:59Z",
	// 			"DeletedAt": null,
	// 			"redirect": "http://go.dev",
	// 			"goly": "http://localhost:8001/r/ztgPsGvU",
	// 			"clicked": 0,
	// 			"custom": false,
	// 			"public": true,
	// 			"userId": 1
	// 		}
	// 	],
	// 	"number of redirects": 3,
	// 	"status": "success"
	// }
	Golies []GolyResponse
}

// PublicGoliesResponse
//
// # Response generated when an unknown user opens the home page
//
// swagger:model PublicGoliesReponse
type PublicGoliesReponse struct {
	// List of PublicGolies
	// example:
	// {
	// 	"Golies": [
	// 		{
	// 			"redirect": "http://go.dev",
	// 			"goly": "http://localhost:8001/r/OerZMapl",
	// 		},
	// 		{
	// 			"redirect": "http://go.dev",
	// 			"goly": "http://localhost:8001/r/ztgPsGvU",
	// 		}
	// 	],
	// 	"number of redirects": 2,
	// 	"status": "success"
	// }
	PublicGolies []PublicGolyResponse
}

// User response
//
// # This response is generated when a user signs up
//
// swagger:model UserResponse
type UserResponse struct {
	// user ID
	// example: 1
	ID uint `json:"id"`
	// name
	// example: Goku
	FirstName string `json:"name"`
	// email
	// example: supersayan@DBZ.com
	Email string `json:"email"`
	// password
	// example: Goku
	Password string `json:"Son44Goku"`
	// user creation time
	// example: 2022-06-01T12:30:00Z
	CreatedAt time.Time `json:"created_at"`
	// user last update time
	// example: 2022-07-12T08:45:00Z
	UpdatedAt time.Time `json:"updated_at"`
}
