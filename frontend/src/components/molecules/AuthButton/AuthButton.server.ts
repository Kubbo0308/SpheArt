"use server"

import { cookies } from "next/headers"

export const serverAuthButton = async () => {
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return token
}