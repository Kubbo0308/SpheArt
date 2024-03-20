import { CONST } from "@/const"

export async function GetAllBookmark() {
  const res = await fetch(`http://localhost:8080${CONST.BOOKMARK}`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function PostBookmark(articleId: string) {
  const res = await fetch(`http://localhost:8080${CONST.BOOKMARK}/${articleId}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function DeleteBookmark(articleId: string) {
  const res = await fetch(`http://localhost:8080${CONST.BOOKMARK}/${articleId}`, {
    method: "DELETE",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })
  const data = await res.json()
  return { status: res.status, data: data }
}
