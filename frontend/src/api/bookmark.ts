import { CONST } from "@/const"

export async function GetBookmark(pageNum: number) {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}${CONST.BOOKMARK}?per_page=${pageNum}`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json',
    },
    cache: 'no-cache',
    credentials: "include", // Cookieを含める
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function PostBookmark(articleId: string) {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}${CONST.BOOKMARK}/${articleId}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })
  const data = await res.json()
  return { status: res.status, data: data }
}
