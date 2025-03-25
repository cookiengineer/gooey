package layout

func getHeaderItemByPath(path string, items map[string]*header_view_item) header_view_item {

	var result header_view_item

	for _, item := range items {

		if item.Path == path {
			result = *item
			break
		}

	}

	return result

}
