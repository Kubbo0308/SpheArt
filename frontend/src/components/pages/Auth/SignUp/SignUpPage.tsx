'use client'

import { FormProvider } from 'react-hook-form'
import { Button, Container } from '@chakra-ui/react'
import { FormInput } from '@/components/atoms/FormInput'
import { useSignUpPage } from './SignUpPage.hooks'

export const SignUpPage = () => {
  const { methods, handleSubmit, onSubmit, register, errors, isDisabled } = useSignUpPage()

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
