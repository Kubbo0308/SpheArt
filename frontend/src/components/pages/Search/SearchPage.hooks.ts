import { ArticleProps } from "@/components/molecules/ArticleCard/ArticleCard"
import { CONST, STATUS_CODE } from "@/const"
import { useSearchParams } from "next/navigation"
import { MutableRefObject, useCallback, useEffect, useRef, useState } from "react"

interface returnValue {
  articles: ArticleProps[]
  loader: MutableRefObject<HTMLDivElement | null>
  isVisible: boolean
}

export const useSearchPageHooks = (): returnValue => {
  const [ offset, setOffset ] = useState(1)
  const [ articles, setArticles ] = useState<ArticleProps[]>([])
  const loader = useRef<HTMLDivElement | null>(null)
  const [ isVisible, setIsVisible ] = useState(true)
  const searchParams = useSearchParams()
  const title = searchParams.get('title')

  const handleObserver = useCallback((entities: IntersectionObserverEntry[]) => {
    const target = entities[0]
    if (target.isIntersecting) {
      setOffset((prev) => prev + 1)
    }
  }, [])

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api${CONST.ARTICLES}/search?title=${title}&per_page=${offset}`)
      if (response.ok) {
        const result = await response.json()
        switch (result.status) {
          case STATUS_CODE.OK:
            if (result.data.length === 0) {
              setIsVisible(false)
            }
            setArticles((prev) => [...prev, ...result.data])
            break // 成功時の処理が完了したらbreakを忘れずに
          default:
            alert(result.status)
            break
        }
      } else {
        alert(response.statusText);
      }
    }

    fetchData()
  }, [title, offset])

  useEffect(() => {
    const observer = new IntersectionObserver(handleObserver, {
      root: null,
      rootMargin: "20px",
      threshold: 0.5
    });
    if (loader.current) observer.observe(loader.current);

    return () => {
      observer.disconnect();
    };
  }, [handleObserver]);

  return { articles, loader, isVisible }
}
