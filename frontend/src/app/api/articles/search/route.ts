import { CONST } from "@/const"

export async function GET(reqest: Request) {
  const { searchParams } = new URL(reqest.url)
  const title = searchParams.get('title')
  const per_page = searchParams.get('per_page')

  const res = await fetch(`${process.env.API_URL}/api${CONST.ARTICLES}${CONST.SEARCH}?title=${title}&per_page=${per_page}`, { next: { revalidate: 43200 } })
  const data = await res.json()

  // Responseオブジェクトを生成してステータスコードとデータを返します
  return new Response(JSON.stringify({
    status: res.status,
    data: data
  }), {
    status: 200,  // 成功した処理のHTTPステータスコード
    headers: {
      'Content-Type': 'application/json'
    }
  });
}