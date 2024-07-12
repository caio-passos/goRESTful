package main


import (
	"net/http"
  "github.com/gin-gonic/gin"
)

func main(){
  router := gin.Default()
  router.GET("/stocks", getStocks)
	router.GET("/stocks/:id", getStockByID)
  router.POST("/stocks", postStocks)

  router.Run("localhost:8080")
}


type stocks struct {
  ID string `json:"id"` 
  Name string `json:"name"`
  Company string `json:"company"`
  Price float64 `json:"price"`
}

var stocks_slices = []stocks {
  {ID: "1", Name: "MSTF", Company: "Microsoft", Price:454.0},
  {ID: "2", Name: "GOOG", Company: "Google", Price:187.30},
  {ID: "3", Name: "BRK-A", Company: "Berkshire Hathaway Inc.", Price:418.78},
  {ID: "4", Name: "AMD", Company: "AMD", Price:181.94},

}

// retorns a list with all stocks as JSON
func getStocks(c *gin.Context){
  c.IndentedJSON(http.StatusOK, stocks_slices)
}

func postStocks(c *gin.Context) {
  var newStocks stocks

  if err := c.BindJSON(&newStocks); err != nil {
    return
  }

  // add new stock to the slice
  stocks_slices = append(stocks_slices, newStocks)
  c.IndentedJSON(http.StatusCreated, newStocks)
}

func getStockByID(c *gin.Context){
  id := c.Param("id")

  for _, a := range stocks_slices {
        if a.ID == id {
          c.IndentedJSON(http.StatusOK, a)
          return
        }
      }
      c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Stock not found"})
}
