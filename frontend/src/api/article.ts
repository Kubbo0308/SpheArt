import { CONST } from "@/const";

export async function getArticles() {
  const res = await fetch(`${CONST.API_BASE_PATH}${CONST.ARTICLES}`, { next: { revalidate: 3600 } })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return res.json()
}