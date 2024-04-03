'use client'

import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { useAuthButton } from './AuthButton.hooks'
import { LinkButton } from '@/components/atoms/LinkButton'
import { CONST } from '@/const'

export const AuthButton = ({ token }: { token: RequestCookie | undefined }) => {
  const { onSignOut } = useAuthButton()

  return (
    <>
      {token === undefined ? (
        <LinkButton title="ログイン" url={`${CONST.AUTH}${CONST.SIGN_IN}`} />
      ) : (
        <LinkButton title="ログアウト" onClick={onSignOut} />
      )}
    </>
  )
}
