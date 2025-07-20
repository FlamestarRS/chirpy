# Welcome to Chirpy!

Chirpy allows users to post "chirps" of 140 characters or less about anything they want.

Requests must have appropriate headers and JSON formatted body content as detailed in ENDPOINTS below.

## Getting Started

1. Clone the repository
2. Install Go and PostgreSQL
3. Run database migrations with Goose
4. Start the server: `go run .`
5. Create a user, login, and start chirping!

## --- ENDPOINTS ---

### POST:

/admin/reset        - Reset user and chirp databases and all metrics **WARNING: DEV ONLY**

/api/chirps         - Post a chirp
    Required Header: "Authorization: Bearer ${userRefreshToken}"
    Example Request Body:
    {
        "body": "Let’s just say I know a guy... who knows a guy... who knows another guy."
    }

/api/users          - Create a new user
    Example Request Body:
    {
        "email": "saul@bettercall.com",
        "password": "123456"
    }

/api/login          - Log into an existing user, issues access and refresh token
    Example Request Body:
    {
        "email": "saul@bettercall.com",
        "password": "123456"
    }

/api/refresh        - Create new access token using valid refresh token
    Required Header: "Authorization: Bearer ${userRefreshToken}"

/api/revoke         - Revoke refresh token
    Required Header: "Authorization: Bearer ${userRefreshToken}"
    
/api/polka/webhooks - Third party endpoint used to upgrade user to ChirpyRed
    Required Header: "Authorization: ApiKey ${ApiKey}"


### GET:

/api/healthz          - Confirm site is online
/admin/metrics        - Number of times the site has been visited
/api/chirps           - Retrieve all chirps. Optional query params "author_id={author_id}", "sort={asc || desc}"
/api/chirps/{chirpID} - Retrieve specific chirp by ID

### PUT:

/api/users - Update user email or password
    Required Header: "Authorization: Bearer ${userAccessToken}"
    Example Body:
    {
        "email": "walter@breakingbad.com",
        "password": "losPollosHermanos"
    }

### DELETE:

/api/chirps/{chirpID} - Delete a chirp
    Required Header: "Authorization: Bearer ${userAccessToken}"