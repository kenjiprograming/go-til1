// https://qiita.com/sedori/items/840e39a0cbf9d5bff006
// https://qiita.com/sedori/items/1a3b8142ae468c7ed053

package main

import (
	"til2_go_gin_gorm/routes"
)

func main()  {
	router := routes.GetApiRouter()
	router.Run(":8080")
}