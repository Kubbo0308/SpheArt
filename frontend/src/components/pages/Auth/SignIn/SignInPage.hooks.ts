import { FieldErrors, useForm, UseFormHandleSubmit, UseFormRegister, UseFormReturn, useWatch } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { CONST, STATUS_CODE } from '@/const'
import { SignInFormSchema, SignInFormType } from '@/schemas/SignInFormSchema'
import { useRouter } from 'next/navigation'

interface returnValue {
  methods: UseFormReturn<SignInFormType>
  handleSubmit: UseFormHandleSubmit<SignInFormType>
  onSubmit: (params: SignInFormType) => Promise<void>
  register: UseFormRegister<SignInFormType>
  errors: FieldErrors<SignInFormType>
  isDisabled: () => boolean
}

export const useSignInPage = (): returnValue => {
  const router = useRouter()
  const methods = useForm<SignInFormType>({
    mode: 'onChange',
    reValidateMode: 'onChange',
    resolver: zodResolver(SignInFormSchema),
    defaultValues: {
      email: '',
      password: ''
    }
  })

  const {
    register,
    formState: { errors },
    handleSubmit,
    getFieldState,
    control
  } = methods

  const watchEmail = useWatch({
    control,
    name: 'email'
  })

  const watchPassword = useWatch({
    control,
    name: 'password'
  })

  // 必須入力の項目が全て正しく入力されているかチェック
  const isDisabled = (): boolean => {
    let isDisabled = false
    if (getFieldState('email').invalid || !watchEmail) {
      isDisabled = true
    }
    if (getFieldState('password').invalid || !watchPassword) {
      isDisabled = true
    }
    return isDisabled
  }

  const onSubmit = async (params: SignInFormType) => {
    const { email, password } = params
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api${CONST.AUTH}${CONST.SIGN_IN}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: email,
        password: password
      }),
      credentials: "include", // Cookieを含める
    })
    if (response.ok) {
      const result = await response.json();
      switch (result.status) {
        case STATUS_CODE.OK:
          // ログイン成功時
          alert('login')
          router.push(CONST.TOP)
          window.location.reload()
          break // 成功時の処理が完了したらbreakを忘れずに
        case STATUS_CODE.UNAUTHORIZED:
          alert('メールアドレスかパスワードが間違っています😭')
          break
        default:
          alert(result.status)
          break
      }
    }
  }

  return { methods, handleSubmit, onSubmit, register, errors, isDisabled }
}
