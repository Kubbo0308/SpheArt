'use client'

import { useAuthButton } from './AuthButton.hooks'
import { LinkButton } from '@/components/atoms/LinkButton'
import { CONST } from '@/const'
import { HeaderMenu } from '../HeaderMenu'

export const AuthButton = () => {
  const { onSignOut, token } = useAuthButton()

  return (
    <>
      {token === undefined ? (
        <LinkButton title="Sign In" url={`${CONST.AUTH}${CONST.SIGN_IN}`} />
      ) : (
        <HeaderMenu onClick={onSignOut} />
      )}
    </>
  )
}
