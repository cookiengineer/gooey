//go:build wasm

package fetch

type ReferrerPolicy string

const (
	ReferrerPolicyNone                        ReferrerPolicy = "no-referrer"
	ReferrerPolicyNoDowngrade                 ReferrerPolicy = "no-referrer-when-downgrade"
	ReferrerPolicyOrigin                      ReferrerPolicy = "origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicy = "origin-when-cross-origin"
	ReferrerPolicySameOrigin                  ReferrerPolicy = "same-origin"
	ReferrerPolicyStrictOrigin                ReferrerPolicy = "strict-origin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strict-origin-when-cross-origin"
	ReferrerPolicyUnsafeURL                   ReferrerPolicy = "unsafe-url"
)
