import { SignOut } from '@/api/user'
import { CONST, STATUS_CODE } from '@/const'
import { useRouter } from 'next/navigation'

interface returnValue {
  onSignOut: () => void
}

export const useAuthButton = (): returnValue => {
  const router = useRouter()
  const onSignOut = async () => {
    const { status } = await SignOut()
    switch (status) {
      case STATUS_CODE.OK:
        alert('logout')
        break // 成功時の処理が完了したらbreakを忘れずに
      default:
        alert(status)
        break
    }
    router.push(CONST.TOP)
  }

  return { onSignOut }
}
