
import { GetBookmark } from "@/api/bookmark"
import { ArticleProps } from "@/components/molecules/ArticleCard/ArticleCard"
import { STATUS_CODE } from "@/const"
import { useEffect, useState } from "react"

interface returnValue {
  currentPage: number
  articles: ArticleProps[]
  goNextPage: () => void
  backPreviousPage: () => void
}

export const useBookmarkPageHooks = (): returnValue => {
  const [ currentPage, setCurrentPage ] = useState(1)
  const [ articles, setArticles ] = useState<ArticleProps[]>([])

  useEffect(() => {
    const fetchData = async () => {
      const { data, status } = await GetBookmark(currentPage)
      switch (status) {
        case STATUS_CODE.OK:
          setArticles(data)
          break // 成功時の処理が完了したらbreakを忘れずに
        default:
          alert(status)
          break
      }
    }

    fetchData()
  }, [currentPage])

  const goNextPage = () => {
    setCurrentPage(currentPage + 1)
  }

  const backPreviousPage = () => {
    setCurrentPage(currentPage - 1)
  }

  return { currentPage, articles, goNextPage, backPreviousPage }
}
