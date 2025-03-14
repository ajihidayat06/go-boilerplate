package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ReadRequestParamID(c *fiber.Ctx) (int64, error) {
	idStr := c.Params("id")

	idTmp, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid request id")
	}

	return int64(idTmp), nil
}

func GetFiltersAndPagination(c *fiber.Ctx) (map[string]interface{}, int, int) {
	filters := make(map[string]interface{})

	// Loop semua query params
	for k, v := range c.Queries() {
		if k == "page" || k == "limit" {
			continue
		}
		filters[k] = v
	}

	page := c.QueryInt("page")
	limit := c.QueryInt("limit")

	return filters, page, limit
}
