'use client'

import { FormProvider } from 'react-hook-form'
import { Button, Container } from '@chakra-ui/react'
import { FormInput } from '@/components/atoms/FormInput'
import { useSignInPage } from './SingInPage.hooks'

export const SignInPage = () => {
  const { methods, handleSubmit, onSubmit, register, errors, isDisabled } = useSignInPage()
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
