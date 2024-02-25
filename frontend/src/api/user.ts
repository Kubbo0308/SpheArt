import { CONST } from "@/const"
import { fetchCsrfToken } from "./csrf"

export interface UserInfo {
  email: string
  password: string
}

export async function SignUp(props: UserInfo) {
  const {email, password} = props
  const csrfToken = await fetchCsrfToken()
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_UP}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
      'X-CSRF-Token': csrfToken, // ヘッダーにCSRFトークンをセット
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return res.json
}

export async function SignIn(props: UserInfo) {
  const {email, password} = props
  const csrfToken = await fetchCsrfToken()
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_IN}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
      'X-CSRF-Token': csrfToken, // ヘッダーにCSRFトークンをセット
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return
}

export async function SignOut(props: UserInfo) {
  const {email, password} = props
  const csrfToken = await fetchCsrfToken()
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_OUT}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
      'X-CSRF-Token': csrfToken, // ヘッダーにCSRFトークンをセット
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return
}
