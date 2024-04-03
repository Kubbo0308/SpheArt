import { SignInPage } from '@/components/pages/Auth/SignIn/SignInPage'
import { CONST } from '@/const'
import { Link } from '@chakra-ui/react'

export default function SignIn() {
  return (
    <>
      <SignInPage />
      <Link href={`${CONST.AUTH}${CONST.SIGN_UP}`}>新規登録はこちら</Link>
    </>
  )
}
