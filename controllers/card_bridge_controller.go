package controllers

import (
	"bab6_golang/database"
	"bab6_golang/entities"
	"bab6_golang/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCards menangani permintaan GET untuk mengambil semua kartu.
func GetCards(c *gin.Context) {
	var result gin.H
	card, err := repositories.GetCards(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": card,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertCard menangani permintaan POST untuk memasukkan kartu baru.
func InsertCard(c *gin.Context) {
	var card entities.Card
	// Mengambil parameter 'id' dari URL
	idCard := c.Param("id")
	card.ID = idCard

	err := repositories.InsertCard(database.DbConnection, card)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan status OK bersama data yang berhasil dimasukkan
	c.JSON(http.StatusOK, card)
}

// DeleteCard menangani permintaan DELETE untuk menghapus kartu berdasarkan ID.
func DeleteCard(c *gin.Context) {
	var card entities.Card
	// Mengambil parameter 'id' dari URL
	idCard := c.Param("id")
	card.ID = idCard

	err := repositories.DeleteCard(database.DbConnection, card)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan pesan sukses jika berhasil dilakukan
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus", "id": idCard})
}
