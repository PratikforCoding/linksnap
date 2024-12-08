package handlers

import (
	"context"
	"net/http"
	"os"
	"time"
	"github.com/PratikforCoding/linksnap/database"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson"
)

// GetQRCode generates a QR code for a short URL
func GetQRCode(c *gin.Context) {
	shortCode := c.Param("code")

	// Get the collection
	collection := database.GetCollection(os.Getenv("COLLECTION_NAME"))
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var url struct {
		ShortCode string `bson:"shortCode"`
		LongURL string   `bson:"longUrl"`
	}

	// Find the URL in the database
	err := collection.FindOne(ctx, bson.M{"shortCode": shortCode}).Decode(&url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	// Generate the QR code
	qrCode, err := qrcode.Encode(url.LongURL, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// Return the QR code as an image
	c.Header("content-type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(qrCode)
}