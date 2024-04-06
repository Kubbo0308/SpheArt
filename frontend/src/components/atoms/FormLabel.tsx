import { Text } from '@chakra-ui/react'

interface FormLabelProps {
  label: string
}

export const FormLabel = (props: FormLabelProps) => {
  const { label } = props
  return (
    <Text fontSize="16px" fontWeight={600} lineHeight={1.8}>
      {label}
    </Text>
  )
}
