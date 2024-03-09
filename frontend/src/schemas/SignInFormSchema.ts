import { z } from 'zod'

export const SignInFormSchema = z
  .object({
    email: z
      .string()
      .min(1, { message: 'メールアドレスを入力して下さい。' })
      .email({ message: 'メールアドレスの形式で入力して下さい。' })
      .max(100, { message: 'メールアドレスの形式で入力して下さい。' }),
    password: z
      .string()
      .min(1, { message: 'パスワードを入力して下さい。' })
      .max(100, { message: 'パスワードは100文字以内で入力して下さい。' })
  })

export type SignInFormType = z.infer<typeof SignInFormSchema>
