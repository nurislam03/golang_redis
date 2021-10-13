# magic.pathao.com/veritas/auth

Warden REST API.

## Routes

<details>
<summary>`/api/v1/profile`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/api**
	- **/v1/profile**
		- [(*ProfileResource).profileCtx-fm]()
		- **/**
			- _GET_
				- [(*ProfileResource).get-fm]()
			- _PUT_
				- [(*ProfileResource).update-fm]()

</details>
<details>
<summary>`/api/v1/user`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/api**
	- **/v1/user**
		- [(*UserResource).userCtx-fm]()
		- **/**
			- _PUT_
				- [(*UserResource).update-fm]()
			- _DELETE_
				- [(*UserResource).delete-fm]()
			- _GET_
				- [(*UserResource).get-fm]()

</details>
<details>
<summary>`/api/v1/user/token/{tokenID}`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/api**
	- **/v1/user**
		- [(*UserResource).userCtx-fm]()
		- **/token/{tokenID}**
			- **/**
				- _PUT_
					- [(*UserResource).updateToken-fm]()
				- _DELETE_
					- [(*UserResource).deleteToken-fm]()

</details>
<details>
<summary>`/auth/challenges/resend`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/auth**
	- [SetContentType.func1]()
	- **/challenges**
		- **/resend**
			- _PATCH_
				- [(*Resource).resendOTP-fm]()

</details>
<details>
<summary>`/auth/challenges/validate`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/auth**
	- [SetContentType.func1]()
	- **/challenges**
		- **/validate**
			- _PATCH_
				- [(*Resource).validateOTP-fm]()

</details>
<details>
<summary>`/auth/login`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/auth**
	- [SetContentType.func1]()
	- **/login**
		- _POST_
			- [(*Resource).login-fm]()

</details>
<details>
<summary>`/auth/logout`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/auth**
	- [SetContentType.func1]()
	- **/logout**
		- _POST_
			- [v5.Verifier.func1]()
			- [AuthenticateRefreshJWT]()
			- [(*Resource).logout-fm]()

</details>
<details>
<summary>`/auth/refresh`</summary>

- [Recoverer]()
- [RequestID]()
- [RealIP]()
- [Timeout.func1]()
- [Heartbeat.func1]()
- [RequestLogger.func1]()
- [SetContentType.func1]()
- **/auth**
	- [SetContentType.func1]()
	- **/refresh**
		- _POST_
			- [v5.Verifier.func1]()
			- [AuthenticateRefreshJWT]()
			- [(*Resource).refresh-fm]()

</details>

Total # of routes: 8
