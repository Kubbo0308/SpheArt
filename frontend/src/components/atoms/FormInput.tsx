import { Box, FormControl, FormErrorMessage, Input } from '@chakra-ui/react'
import { HTMLInputTypeAttribute } from 'react'
import { UseFormRegisterReturn } from 'react-hook-form'
import { FormLabel } from './FormLabel'

interface FormInputProps {
  label: string
  type?: HTMLInputTypeAttribute
  placeholder?: string
  register?: UseFormRegisterReturn
  errMessage?: string
}

export const FormInput = (props: FormInputProps) => {
  const { label, type = 'text', placeholder, register, errMessage } = props
  return (
    <Box>
      <FormControl isInvalid={!!errMessage}>
        <FormLabel label={label} />
        <Input type={type} {...register} placeholder={placeholder} mt="5px" />
        {!!errMessage && <FormErrorMessage>{errMessage}</FormErrorMessage>}
      </FormControl>
    </Box>
  )
}
