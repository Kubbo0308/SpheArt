'use client'

import { FormProvider } from 'react-hook-form'
import { Box, Button, Container, Divider, Flex, Link, Text } from '@chakra-ui/react'
import NextLink from 'next/link'
import { FormInput } from '@/components/atoms/FormInput'
import { useSignInPage } from './SingInPage.hooks'
import { CONST } from '@/const'

export const SignInPage = () => {
  const { methods, handleSubmit, onSubmit, register, errors, isDisabled } = useSignInPage()
  return (
    <Container pt={{ base: '30px', md: '50px' }} pb="50px" px="20px">
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Container bg="white.primary" p="30px">
            <Text fontSize="24px" fontWeight={700} lineHeight={1.8} textAlign="center">
              Sign In
            </Text>
            <Flex direction="column" gap="15px" mt="20px">
              <FormInput
                type="email"
                register={register('email')}
                label="メールアドレス"
                placeholder="example@example.com"
                errMessage={errors.email?.message}
              />
              <FormInput
                type="password"
                register={register('password')}
                label="パスワード"
                placeholder="example"
                errMessage={errors.password?.message}
              />
            </Flex>
            <Button
              type="submit"
              isDisabled={isDisabled()}
              mt="20px"
              bg="blue.accent"
              borderRadius="6px"
              display="block"
              mx="auto"
              cursor="pointer"
            >
              <Text fontSize="16px" fontWeight={600} lineHeight={1.8} color="white.primary">
                ログイン
              </Text>
            </Button>
            <Divider mt="20px" />
            <Box mt="20px">
              <Link
                as={NextLink}
                href={`${CONST.AUTH}${CONST.SIGN_UP}`}
                textAlign="center"
                mt="20px"
                textDecoration="underline"
                color="blue.accent"
              >
                <Text fontSize="12px" fontWeight={600} lineHeight={1.8}>
                  新規登録はこちらから
                </Text>
              </Link>
            </Box>
          </Container>
        </form>
      </FormProvider>
    </Container>
  )
}
