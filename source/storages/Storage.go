//go:build wasm

package storages

import "strconv"
import "syscall/js"

type Storage struct {
	name string `json:"name"`
}

func (storage *Storage) Clear() {
	js.Global().Get(storage.name).Call("clear")
}

func (storage *Storage) GetItem(key string) string {

	var result string

	wrapped_key := js.ValueOf(key)
	tmp := js.Global().Get(storage.name).Call("getItem", wrapped_key)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = tmp.String()
	}

	return result

}

func (storage *Storage) GetItemBytes(key string) []byte {

	var result []byte

	tmp := storage.GetItem(key)

	if len(tmp) > 0 {
		result = []byte(tmp)
	}

	return result

}

func (storage *Storage) GetItemFloat32(key string) float32 {

	var result float32

	tmp := storage.GetItem(key)
	num, err := strconv.ParseFloat(tmp, 32)

	if err == nil {
		result = float32(num)
	}

	return result

}

func (storage *Storage) GetItemFloat64(key string) float64 {

	var result float64

	tmp := storage.GetItem(key)
	num, err := strconv.ParseFloat(tmp, 64)

	if err == nil {
		result = float64(num)
	}

	return result

}

func (storage *Storage) GetItemInt(key string) int {

	var result int

	tmp := storage.GetItem(key)
	num, err := strconv.ParseInt(tmp, 10, 64)

	if err == nil {
		result = int(num)
	}

	return result

}

func (storage *Storage) GetItemInt8(key string) int8 {

	var result int8

	tmp := storage.GetItem(key)
	num, err := strconv.ParseInt(tmp, 10, 8)

	if err == nil {
		result = int8(num)
	}

	return result

}

func (storage *Storage) GetItemInt16(key string) int16 {

	var result int16

	tmp := storage.GetItem(key)
	num, err := strconv.ParseInt(tmp, 10, 16)

	if err == nil {
		result = int16(num)
	}

	return result

}

func (storage *Storage) GetItemInt32(key string) int32 {

	var result int32

	tmp := storage.GetItem(key)
	num, err := strconv.ParseInt(tmp, 10, 32)

	if err == nil {
		result = int32(num)
	}

	return result

}

func (storage *Storage) GetItemInt64(key string) int64 {

	var result int64

	tmp := storage.GetItem(key)
	num, err := strconv.ParseInt(tmp, 10, 64)

	if err == nil {
		result = int64(num)
	}

	return result

}

func (storage *Storage) GetItemUint(key string) uint {

	var result uint

	tmp := storage.GetItem(key)
	num, err := strconv.ParseUint(tmp, 10, 64)

	if err == nil {
		result = uint(num)
	}

	return result

}

func (storage *Storage) GetItemUint8(key string) uint8 {

	var result uint8

	tmp := storage.GetItem(key)
	num, err := strconv.ParseUint(tmp, 10, 8)

	if err == nil {
		result = uint8(num)
	}

	return result

}

func (storage *Storage) GetItemUint16(key string) uint16 {

	var result uint16

	tmp := storage.GetItem(key)
	num, err := strconv.ParseUint(tmp, 10, 16)

	if err == nil {
		result = uint16(num)
	}

	return result

}

func (storage *Storage) GetItemUint32(key string) uint32 {

	var result uint32

	tmp := storage.GetItem(key)
	num, err := strconv.ParseUint(tmp, 10, 32)

	if err == nil {
		result = uint32(num)
	}

	return result

}

func (storage *Storage) GetItemUint64(key string) uint64 {

	var result uint64

	tmp := storage.GetItem(key)
	num, err := strconv.ParseUint(tmp, 10, 64)

	if err == nil {
		result = uint64(num)
	}

	return result

}

func (storage *Storage) Key(index uint) string {

	var result string

	wrapped_index := js.ValueOf(index)

	tmp := js.Global().Get(storage.name).Call("key", wrapped_index)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = tmp.String()
	}

	return result

}

func (storage *Storage) Length() uint {

	var result uint

	tmp := js.Global().Get(storage.name).Get("length")

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = uint(tmp.Int())
	}

	return result

}

func (storage *Storage) RemoveItem(key string) {
	js.Global().Get(storage.name).Call("removeItem", js.ValueOf(key))
}

func (storage *Storage) SetItem(key string, value any) {

	switch value.(type) {
	case []byte:

		tmp1, ok := value.([]byte)

		if ok {
			tmp2 := string(tmp1)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case float32:

		tmp1, ok := value.(float32)

		if ok {
			tmp2 := strconv.FormatFloat(float64(tmp1), 'g', -1, 32)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case float64:

		tmp1, ok := value.(float64)

		if ok {
			tmp2 := strconv.FormatFloat(float64(tmp1), 'g', -1, 64)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case int:

		tmp1, ok := value.(int)

		if ok {
			tmp2 := strconv.FormatInt(int64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case int8:

		tmp1, ok := value.(int8)

		if ok {
			tmp2 := strconv.FormatInt(int64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case int16:

		tmp1, ok := value.(int16)

		if ok {
			tmp2 := strconv.FormatInt(int64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case int32:

		tmp1, ok := value.(int32)

		if ok {
			tmp2 := strconv.FormatInt(int64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case int64:

		tmp1, ok := value.(int64)

		if ok {
			tmp2 := strconv.FormatInt(int64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case uint:

		tmp1, ok := value.(uint)

		if ok {
			tmp2 := strconv.FormatUint(uint64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case uint8:

		tmp1, ok := value.(uint8)

		if ok {
			tmp2 := strconv.FormatUint(uint64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case uint16:

		tmp1, ok := value.(uint16)

		if ok {
			tmp2 := strconv.FormatUint(uint64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case uint32:

		tmp1, ok := value.(uint32)

		if ok {
			tmp2 := strconv.FormatUint(uint64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case uint64:

		tmp1, ok := value.(uint64)

		if ok {
			tmp2 := strconv.FormatUint(uint64(tmp1), 10)
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp2))
		}

	case string:

		tmp1, ok := value.(string)

		if ok {
			js.Global().Get(storage.name).Call("setItem", js.ValueOf(key), js.ValueOf(tmp1))
		}

	}

}
