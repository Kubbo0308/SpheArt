import { CONST } from "@/const"

export async function POST(request: Request) {
  const requestData = await request.json()
  const res = await fetch(`${process.env.API_URL}/api${CONST.SIGN_IN}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: requestData.email,
      password: requestData.password
    }),
    credentials: "include", // Cookieを含める
  })

  const cookies = res.headers.get('set-cookie');
  // Cookieが存在する場合のみヘッダーに設定
  let headersInit: HeadersInit = {
    'Content-Type': 'application/json'
  };

  if (cookies) {
    headersInit['Set-Cookie'] = cookies;
  }

  // Responseオブジェクトを生成してステータスコードとデータを返します
  return new Response(JSON.stringify({
    status: res.status
  }), {
    status: 200,  // 成功した処理のHTTPステータスコード
    headers: headersInit
  });
}
