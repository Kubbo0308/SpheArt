import { CONST } from '@/const';
import { STATUS_CODE } from '@/const'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { useState } from 'react'
import { useRouter } from 'next/navigation'

interface useArticleCardProps {
  token: RequestCookie | undefined
  isBookmarkPage: boolean
}

interface returnValue {
  isBookmark: boolean
  postBookmark: (articleId: string) => Promise<void>
  formatDate: (dateString: string) => string
}

export const useArticleCard = (props: useArticleCardProps): returnValue => {
  const { token, isBookmarkPage } = props
  const isLogin = token !== undefined
  const [isBookmark, setIsBookmark] = useState(isBookmarkPage)
  const router = useRouter()

  const postBookmark = async (articleId: string) => {
    if (isLogin) {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api${CONST.BOOKMARK}/${articleId}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        credentials: "include", // Cookieを含める
      })
      if (response.ok) {
        const result = await response.json()
        switch (result.status) {
        case STATUS_CODE.OK:
          setIsBookmark(!isBookmark)
          break
        default:
          break
        }
      }
    } else {
      alert("ブックマークするにはサインインしてください。")
      router.push(`${CONST.AUTH}${CONST.SIGN_IN}`)
    }
  }

  // 日付のフォーマット
  const formatDate = (dateString: string) => {
    const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' }
    return new Date(dateString).toLocaleDateString(undefined, options)
  }

  return { isBookmark, postBookmark, formatDate }
}