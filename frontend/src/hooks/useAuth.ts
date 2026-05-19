import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { authService } from '@/services/api'
import { useAuthStore } from '@/store/authStore'
import type { AuthTokens, User } from '@/types'

export function useLogin() {
  const { setAuth } = useAuthStore()
  const navigate = useNavigate()
  const qc = useQueryClient()

  return useMutation({
    mutationFn: (data: { email: string; password: string }) => authService.login(data),
    onSuccess: (res) => {
      const { access_token, user } = res.data.data as AuthTokens
      setAuth(user, access_token)
      qc.invalidateQueries({ queryKey: ['me'] })
      navigate('/')
    },
  })
}

export function useRegister() {
  const navigate = useNavigate()

  return useMutation({
    mutationFn: (data: { name: string; email: string; password: string }) =>
      authService.register(data),
    onSuccess: () => navigate('/login'),
  })
}

export function useLogout() {
  const { logout } = useAuthStore()
  const navigate = useNavigate()
  const qc = useQueryClient()

  return useMutation({
    mutationFn: () => authService.logout(),
    onSettled: () => {
      logout()
      qc.clear()
      navigate('/login')
    },
  })
}

export function useMe() {
  const { user } = useAuthStore()

  return useQuery({
    queryKey: ['me'],
    queryFn: async () => {
      const res = await authService.me()
      return res.data.data as User
    },
    enabled: !!user,
    staleTime: 5 * 60 * 1000,
  })
}
