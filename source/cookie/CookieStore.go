//go:build wasm

// CookieStore API support
// Chromium ✅
// Edge ✅
// Firefox ❌
// Webkit ❌

package cookie

import (
	"errors"
	"syscall/js"
)

var CookieStore cookieStore

func init() {
	cookieStoreValue := js.Global().Get("cookieStore")
	if cookieStoreValue.IsNull() || cookieStoreValue.IsUndefined() {
		panic("cannot get 'cookieStore'")
	}
	CookieStore = cookieStore{
		Value: &cookieStoreValue,
	}
}

type cookieStore struct {
	Value *js.Value `json:"value"`
}

func (cs *cookieStore) Delete(deleteOptions DeleteOptions) error {
	errChan := make(chan error)
	wrappedDeleteOptions := js.ValueOf(deleteOptions.MapToJS())

	// on_success represents the callback if the promise is resolved
	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]
		if value.IsUndefined() {
			errChan <- nil
		} else {
			errChan <- errors.New("onSuccess callback called, but the result should have been undefined. Check out the docs: https://developer.mozilla.org/en-US/docs/Web/API/CookieStore/delete#return_value")
		}
		return nil
	})

	// on_failure represents the callback that catches an exception when one is thrown
	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]
		errChan <- errors.New(value.Get("message").String())
		return nil
	})

	go cs.Value.Call("delete", wrappedDeleteOptions).Call("then", on_success).Call("catch", on_failure)

	// block and wait for error
	return <-errChan
}

func (cs *cookieStore) Get(getOptions GetOptions) (Cookie, error) {
	type Result struct {
		cookie *Cookie
		err    error
	}

	resultChan := make(chan Result, 1)

	wrappedGetOptions := js.ValueOf(getOptions.MapToJS())

	// on_success represents the callback if the promise is resolved
	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]

		var domain string = ""
		domainValue := value.Get("domain")
		if !domainValue.IsNull() && !domainValue.IsUndefined() {
			domain = domainValue.String()
		}

		var expires int64 = 0
		expiresValue := value.Get("expires")
		if !expiresValue.IsNull() && !expiresValue.IsUndefined() {
			expires = int64(expiresValue.Float())
		}

		var path string = ""
		pathValue := value.Get("path")
		if !pathValue.IsNull() && !pathValue.IsUndefined() {
			path = pathValue.String()
		}

		var sameSite SameSite = Strict
		sameSiteValue := value.Get("sameSite")
		if !sameSiteValue.IsNull() && !sameSiteValue.IsUndefined() {
			sameSite = SameSite(sameSiteValue.String())
		}
		resultChan <- Result{
			cookie: &Cookie{
				Domain:      domain,
				Expires:     expires,
				Name:        value.Get("name").String(),
				Partitioned: value.Get("partitioned").Bool(),
				Path:        path,
				SameSite:    sameSite,
				Secure:      value.Get("secure").Bool(),
				Value:       value.Get("value").String(),
			},
			err: nil,
		}

		return nil
	})

	// on_failure represents the callback that catches an exception when one is thrown
	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]

		resultChan <- Result{
			cookie: nil,
			err:    errors.New(value.Get("message").String()),
		}

		return nil
	})

	go cs.Value.Call("get", wrappedGetOptions).Call("then", on_success).Call("catch", on_failure)

	// block and wait for results
	result := <-resultChan
	return *result.cookie, result.err
}

func (cs *cookieStore) GetAll(getOptions *GetOptions) ([]Cookie, error) {
	type Result struct {
		cookies []Cookie
		err     error
	}

	resultChan := make(chan Result, 1)

	var wrappedGetOptions js.Value = js.Value{}
	if getOptions != nil {
		wrappedGetOptions = js.ValueOf(getOptions.MapToJS())
	}

	// on_success represents the callback if the promise is resolved
	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 || args[0].IsUndefined() {
			resultChan <- Result{cookies: nil, err: errors.New("no cookies found!")}
			return nil
		}

		jsCookies := args[0]
		var cookies []Cookie

		for i := 0; i < jsCookies.Length(); i++ {
			value := jsCookies.Index(i)
			var domain string = ""
			domainValue := value.Get("domain")
			if !domainValue.IsNull() && !domainValue.IsUndefined() {
				domain = domainValue.String()
			}

			var expires int64 = 0
			expiresValue := value.Get("expires")
			if !expiresValue.IsNull() && !expiresValue.IsUndefined() {
				expires = int64(expiresValue.Float())
			}

			var path string = ""
			pathValue := value.Get("path")
			if !pathValue.IsNull() && !pathValue.IsUndefined() {
				path = pathValue.String()
			}

			var sameSite SameSite = Strict
			sameSiteValue := value.Get("sameSite")
			if !sameSiteValue.IsNull() && !sameSiteValue.IsUndefined() {
				sameSite = SameSite(sameSiteValue.String())
			}
			cookies = append(cookies, Cookie{
				Domain:      domain,
				Expires:     expires,
				Name:        value.Get("name").String(),
				Partitioned: value.Get("partitioned").Bool(),
				Path:        path,
				SameSite:    sameSite,
				Secure:      value.Get("secure").Bool(),
				Value:       value.Get("value").String(),
			})
		}

		resultChan <- Result{
			cookies: cookies,
			err:     nil,
		}

		return nil
	})

	// on_failure represents the callback that catches an exception when one is thrown
	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]

		resultChan <- Result{
			cookies: nil,
			err:     errors.New(value.Get("message").String()),
		}

		return nil
	})

	go cs.Value.Call("getAll", wrappedGetOptions).Call("then", on_success).Call("catch", on_failure)

	// block and wait for results
	result := <-resultChan
	return result.cookies, result.err
}

func (cs *cookieStore) Set(setOptions SetOptions) error {
	errChan := make(chan error)
	wrappedSetOptions := js.ValueOf(setOptions.MapToJS())

	// on_success represents the callback if the promise is resolved
	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]
		if value.IsUndefined() {
			errChan <- nil
		} else {
			errChan <- errors.New("onSuccess callback called, but the result should have been undefined. Check out the docs: https://developer.mozilla.org/en-US/docs/Web/API/CookieStore/set#return_value")
		}
		return nil
	})

	// on_failure represents the callback that catches an exception when one is thrown
	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {
		value := args[0]
		errChan <- errors.New(value.Get("message").String())
		return nil
	})

	go cs.Value.Call("set", wrappedSetOptions).Call("then", on_success).Call("catch", on_failure)

	// block and wait for error
	return <-errChan
}
