package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userUrl string
}

func NewUserHandler(userUrl string) *UserHandler {
	return &UserHandler{userUrl}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body user.Request true "User data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	req, err := http.NewRequest("POST", u.userUrl, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp)
}

// ListUsers godoc
// @Summary List all users
// @Description List all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func (u *UserHandler) ListUsers(c *gin.Context) {
	req, err := http.NewRequest("GET", u.userUrl, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp)
}

// GetUser godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [get]
func (u *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest("GET", u.userUrl+"/"+id, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp)
}

// UpdateUser godoc
// @Summary Update user by id
// @Description Update user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body user.Request true "User data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest("PUT", u.userUrl+"/"+id, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp)
}

// DeleteUser godoc
// @Summary Delete user by id
// @Description Delete user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [delete]
func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest("DELETE", u.userUrl+"/"+id, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp)
}

// SearchUser godoc
// @Summary Search user
// @Description Search user
// @Tags users
// @Accept  json
// @Produce  json
// @Param filter query string true "Search filter"
// @Param value query string true "Search value"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/search [get]
func (u *UserHandler) SearchUser(c *gin.Context) {
	filter := c.Query("filter")
	val := c.Query("value")
	req, err := http.NewRequest("GET", u.userUrl+"/search?filter="+filter+"&value="+val, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, resp)
}
