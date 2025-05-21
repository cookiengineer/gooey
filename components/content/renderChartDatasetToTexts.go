package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/data"
import "strconv"

func renderTextAt(x int, y int, label string) *dom.Element {

	text := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "text")
	text.SetAttribute("text-anchor", "middle")
	text.SetAttribute("dominant-baseline", "middle")
	text.SetAttribute("x", strconv.Itoa(x))
	text.SetAttribute("y", strconv.Itoa(y))
	text.SetInnerHTML(label)

	return text

}

func renderChartDatasetToTexts(dataset *data.Dataset, width int, height int, min_value int64, max_value int64, property string) []*dom.Element {

	texts   := make([]*dom.Element, 0)
	delta_x := int(width)
	delta_y := int(height)

	if dataset.Length() > 1 {
		delta_x = int(width / (dataset.Length() - 1))
	}

	if max_value > 0 {
		delta_y = int(int64(height) / max_value)
	}

	if min_value < 0 {
		delta_y = int(int64(height) / (int64(max_value) + int64(-1 * min_value)))
	}

	for index := 0; index < dataset.Length(); index++ {

		data    := dataset.Get(index)
		val, ok := (*data)[property]

		if ok == true {

			switch val.(type) {

			case []byte:
				// Do Nothing

			case bool:

				value := val.(bool)

				if value == true {

					pos_x := int(delta_x * index)
					pos_y := 0
					label := strconv.FormatBool(value)

					texts = append(texts, renderTextAt(pos_x, pos_y, label))

				} else if value == false {

					pos_x := int(delta_x * index)
					pos_y := height
					label := strconv.FormatBool(value)

					texts = append(texts, renderTextAt(pos_x, pos_y, label))

				}

			case float32:

				value := val.(float32)
				pos_x := int(delta_x * index)
				pos_y := height - int(float32(delta_y) * (value - float32(min_value)))
				label := strconv.FormatFloat(float64(value), 'g', -1, 32)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case float64:

				value := val.(float64)
				pos_x := int(delta_x * index)
				pos_y := height - int(float64(delta_y) * (value - float64(min_value)))
				label := strconv.FormatFloat(value, 'g', -1, 32)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int:

				value := val.(int)
				pos_x := int(delta_x * index)
				pos_y := height - int(int(delta_y) * (value - int(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int8:

				value := val.(int8)
				pos_x := int(delta_x * index)
				pos_y := height - int(int8(delta_y) * (value - int8(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int16:

				value := val.(int16)
				pos_x := int(delta_x * index)
				pos_y := height - int(int16(delta_y) * (value - int16(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int32:

				value := val.(int32)
				pos_x := int(delta_x * index)
				pos_y := height - int(int32(delta_y) * (value - int32(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int64:

				value := val.(int64)
				pos_x := int(delta_x * index)
				pos_y := height - int(int64(delta_y) * (value - int64(min_value)))
				label := strconv.FormatInt(value, 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint:

				value := val.(uint)
				pos_x := int(delta_x * index)
				pos_y := height - int(uint(delta_y) * (value - uint(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint8:

				value := val.(uint8)
				pos_x := int(delta_x * index)
				pos_y := height - int(uint8(delta_y) * (value - uint8(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint16:

				value := val.(uint16)
				pos_x := int(delta_x * index)
				pos_y := height - int(uint16(delta_y) * (value - uint16(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint32:

				value := val.(uint32)
				pos_x := int(delta_x * index)
				pos_y := height - int(uint32(delta_y) * (value - uint32(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint64:

				value := val.(uint64)
				pos_x := int(delta_x * index)
				pos_y := height - int(uint64(delta_y) * (value - uint64(min_value)))
				label := strconv.FormatUint(value, 10)

				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			}

		}

	}

	if len(texts) > 0 {
		texts[0].SetAttribute("text-anchor", "start")
	}

	if len(texts) >= 2 {
		texts[len(texts)-1].SetAttribute("text-anchor", "end")
	}

	return texts

}
