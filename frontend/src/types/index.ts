export interface User {
  id: string
  name: string
  email: string
  created_at: string
}

export interface ApiResponse<T> {
  success: boolean
  data?: T
  message?: string
  error?: string
}

export interface AuthTokens {
  access_token: string
  user: User
}
