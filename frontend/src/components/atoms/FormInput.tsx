import { Box, FormControl, FormErrorMessage, Input } from '@chakra-ui/react'
import { HTMLInputTypeAttribute } from 'react'
import { UseFormRegisterReturn } from 'react-hook-form'
import { FormLabel } from './FormLabel'

interface FormInputProps {
  label: string
  required: boolean
  type?: HTMLInputTypeAttribute
  placeholder?: string
  register?: UseFormRegisterReturn
  errMessage?: string
}

export const FormInput = (props: FormInputProps) => {
  const { label, required, type = 'text', placeholder, register, errMessage } = props
  return (
    <Box>
      <FormControl isInvalid={!!errMessage}>
        <FormLabel label={label} required={required} />
        <Input type={type} {...register} placeholder={placeholder} />
        {!!errMessage && <FormErrorMessage>{errMessage}</FormErrorMessage>}
      </FormControl>
    </Box>
  )
}
