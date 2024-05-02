import { CONST } from "@/const"

export async function POST(request: Request) {
  const requestData = await request.json()
  const res = await fetch(`${process.env.API_URL}/api${CONST.SIGN_UP}`, {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: requestData.email,
      password: requestData.password
    })
  })

  // Responseオブジェクトを生成してステータスコードとデータを返します
  return new Response(JSON.stringify({
    status: res.status
  }), {
    status: 200,  // 成功した処理のHTTPステータスコード
    headers: {
      'Content-Type': 'application/json'
    }
  });
}
