import { CONST } from "@/const"

export async function GetAllBookmark(userId: string) {
  const res = await fetch(`http://localhost:8080${CONST.BOOKMARK}/${userId}`, {
    method: "GET",
    headers: {
      'Content-Type': 'application/json',
    },
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function PostBookmark(userId: string, articleId: string) {
  const res = await fetch(`http://localhost:8080${CONST.BOOKMARK}/${userId}/${articleId}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
  })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function DeleteBookmark(userId: string, articleId: string) {
  const res = await fetch(`http://localhost:8080${CONST.BOOKMARK}/${userId}/${articleId}`, {
    method: "DELETE",
    headers: {
      'Content-Type': 'application/json',
    },
  })
  const data = await res.json()
  return { status: res.status, data: data }
}
