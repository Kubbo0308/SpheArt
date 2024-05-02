import { CONST } from "@/const"

export async function POST(request: Request) {
  const res = await fetch(`${process.env.API_URL}/api${CONST.SIGN_OUT}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: "include", // Cookieを含める
  })

  // クッキーを削除するための設定
  const cookieHeader = `token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT; Max-Age=0`;

  // Responseオブジェクトを生成してステータスコードとデータを返します
  return new Response(JSON.stringify({
    status: res.status
  }), {
    status: 200,  // 成功した処理のHTTPステータスコード
    headers: {
      'Content-Type': 'application/json',
      'Set-Cookie': cookieHeader // クライアントにCookieを設定
    }
  });
}
