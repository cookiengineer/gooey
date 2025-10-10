package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/data"
import "strconv"
import "strings"

func renderTextAt(x int, y int, label string) *dom.Element {

	text := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "text")
	text.SetAttribute("text-anchor", "middle")
	text.SetAttribute("dominant-baseline", "middle")
	text.SetAttribute("x", strconv.Itoa(x))
	text.SetAttribute("y", strconv.Itoa(y))
	text.SetInnerHTML(label)

	return text

}

func renderLineChartDataset(dataset *data.Dataset, width int, height int, min_value int64, max_value int64, property string) (*dom.Element, []*dom.Element) {

	path := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "path")
	texts := make([]*dom.Element, 0)
	delta_x := float64(width)
	delta_y := float64(height)

	if dataset.Length() > 1 {
		delta_x = float64(float64(width) / (float64(dataset.Length()) - float64(1)))
	}

	if max_value > 0 {
		delta_y = float64(float64(height) / float64(max_value))
	}

	if min_value < 0 {
		delta_y = float64(float64(height) / (float64(max_value) + float64(float64(-1)*float64(min_value))))
	}

	description := make([]string, 0)

	description = append(description, "M 0 "+strconv.Itoa(height+1))

	for index := 0; index < dataset.Length(); index++ {

		data := dataset.Get(index)
		val, ok := (*data)[property]

		if ok == true {

			switch val.(type) {

			case []byte:
				// Do Nothing
			case bool:

				value := val.(bool)

				if value == true {

					pos_x := int(delta_x * float64(index))
					pos_y := 0
					label := strconv.FormatBool(value)

					description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
					texts = append(texts, renderTextAt(pos_x, pos_y, label))

				} else if value == false {

					pos_x := int(delta_x * float64(index))
					pos_y := height
					label := strconv.FormatBool(value)

					description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
					texts = append(texts, renderTextAt(pos_x, pos_y, label))

				}

			case float32:

				value := val.(float32)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatFloat(float64(value), 'g', -1, 32)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case float64:

				value := val.(float64)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatFloat(value, 'g', -1, 32)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int:

				value := val.(int)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int8:

				value := val.(int8)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int16:

				value := val.(int16)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int32:

				value := val.(int32)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatInt(int64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case int64:

				value := val.(int64)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatInt(value, 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case string:

				value := val.(string)

				if strings.Contains(value, ".") && strings.HasSuffix(value, "%") {

					tmp, err := strconv.ParseFloat(value[0:len(value)-1], 32)

					if err == nil && tmp >= 0.0 && tmp <= 100.0 {

						percentage := float64(tmp) / float64(100.0)
						pos_x := int(delta_x * float64(index))
						pos_y := height - int(delta_y*percentage)
						label := value

						description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
						texts = append(texts, renderTextAt(pos_x, pos_y, label))

					}

				} else if strings.HasSuffix(value, "%") {

					tmp, err := strconv.ParseInt(value[0:len(value)-1], 10, 32)

					if err == nil && tmp >= 0 && tmp <= 100 {

						percentage := float64(tmp) / float64(100.0)
						pos_x := int(delta_x * float64(index))
						pos_y := height - int(delta_y*percentage)
						label := value

						description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
						texts = append(texts, renderTextAt(pos_x, pos_y, label))

					}

				}

			case uint:

				value := val.(uint)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint8:

				value := val.(uint8)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint16:

				value := val.(uint16)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint32:

				value := val.(uint32)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatUint(uint64(value), 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			case uint64:

				value := val.(uint64)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y*(float64(value)-float64(min_value)))
				label := strconv.FormatUint(value, 10)

				description = append(description, "L "+strconv.Itoa(pos_x)+" "+strconv.Itoa(pos_y))
				texts = append(texts, renderTextAt(pos_x, pos_y, label))

			}

		}

	}

	description = append(description, "L "+strconv.Itoa(width-1)+" "+strconv.Itoa(height+1))

	path.SetAttribute("d", strings.Join(description, " ")+" Z")

	if len(texts) > 0 {
		texts[0].SetAttribute("text-anchor", "start")
	}

	if len(texts) >= 2 {
		texts[len(texts)-1].SetAttribute("text-anchor", "end")
	}

	return path, texts

}
