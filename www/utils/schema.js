import { z } from 'zod'

export const emailSchema = z.string().email('Must be a valid email')

export const usernameSchema = z
    .string()
    .min(3, 'At least 3 characters')
    .max(24, 'Less than 24 characters')
    .regex(
        /^[a-zA-Z0-9._-]+$/,
        'Alphanumeric, hyphen, underscore, or full stop only'
    )

export const nameSchema = z
    .string()
    .min(1, 'At least 1 character')
    .max(255, 'At most 255 characters')

export const passwordSchema = z
    .string()
    .min(8, 'At least 8 characters')
    .max(255, 'At most 255 characters')
    .regex(/[A-Z]/, '1 uppercase')
    .regex(/[a-z]/, '1 lowercase')
    .regex(/\d/, '1 digit')

export const userSchema = z.object({
    email: emailSchema,
    username: usernameSchema,
    name: nameSchema,
    password: passwordSchema,
})

export const matchingSchema = (data, name) =>
    z.string().refine((arg) => arg == data, name + ' must match')
