"use client"

import { getArticles } from "@/api/article"
import { ArticleProps } from "@/components/atoms/ArticleListItem"
import { useEffect, useState } from "react"

interface returnValue {
  currentPage: number
  articles: ArticleProps[]
  goNextPage: () => void
  backPreviousPage: () => void
}

export const useTopPageHooks = (): returnValue => {
  const [ currentPage, setCurrentPage ] = useState(1)
  const [ articles, setArticles ] = useState<ArticleProps[]>([])

  useEffect(() => {
    const fetchData = async () => {
      const response = await getArticles(currentPage)
      setArticles(response)
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
