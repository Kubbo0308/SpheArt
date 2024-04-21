import { CONST } from "@/const";

export async function getArticles(pageNum: number) {
  const res = await fetch(`${CONST.API_BASE_PATH}/api${CONST.ARTICLES}?per_page=${pageNum}`, { next: { revalidate: 43200 } })
  const data = await res.json()
  return { status: res.status, data: data }
}

export async function searchArticlesInTitle(title: string, pageNum: number) {
  const res = await fetch(`${CONST.API_BASE_PATH}/api${CONST.SEARCH_ARTICLES_IN_TITLE}${title}&per_page=${pageNum}`, { next: { revalidate: 43200 } })
  const data = await res.json()
  return { status: res.status, data: data }
}