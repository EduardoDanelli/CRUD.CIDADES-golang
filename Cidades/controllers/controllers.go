package controllers

import (
	"net/http"
	"teste/database"
	"teste/models"

	"github.com/gin-gonic/gin"
)

// Cidade...
type Cidade struct {
	Nome        string `json:"nome"`
	SiglaEstado string `json:"siglaEstado"`
}

// Criar ...
func CriarCidade(c *gin.Context) {

	db := database.SetupDB()
	defer db.Close()

	var criar Cidade
	if err := c.ShouldBindJSON(&criar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cidade := models.Cidades{
		ID:          0,
		Nome:        criar.Nome,
		SiglaEstado: criar.SiglaEstado,
	}
	if err := db.Create(&cidade).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dados": cidade})
}

// ListarCidades ...
func ListarCidades(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	var cidades []models.Cidades
	if err := db.Table("cidades").Where("siglaestado ilike ?", "%"+c.Query("sigla")+"%").Find(&cidades).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível encotrar cidades"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dados": cidades})
}

// BuscarCidade ...
func BuscarCidade(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	var cidade models.Cidades

	if err := db.Table("cidades").Where("id = ?", c.Param("id")).Take(&cidade).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cidade não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dados": cidade})

}

// Editar...
func EditarCidade(c *gin.Context) {

	db := database.SetupDB()
	defer db.Close()
	
	var editar Cidade
	if err := c.ShouldBindJSON(&editar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	 var cidade = models.Cidades{
		ID:          0,
		Nome:        editar.Nome,
		SiglaEstado: editar.SiglaEstado,
	}
	if err := db.Table("cidades").Where("id = ?", c.Param("id")).Update(&cidade).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dados": cidade})

}

// ExcluirCidade
func ExcluirCidade(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	var cidade models.Cidades

	if err := db.Table("cidades").Where("id = ?", c.Param("id")).Delete(&cidade).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dados": cidade})
}
