import { CONST } from "@/const"

export async function fetchCsrfToken() {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.CSRF_TOKEN}`, {
    method: "GET",
    credentials: "include",
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return data.csrf_token;
}