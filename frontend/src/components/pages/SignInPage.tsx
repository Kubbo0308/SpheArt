'use client'

import { useForm, FormProvider, useWatch } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button, Container } from '@chakra-ui/react'
import { FormInput } from '@/components/atoms/FormInput'
import { SignIn } from '@/api/user'
import { CONST, STATUS_CODE } from '@/const'
import { SignInFormSchema, SignInFormType } from '@/schemas/SignInFormSchema'
import { useCookies } from 'next-client-cookies'
import { useRouter } from 'next/navigation'

interface FormData {
  email: string
  password: string
}

export const SignInPage = () => {
  const cookies = useCookies()
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

  const onSubmit = async (params: FormData) => {
    const { email, password } = params
    const { data, status } = await SignIn(email, password)
    switch (status) {
      case STATUS_CODE.OK:
        // ログイン成功時
        // cookies.set('token', data)
        alert('login')
        router.push(CONST.TOP)
        break // 成功時の処理が完了したらbreakを忘れずに
      default:
        alert(status)
        break
    }
  }

  return (
    <Container size="md">
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <FormInput
            type="email"
            register={register('email')}
            label="メールアドレス"
            placeholder="example@example.com"
            required={true}
            errMessage={errors.email?.message}
          />
          <FormInput
            type="password"
            register={register('password')}
            label="パスワード"
            placeholder="パスワードを入力"
            required={true}
            errMessage={errors.password?.message}
          />
          <Button type="submit" isDisabled={isDisabled()}>
            ログイン
          </Button>
        </form>
      </FormProvider>
    </Container>
  )
}
