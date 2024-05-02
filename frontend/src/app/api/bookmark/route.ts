import { CONST } from "@/const"

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)
  const per_page = searchParams.get('per_page')
  // クライアントからのリクエストヘッダーからCookieを取得
  const clientCookies = request.headers.get('cookie');
  // headersオブジェクトを定義し、Cookieがnullでない場合のみCookieヘッダーを設定
  const headers: HeadersInit = {
    'Content-Type': 'application/json'
  };
  if (clientCookies) {
    headers.Cookie = clientCookies; // Cookieがnullでなければヘッダーに追加
  }

  const res = await fetch(`${process.env.API_URL}/api${CONST.BOOKMARK}?per_page=${per_page}`,
  {
    method: "GET",
    headers: headers,
    cache: 'no-cache',
    credentials: "include", // Cookieを含める
  })
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
