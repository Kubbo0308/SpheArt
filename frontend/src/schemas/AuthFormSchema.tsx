import { z } from 'zod'

export const AuthFormSchema = z
  .object({
    email: z
      .string()
      .min(1, { message: 'メールアドレスを入力して下さい。' })
      .email({ message: 'メールアドレスの形式で入力して下さい。' })
      .max(100, { message: 'メールアドレスの形式で入力して下さい。' }),
    password: z
      .string()
      .min(1, { message: 'パスワードを入力して下さい。' })
      .max(100, { message: 'パスワードは100文字以内で入力して下さい。' }),
    confirmPassword: z
      .string()
      .min(1, { message: 'パスワードを再入力して下さい。' })
      .max(100, { message: 'パスワードは100文字以内で入力して下さい。' })
  })
  .superRefine((data, ctx) => {
    if (data.password !== data.confirmPassword) {
      ctx.addIssue({
        message: 'パスワードが異なります。',
        path: ['confirmPassword'],
        code: z.ZodIssueCode.custom
      })
    }
  })

export type AuthFormType = z.infer<typeof AuthFormSchema>
