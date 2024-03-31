'use client'

import { useCookies } from 'next-client-cookies'
import { LinkButton } from './LinkButton'

interface RemoveCookieButtonProps {
  title: string
  url: string
}

export const RemoveCookieButton = (props: RemoveCookieButtonProps) => {
  const { title, url } = props

  return <LinkButton title={title} url={url}></LinkButton>
}
