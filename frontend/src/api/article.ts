import { CONST } from "@/const";

export async function getArticles(pageNum: number) {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}${CONST.ARTICLES}?per_page=${pageNum}`, { next: { revalidate: 3600 } })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function searchArticlesInTitle(title: string, pageNum: number) {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}${CONST.SEARCH_ARTICLES_IN_TITLE}${title}&per_page=${pageNum}`, { next: { revalidate: 3600 } })
  const data = await res.json()
  return { status: res.status, data: data }
}