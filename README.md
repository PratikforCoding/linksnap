# LinkSnap  

LinkSnap is a URL shortening service built using Go, Gin, and MongoDB. It allows users to create short URLs, view their details, track analytics, and generate QR codes for easy sharing.  

## Features  
- Create short URLs with optional custom aliases  
- List all short URLs  
- Retrieve detailed information about a specific URL  
- Delete short URLs  
- View analytics and usage statistics  
- Generate QR codes for short URLs  

---

## API Endpoints  

### URL Operations  
| **Endpoint**                  | **Method** | **Description**                  |  
|-------------------------------|------------|----------------------------------|  
| `/api/urls`                   | POST       | Create a short URL              |  
| `/api/urls`                   | GET        | List all URLs                   |  
| `/api/urls/:code`             | GET        | Get URL details by short code   |  
| `/api/urls/:code`             | DELETE     | Delete a short URL              |  

### Analytics  
| **Endpoint**                  | **Method** | **Description**                  |  
|-------------------------------|------------|----------------------------------|  
| `/api/urls/:code/stats`       | GET        | Get URL analytics and statistics |  
| `/api/urls/:code/qr`          | GET        | Generate a QR code for a short URL |  

---

## Error Codes  

| **Error Code** | **Message**                   | **Description**                                       |  
|----------------|-------------------------------|-------------------------------------------------------|  
| 400            | `Invalid Request`            | The request payload is malformed or incomplete        |  
| 404            | `URL not found`              | The specified short code does not exist in the system |  
| 409            | `Custom Alias Already Exists`| The requested custom alias is already in use          |  
| 500            | `Internal Server Error`      | An unexpected error occurred on the server            |  

---

## Rate Limits  

- **Max Requests per User**: 100 requests per hour  
- **Rate Limiting**: Excessive requests will return a `429 Too Many Requests` error.  

---

## Setup Guide  

### Prerequisites  
1. Install [Go](https://go.dev/dl/) (version 1.18 or higher).  
2. Install [MongoDB](https://www.mongodb.com/try/download/community).  
3. Install [Postman](https://www.postman.com/) for testing API endpoints (optional).  

### Steps  
1. Clone the repository:  
   ```bash  
   git clone https://github.com/your-username/linksnap.git  
   cd linksnap  
   ```
2. Set up environment variables:
	Create a `.env` file in the project root with the following variables:
	```
	DB_URI=mongodb://localhost:27017  
	COLLECTION_NAME=linksnap  
	PORT=8080  
	```
3. Install dependencies:
	```
	go mod tidy
	```
4. Start the server:
	```
	go run main.go  
	```
5. Test the API endpoints using Postman.
---
## API Documentation

API documentation is published on [here with example requests and responses](https://documenter.getpostman.com/view/26709386/2sAYBd6T15)

---
## License
This project is licensed under the MIT License.
Feel free to contribute or suggest improvements! 🚀


