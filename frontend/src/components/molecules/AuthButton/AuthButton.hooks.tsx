import { SignOut } from '@/api/user'
import { CONST, STATUS_CODE } from '@/const'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { serverAuthButton } from './AuthButton.server'

interface returnValue {
  onSignOut: () => void
  token: RequestCookie | undefined
}

export const useAuthButton = (): returnValue => {
  const router = useRouter()
  const [token, setToken] = useState<RequestCookie | undefined>(undefined)

  useEffect(() => {
    const fetchToken = async () => {
      const tokenString = await serverAuthButton()
      setToken(tokenString)
    }
    fetchToken()
  }, [])

  const onSignOut = async () => {
    const { status } = await SignOut()
    switch (status) {
      case STATUS_CODE.OK:
        alert('logout')
        router.push(CONST.TOP)
        window.location.reload()
        break // 成功時の処理が完了したらbreakを忘れずに
      default:
        alert(status)
        break
    }
  }

  return { onSignOut, token }
}
