'use client'

import { SignOut } from '@/api/user'
import { CONST, STATUS_CODE } from '@/const'
import { Button } from '@chakra-ui/react'
import { useRouter } from 'next/navigation'

export default function SignOutPage() {
  const router = useRouter()

  const signOut = async () => {
    const { status } = await SignOut()
    switch (status) {
      case STATUS_CODE.OK:
        break // 成功時の処理が完了したらbreakを忘れずに
      default:
        alert(status)
        break
    }
    router.push(CONST.TOP)
  }

  return <Button onClick={signOut}>ログアウト</Button>
}
