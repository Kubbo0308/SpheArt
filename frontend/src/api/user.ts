import { CONST } from "@/const"

export async function SignUp(email: string, password: string) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_UP}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function SignIn(email: string, password: string) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_IN}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: email,
      password: password
    }),
    credentials: "include", // Cookieを含める
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function SignOut() {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SIGN_OUT}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return { status: res.status }
}
