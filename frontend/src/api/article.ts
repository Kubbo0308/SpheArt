import { CONST } from "@/const";

export async function getArticles(pageNum: number) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.ARTICLES}?per_page=${pageNum}`, { next: { revalidate: 3600 } })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function searchArticlesInTitle(title: string) {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.SEARCH_ARTICLES_IN_TITLE}${title}`, { next: { revalidate: 3600 } })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  const data = await res.json()
  return { status: res.status, data: data }
}