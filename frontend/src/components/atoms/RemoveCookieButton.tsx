'use client'

import { useCookies } from 'next-client-cookies'
import { LinkButton } from './LinkButton'

interface RemoveCookieButtonProps {
  title: string
  url: string
}

export const RemoveCookieButton = (props: RemoveCookieButtonProps) => {
  const { title, url } = props

  const cookies = useCookies()
  return <LinkButton title={title} url={url} onClick={() => cookies.remove('token')}></LinkButton>
}
