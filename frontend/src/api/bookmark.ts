import { CONST } from "@/const"

export async function GetBookmark(pageNum: number) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.BOOKMARK}?per_page=${pageNum}`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json',
    },
    cache: 'no-cache',
    credentials: "include", // Cookieを含める
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function PostBookmark(articleId: string) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.BOOKMARK}/${articleId}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return { status: res.status, data: data }
}