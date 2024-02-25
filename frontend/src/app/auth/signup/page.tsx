'use client'

import { AuthFormSchema, AuthFormType } from '@/schemas/AuthFormSchema'
import { useForm, FormProvider } from 'react-hook-form'
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
    handleSubmit
  } = methods

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
          <Button type="submit">新規登録</Button>
        </form>
      </FormProvider>
    </Container>
  )
}
