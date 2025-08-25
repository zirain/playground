package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct {
	CreatedAt    string                 `json:"created_at"`
	RequestCount int                    `json:"request_count"`
	UserData     map[string]interface{} `json:"user_data"`
}

type Envelope struct {
	SessionID    string `json:"session_id"`
	RequestCount int    `json:"request_count,omitempty"`
}

var (
	sessions   = make(map[string]*Session)
	sessionsMu sync.Mutex
)

func parseEnvelopeSession(sessionHeader string) (*Envelope, error) {
	decoded, err := base64.StdEncoding.DecodeString(sessionHeader)
	if err != nil {
		return nil, err
	}
	var envelope Envelope
	if err := json.Unmarshal(decoded, &envelope); err != nil {
		return nil, err
	}
	return &envelope, nil
}

func createEnvelopeSession(sessionID string, additional map[string]interface{}) (string, error) {
	envelope := map[string]interface{}{
		"session_id": sessionID,
	}
	for k, v := range additional {
		envelope[k] = v
	}
	jsonBytes, err := json.Marshal(envelope)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

func isEnvelopeFormat(sessionHeader string) bool {
	decoded, err := base64.StdEncoding.DecodeString(sessionHeader)
	if err != nil {
		return false
	}
	var js map[string]interface{}
	return json.Unmarshal(decoded, &js) == nil
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		sessionHeader := c.GetHeader("x-session-id")
		var sessionID string
		var format string
		var envelope *Envelope

		sessionsMu.Lock()
		defer sessionsMu.Unlock()

		if sessionHeader != "" {
			if isEnvelopeFormat(sessionHeader) {
				format = "envelope"
				env, err := parseEnvelopeSession(sessionHeader)
				if err == nil && env.SessionID != "" {
					sessionID = env.SessionID
					envelope = env
					if sess, ok := sessions[sessionID]; ok {
						sess.RequestCount++
						additional := map[string]interface{}{"request_count": sess.RequestCount}
						newEnvelope, _ := createEnvelopeSession(sessionID, additional)
						c.Header("x-session-id", newEnvelope)
						c.JSON(http.StatusOK, gin.H{
							"message":       fmt.Sprintf("Hello! Continuing envelope session %s", sessionID),
							"session_id":    sessionID,
							"format":        format,
							"envelope":      envelope,
							"request_count": sess.RequestCount,
							"session_data":  sess,
						})
						return
					}
				}
			} else {
				format = "header"
				sessionID = sessionHeader
				if sess, ok := sessions[sessionID]; ok {
					sess.RequestCount++
					c.Header("x-session-id", sessionID)
					c.JSON(http.StatusOK, gin.H{
						"message":       fmt.Sprintf("Hello! Continuing header session %s", sessionID),
						"session_id":    sessionID,
						"format":        format,
						"request_count": sess.RequestCount,
						"session_data":  sess,
					})
					return
				}
			}
		}

		// Create new session
		sessionID = uuid.New().String()
		sessions[sessionID] = &Session{
			CreatedAt:    "now",
			RequestCount: 1,
			UserData:     map[string]interface{}{},
		}
		useEnvelope := strings.ToLower(c.Query("envelope")) == "true"
		if useEnvelope {
			envelopeSession, _ := createEnvelopeSession(sessionID, map[string]interface{}{"request_count": 1})
			c.Header("x-session-id", envelopeSession)
			c.JSON(http.StatusOK, gin.H{
				"message":       fmt.Sprintf("Hello! New envelope session created: %s", sessionID),
				"session_id":    sessionID,
				"format":        "envelope",
				"request_count": 1,
			})
		} else {
			c.Header("x-session-id", sessionID)
			c.JSON(http.StatusOK, gin.H{
				"message":       fmt.Sprintf("Hello! New header session created: %s", sessionID),
				"session_id":    sessionID,
				"format":        "header",
				"request_count": 1,
			})
		}
	})

	r.GET("/status", func(c *gin.Context) {
		sessionHeader := c.GetHeader("x-session-id")
		resp := gin.H{
			"endpoint":               "/status",
			"session_header_present": sessionHeader != "",
			"session_header":         sessionHeader,
			"total_sessions":         len(sessions),
		}
		if sessionHeader != "" {
			if isEnvelopeFormat(sessionHeader) {
				env, _ := parseEnvelopeSession(sessionHeader)
				resp["format"] = "envelope"
				resp["envelope"] = env
				if env != nil && env.SessionID != "" {
					resp["session_id"] = env.SessionID
					resp["session_valid"] = sessions[env.SessionID] != nil
				}
			} else {
				resp["format"] = "header"
				resp["session_id"] = sessionHeader
				resp["session_valid"] = sessions[sessionHeader] != nil
			}
		}
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/data", func(c *gin.Context) {
		sessionHeader := c.GetHeader("x-session-id")
		var sessionID string
		if sessionHeader != "" {
			if isEnvelopeFormat(sessionHeader) {
				env, _ := parseEnvelopeSession(sessionHeader)
				if env != nil {
					sessionID = env.SessionID
				}
			} else {
				sessionID = sessionHeader
			}
		}
		resp := gin.H{
			"method":      "GET",
			"session_id":  sessionID,
			"sample_data": gin.H{"key": "value", "number": 42},
			"session_data": func() interface{} {
				if sessionID != "" {
					return sessions[sessionID]
				}
				return gin.H{}
			}(),
			"message": func() string {
				if sessionID != "" {
					return fmt.Sprintf("Data for session %s", sessionID)
				}
				return "Data (no session)"
			}(),
		}
		if sessionID != "" && sessions[sessionID] != nil {
			sessions[sessionID].RequestCount++
		}
		c.Header("x-session-id", sessionHeader)
		c.JSON(http.StatusOK, resp)
	})

	r.POST("/data", func(c *gin.Context) {
		sessionHeader := c.GetHeader("x-session-id")
		var sessionID string
		if sessionHeader != "" {
			if isEnvelopeFormat(sessionHeader) {
				env, _ := parseEnvelopeSession(sessionHeader)
				if env != nil {
					sessionID = env.SessionID
				}
			} else {
				sessionID = sessionHeader
			}
		}
		var data map[string]interface{}
		if err := c.BindJSON(&data); err != nil {
			data = map[string]interface{}{}
		}
		stored := false
		if sessionID != "" && sessions[sessionID] != nil {
			for k, v := range data {
				sessions[sessionID].UserData[k] = v
			}
			sessions[sessionID].RequestCount++
			stored = true
		}
		resp := gin.H{
			"method":         "POST",
			"received_data":  data,
			"session_id":     sessionID,
			"session_stored": stored,
			"message": func() string {
				if sessionID != "" {
					return fmt.Sprintf("Data received for session %s", sessionID)
				}
				return "Data received (no session)"
			}(),
		}
		c.Header("x-session-id", sessionHeader)
		c.JSON(http.StatusOK, resp)
	})

	r.GET("/sessions", func(c *gin.Context) {
		sessionsMu.Lock()
		defer sessionsMu.Unlock()
		c.JSON(http.StatusOK, gin.H{
			"active_sessions": len(sessions),
			"sessions":        sessions,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "backend",
		})
	})

	fmt.Println("Starting Gin backend server...")
	fmt.Println("Supports both header-based and envelope session formats:")
	fmt.Println("  Header format: x-session-id: <session-uuid>")
	fmt.Println("  Envelope format: x-session-id: base64(json({session_id: <uuid>}))")
	fmt.Println("")
	fmt.Println("Test commands:")
	fmt.Println("  # Header-based sessions (default):")
	fmt.Println("  curl http://localhost:8000/")
	fmt.Println("  curl http://localhost:8000/status")
	fmt.Println("  curl http://localhost:8000/sessions")
	fmt.Println("")
	fmt.Println("  # Envelope sessions (add ?envelope=true):")
	fmt.Println("  curl 'http://localhost:8000/?envelope=true'")
	fmt.Println("  curl -X POST -H 'Content-Type: application/json' -d '{\"test\": \"data\"}' http://localhost:8000/data")
	fmt.Println("")

	r.Run("0.0.0.0:8000")
}
