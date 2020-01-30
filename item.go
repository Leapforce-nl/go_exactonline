package exactonline

import (
	"fmt"
	"strconv"

	types "github.com/Leapforce-nl/go_types"
)

type Item struct {
	ID   types.GUID `json:"ID"`
	Code string     `json:"Code"`
}

func (eo *ExactOnline) GetItems() error {
	selectFields := GetJsonTaggedFieldNames(Item{})
	urlStr := fmt.Sprintf("%s%s/logistics/Items?$select=%s", eo.ApiUrl, strconv.Itoa(eo.Division), selectFields)
	//fmt.Println(urlStr)

	for urlStr != "" {
		it := []Item{}

		str, err := eo.Get(urlStr, &it)
		if err != nil {
			fmt.Println("ERROR in GetItems:", err)
			fmt.Println("url:", urlStr)
			return err
		}

		eo.Items = append(eo.Items, it...)

		urlStr = str
	}

	return nil
}
