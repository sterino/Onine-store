package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type PaymentHandler struct {
	paymentUrl string
}

func NewPaymentHandler(paymentUrl string) *PaymentHandler {
	return &PaymentHandler{paymentUrl}
}

// CreatePayment godoc
// @Summary Create payment
// @Description Create payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param payment body payment.Request true "Payment data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments [post]
func (p *PaymentHandler) CreatePayment(c *gin.Context) {
	req, err := http.NewRequest("POST", p.paymentUrl, c.Request.Body)
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
	_, err = io.Copy(c.Writer, resp.Body)
}

// ListPayments godoc
// @Summary List all payments
// @Description List all payments
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments [get]
func (p *PaymentHandler) ListPayments(c *gin.Context) {
	req, err := http.NewRequest("GET", p.paymentUrl, nil)
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
	_, err = io.Copy(c.Writer, resp.Body)
}

// GetPayment godoc
// @Summary Get payment
// @Description Get payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments/{id} [get]
func (p *PaymentHandler) GetPayment(c *gin.Context) {
	req, err := http.NewRequest("GET", p.paymentUrl+"/"+c.Param("id"), nil)
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

	_, err = io.Copy(c.Writer, resp.Body)
}

// UpdatePayment godoc
// @Summary Update payment
// @Description Update payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Param payment body payment.Request true "Payment data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/{id} [put]
func (p *PaymentHandler) UpdatePayment(c *gin.Context) {
	req, err := http.NewRequest("PUT", p.paymentUrl+"/"+c.Param("id"), c.Request.Body)
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

	_, err = io.Copy(c.Writer, resp.Body)
}

// DeletePayment godoc
// @Summary Delete payment
// @Description Delete payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/{id} [delete]
func (p *PaymentHandler) DeletePayment(c *gin.Context) {
	req, err := http.NewRequest("DELETE", p.paymentUrl+"/"+c.Param("id"), nil)
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
	_, err = io.Copy(c.Writer, resp.Body)
}

// SearchPayments godoc
// @Summary Search payments
// @Description Search payments
// @Tags payments
// @Accept  json
// @Produce  json
// @Param filter query string false "Filter"
// @Param value query string false "Value"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/search [get]
func (p *PaymentHandler) SearchPayments(c *gin.Context) {
	req, err := http.NewRequest("GET", p.paymentUrl+"/search?"+c.Request.URL.RawQuery, nil)
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
	_, err = io.Copy(c.Writer, resp.Body)
}
