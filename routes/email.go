package routes

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"

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
	_ = godotenv.Load()

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

	// initialize googleOauthConfig after ClientID and ClientSecret are set
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://server.lostengineering.com/email/handleCallback",
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
		Endpoint:     google.Endpoint,
	}
	log.Println("OAuth config initialized.")
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
		log.Printf("JWT parsing error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired JWT token"})
		return
	}

	session := sessions.Default(c)
	state := generateStateOauthCookie()
	session.Set("state", state)
	if err := session.Save(); err != nil {
		log.Printf("Error saving session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Session save error"})
		return
	}

	url := googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	log.Printf("Redirecting to: %s", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleCallback(c *gin.Context) {
	session := sessions.Default(c)

	receivedState := c.Query("state")
	originalState := session.Get("state")
	if receivedState != originalState {
		log.Println("Invalid session state")
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"Error": "Invalid session state."})
		return
	}

	code := c.Query("code")
	if code == "" {
		log.Println("Authorization code not found")
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"Error": "Code not found"})
		return
	}

	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		log.Printf("Token exchange failed: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"Error": "Failed to exchange token"})
		return
	}

	// render token.html with a script that will give the token to the client (electron app)
	c.HTML(http.StatusOK, "token.html", gin.H{"AccessToken": token.AccessToken})
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
	var requestData struct {
		AccessToken    string `json:"accessToken"`
		RecipientEmail string `json:"recipientEmail"`
		Subject        string `json:"subject"`
		Content        string `json:"content"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// set up OAuth token using the provided access token and create a new HTTP client using the OAuth token
	oauthToken := &oauth2.Token{
		AccessToken: requestData.AccessToken,
	}
	client := googleOauthConfig.Client(c, oauthToken)

	// create a new Gmail service using the authorized client
	srv, err := gmail.NewService(c, option.WithHTTPClient(client))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Gmail service"})
		return
	}

	// create the email message
	var message gmail.Message
	emailTo := "To: " + requestData.RecipientEmail + "\r\n"
	subject := "Subject: " + requestData.Subject + "\r\n"
	mime := "MIME-version: 1.0; Content-Type: text/plain; charset=\"UTF-8\";\r\n\r\n"
	msgStr := []byte(emailTo + subject + mime + "\n" + requestData.Content)
	message.Raw = base64.URLEncoding.EncodeToString(msgStr)

	// send the email
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

func getEmailAddress(c *gin.Context) {
	var requestData struct {
		AccessToken string `json:"accessToken"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	oauthToken := &oauth2.Token{
		AccessToken: requestData.AccessToken,
	}

	client := googleOauthConfig.Client(c, oauthToken)

	srv, err := gmail.NewService(c, option.WithHTTPClient(client))
	if err != nil {
		log.Print("Failed to create gmail service")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Gmail service"})
		return
	}

	userProfile, err := srv.Users.GetProfile("me").Do()
	if err != nil {
		log.Print("Failed to retrieve user profile")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"emailAddress": userProfile.EmailAddress})
}
