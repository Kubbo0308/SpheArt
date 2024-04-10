'use client'

import { useAuthButton } from './AuthButton.hooks'
import { LinkButton } from '@/components/atoms/LinkButton'
import { CONST } from '@/const'

export const AuthButton = () => {
  const { onSignOut, token } = useAuthButton()

  return (
    <>
      {token === undefined ? (
        <LinkButton title="Sign In" url={`${CONST.AUTH}${CONST.SIGN_IN}`} />
      ) : (
        <LinkButton title="サインアウト" onClick={onSignOut} />
      )}
    </>
  )
}
