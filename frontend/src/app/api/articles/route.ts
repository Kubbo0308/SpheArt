import { CONST } from "@/const"

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)
  const per_page = searchParams.get('per_page')

  const res = await fetch(`${process.env.API_URL}/api${CONST.ARTICLES}?per_page=${per_page}`, { next: { revalidate: 43200 } })
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