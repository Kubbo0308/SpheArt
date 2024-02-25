'use client'

import { AuthFormSchema, AuthFormType } from '@/schemas/AuthFormSchema'
import { useForm, FormProvider, useWatch } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button, Container } from '@chakra-ui/react'
import { FormInput } from '@/components/atoms/FormInput'
import { UserInfo } from '@/api/user'

export default function SignUp() {
  const methods = useForm<AuthFormType>({
    mode: 'onChange',
    reValidateMode: 'onChange',
    resolver: zodResolver(AuthFormSchema),
    defaultValues: {
      email: '',
      password: '',
      confirmPassword: ''
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

  const watchConfirmPassword = useWatch({
    control,
    name: 'confirmPassword'
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
    if (getFieldState('confirmPassword').invalid || !watchConfirmPassword) {
      isDisabled = true
    }
    return isDisabled
  }

  const onSubmit = async (params: UserInfo) => {
    console.log(params)
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
          <FormInput
            type="password"
            register={register('confirmPassword')}
            label="確認用パスワード"
            placeholder="パスワードを再入力"
            required={true}
            errMessage={errors.confirmPassword?.message}
          />
          <Button type="submit" isDisabled={isDisabled()}>
            新規登録
          </Button>
        </form>
      </FormProvider>
    </Container>
  )
}
