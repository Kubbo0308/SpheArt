import { CONST } from "@/const";

export async function POST(
  request: Request,
  { params }: { params: { articleId: string } }
) {
  const articleId = params.articleId
  // クライアントからのリクエストヘッダーからCookieを取得
  const clientCookies = request.headers.get('cookie');
  // headersオブジェクトを定義し、Cookieがnullでない場合のみCookieヘッダーを設定
  const headers: HeadersInit = {
    'Content-Type': 'application/json'
  };
  if (clientCookies) {
    headers.Cookie = clientCookies; // Cookieがnullでなければヘッダーに追加
  }

  const res = await fetch(`${process.env.API_URL}/api${CONST.BOOKMARK}/${articleId}`,
  {
    method: "POST",
    headers: headers,
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