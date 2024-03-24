import { cookies } from 'next/headers';
import { CONST } from "@/const"

export async function GetAllBookmark() {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.BOOKMARK}`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json',
      Cookie: cookies().toString(),
    },
    cache: 'no-cache',
    credentials: "include", // Cookieを含める
  })
  const data = await res.json()
  return { status: res.status, data: data }
}