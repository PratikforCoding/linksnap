package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"strings"

	"github.com/PratikforCoding/linksnap/database"
	"github.com/PratikforCoding/linksnap/models"
	"github.com/PratikforCoding/linksnap/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateShortURL(c *gin.Context) {

	var req struct {
		LongURL		string `json:"longUrl" binding:"required,url"`
		CustomAlias string `json:"customAlias"`
	}

	// Bind raw input first to sanitize it
	rawInput := make(map[string]string)
	if err := c.ShouldBindJSON(&rawInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sanitize inputs
	rawInput["longUrl"] = utils.SanitizeString(rawInput["longUrl"])
	rawInput["customAlias"] = utils.SanitizeString(rawInput["customAlias"])

	// Bind sanitized input to the struct
	req.LongURL = rawInput["longUrl"]
	req.CustomAlias = rawInput["customAlias"]

	if !strings.HasPrefix(req.LongURL, "http://") && !strings.HasPrefix(req.LongURL, "https://") {
		req.LongURL = "https://" + req.LongURL
	}

	if err := utils.ValidateURL(req.LongURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	fmt.Println("Collection name:", os.Getenv("COLLECTION_NAME")) //extra

	// check duplicate custom url
	if req.CustomAlias != "" {
		var existing models.URL
		err := collection.FindOne(ctx, bson.M{"customAlias": req.CustomAlias}).Decode(&existing)
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Custom Alias already esists"})
			return
		}
	}

	// generate short code
	shortCode := req.CustomAlias
	if shortCode == "" {
		shortCode = uuid.New().String()[:6]
	}

	// create short url
	shortURL := "http://linksnap.io/" + shortCode

	// generate qr code
	qrCode, err := utils.GenerateQRCode(shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR Code"})
		return
	}

	// save to database
	newURL := models.URL {
		ShortCode: shortCode,
		LongURL: req.LongURL,
		CustomAlias: req.CustomAlias,
		CreatedAt: time.Now(),
		Clicks: 0,
	}

	
	_, err = collection.InsertOne(ctx, newURL)
	if err != nil {
		fmt.Printf("InsertOne error: %v\n", err) // extra
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"shortUrl": shortURL,
		"qrCode": qrCode,
		"created": newURL.CreatedAt,
	})
}

func GetURLDetails(c *gin.Context) {
	shortCode := c.Param("code")


	collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url models.URL
	err := collection.FindOne(ctx, bson.M{"shortCode": shortCode}).Decode(&url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortCode": url.ShortCode,
		"longUrl": url.LongURL,
		"customAlias": url.CustomAlias,
		"createdAt": url.CreatedAt,
		"clicks": url.Clicks,
	})
}

func ListURLs(c *gin.Context) {
    // Access the collection from the database
    collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
    
    // Set a timeout context for the database operation
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Define a variable to store the list of URLs
    var urls []models.URL
    
    // Query all documents from the collection
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URLs"})
        return
    }
    defer cursor.Close(ctx)

    // Decode the result into the urls slice
    for cursor.Next(ctx) {
        var url models.URL
        if err := cursor.Decode(&url); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode URL"})
            return
        }
        urls = append(urls, url)
    }

    // Check for any errors during iteration
    if err := cursor.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over URLs"})
        return
    }

    // Return the list of URLs in the response
    c.JSON(http.StatusOK, gin.H{
        "urls": urls,
    })
}
