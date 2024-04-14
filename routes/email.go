package routes

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var (
	ClientID          string
	ClientSecret      string
	googleOauthConfig *oauth2.Config
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	ClientID = os.Getenv("ClientID")
	if ClientID == "" {
		log.Fatal("ClientID not set as env variable")
	}

	ClientSecret = os.Getenv("ClientSecret")
	if ClientSecret == "" {
		log.Fatal("ClientSecret not set as env variable")
	}

	secretJWT := os.Getenv("JWT_SECRET")
	if secretJWT == "" {
		log.Fatal("JWT_SECRET not set as env variable")
	}

	jwtKey = []byte(secretJWT)

	// Initialize googleOauthConfig after ClientID and ClientSecret are set
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://server.lostengineering.com/email/handleCallback",
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
		Endpoint:     google.Endpoint,
	}
}

func handleLogin(c *gin.Context) {
	jwtToken := c.Query("jwt")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token is required"})
		return
	}

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired JWT token"})
		return
	}

	// JWT is valid, proceed with OAuth
	session := sessions.Default(c)
	state := generateStateOauthCookie()
	session.Set("state", state)
	session.Save()

	url := googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleCallback(c *gin.Context) {
	session := sessions.Default(c)
	receivedState := c.Query("state")
	originalState := session.Get("state")

	if receivedState != originalState {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"Error": "Invalid session state.",
		})
		return
	}

	code := c.Query("code")
	if code == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": "Code not found",
		})
		return
	}

	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"Error": "Failed to exchange token",
		})
		return
	}

	// Success HTML page
	c.HTML(http.StatusOK, "success.html", gin.H{
		"Message": "You have successfully logged in with Gmail!",
		"Token":   token.AccessToken, // For demonstration, typically you wouldn't show the token directly
	})
}

func generateStateOauthCookie() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error generating random state: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func handleSendEmail(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization token provided"})
		return
	}

	// Validate the token
	if !validateToken(token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	var requestData struct {
		RecipientEmail string `json:"recipientEmail"`
		Subject        string `json:"subject"`
		Content        string `json:"content"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	oauthToken := &oauth2.Token{
		AccessToken: token,
	}

	client := googleOauthConfig.Client(c, oauthToken)
	srv, err := gmail.NewService(c, option.WithHTTPClient(client))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Gmail service"})
		return
	}

	// Create the email message
	var message gmail.Message
	emailTo := "To: " + requestData.RecipientEmail + "\r\n"
	subject := "Subject: " + requestData.Subject + "\r\n"
	mime := "MIME-version: 1.0; Content-Type: text/plain; charset=\"UTF-8\";\r\n\r\n"
	msgStr := []byte(emailTo + subject + mime + "\n" + requestData.Content)
	message.Raw = base64.URLEncoding.EncodeToString(msgStr)

	// Send the email
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

func validateToken(accessToken string) bool {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/tokeninfo?access_token="+accessToken, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var data struct {
		ErrorDescription string `json:"error_description"`
		ExpiresIn        string `json:"expires_in"`
		Scope            string `json:"scope"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	return data.ErrorDescription == "" && strings.Contains(data.Scope, "https://www.googleapis.com/auth/gmail.send")
}
