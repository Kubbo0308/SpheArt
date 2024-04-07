import { CONST } from '@/const';
import { STATUS_CODE } from '@/const'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { PostBookmark } from '@/api/bookmark';

interface returnValue {
  isBookmark: boolean
  postBookmark: (articleId: string) => Promise<void>
  formatDate: (dateString: string) => string
}

export const useArticleCard = (token: RequestCookie | undefined): returnValue => {
  const isLogin = token !== undefined
  const [isBookmark, setIsBookmark] = useState(false)
  const router = useRouter()

  const postBookmark = async (articleId: string) => {
    if (isLogin) {
    const { status } = await PostBookmark(articleId)
    switch (status) {
      case STATUS_CODE.OK:
        setIsBookmark(!isBookmark)
        break
      default:
        break
    }
  } else {
    alert("ブックマークするにはログインしてください。")
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