# Goly Shortener API

This API helps users to shorten their url links

## Features include
- Generation of shortened link
- Customization of shortened link
- Generation of QR CODE
- link analytics  i.e number of times the link has been visited
- User history, etc.

**BASE URL** - ***https://goly.onrender.com***  
**Documentation** - ***https://goly.stoplight.io/docs/goly/branches/main/q6wza6ie16atg-goly-shortener-api*** 

## How it works
A user can access all the features of the API by simply signing up and then
proceed to sign in. The application can still be used without signing in 
but in such a scenario the user is unknown. 

### An unknown user will only be allowed to 
- Generation of shortened link
- Customization of shortened link
- Generation of QR CODE
- An unknown user will not have access to user history and link analytics

## Enpoints
| Method | Endpoint | Description | Requires Authentication? |  
| :----- | :------- | :---------- | :----------------------- |  
| GET | **/** | Homepage | NO |  
| POST | **/users/signup** | sign up | NO |  
| POST | **/users/signin** | sign in | NO |  
| GET | **/users/:id/history** | user history | YES |  
| GET | **/golies** | get all public golies | NO |  
| POST | **/golies** | create new goly | Optional auth |  
| GET | **/golies/:id** | get goly by id | Optional auth |  
| PUT | **/golies/:id** | Modify goly | YES |  
| DELETE | **/golies/:id** | delete goly | YES |  
| GET | **/r/:redirect** | Redirects user to the original url i.e long link | NO auth |  
| GET | **/:redirect/generateQRCode** | generates QR CODE | NO auth |  


## Definitions
* User: A registered user who has signed into their account on our website using email address, password
* Unknown User: A user who has not signed in.
* Optional Auth: Resources with optional auth can be accessed by a User and Unknown user, but the 
experience of both will be different i.e a User will be able to see more information on golies and 
have their User ID stored when they create a goly. 
* Required Auth: These resources require authentication before accessing them e.g creating or modifying an existing resource
* Public Golies: These are golies which the author created and made public i.e public = true. Authors can chose to 
make their golies private i.e public = false. Private golies can only be viewed by the creator

# GOLY
 Goly is the name for the user object created, as well as the name of the shrotened URL. The
 user object goly contains the following
- **Redirect** - It is a string that represents the long url to be shortened i.e original url
-	**Goly** - This is the string of the  shortened link. It can be customized or randomly generated. 
To customize the link Custom must be set to true i.e custom = true, the Goly string must contain 
at least 10 characters. If it is randomly generated custom is set to false i.e Custom = false.
-	**Clicked** -  An integer that tracks/shows the number of times a link/goly was clicked i.e used
-	**Custom** - This determines if the Goly(shortened string) is randomly generated or customized by 
the user. To customize set custom to true i.e custom = true, otherwise set custom = false. 
-	**Public** - Public golies can be viewed by anyone, while private golies can only be accessed by the 
user. To make golies private set public = false
-	**UserID** - An integer that stores the ID of the user if the user has been authenticated.




