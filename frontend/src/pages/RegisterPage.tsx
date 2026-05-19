import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { Link } from 'react-router-dom'
import { useRegister } from '@/hooks/useAuth'

const schema = z.object({
  name: z.string().min(2, 'Name must be at least 2 characters'),
  email: z.string().email('Enter a valid email'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
})

type FormData = z.infer<typeof schema>

export default function RegisterPage() {
  const register_ = useRegister()
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>({
    resolver: zodResolver(schema),
  })

  const onSubmit = (data: FormData) => register_.mutate(data)

  return (
    <div style={styles.container}>
      <div style={styles.card}>
        <h1 style={styles.title}>Create Account</h1>

        <form onSubmit={handleSubmit(onSubmit)} style={styles.form}>
          <div style={styles.field}>
            <label style={styles.label}>Full Name</label>
            <input {...register('name')} style={styles.input} placeholder="Alice Smith" />
            {errors.name && <span style={styles.error}>{errors.name.message}</span>}
          </div>

          <div style={styles.field}>
            <label style={styles.label}>Email</label>
            <input {...register('email')} type="email" style={styles.input} placeholder="you@example.com" />
            {errors.email && <span style={styles.error}>{errors.email.message}</span>}
          </div>

          <div style={styles.field}>
            <label style={styles.label}>Password</label>
            <input {...register('password')} type="password" style={styles.input} placeholder="At least 8 characters" />
            {errors.password && <span style={styles.error}>{errors.password.message}</span>}
          </div>

          {register_.error && (
            <div style={styles.errorBox}>
              {(register_.error as { response?: { data?: { error?: string } } })?.response?.data?.error ?? 'Something went wrong. Please try again.'}
            </div>
          )}

          <button type="submit" disabled={register_.isPending} style={styles.button}>
            {register_.isPending ? 'Creating account...' : 'Create Account'}
          </button>
        </form>

        <p style={styles.footer}>
          Already have an account?{' '}
          <Link to="/login" style={styles.link}>Sign in</Link>
        </p>
      </div>
    </div>
  )
}

const styles: Record<string, React.CSSProperties> = {
  container: { minHeight: '100vh', display: 'flex', alignItems: 'center', justifyContent: 'center', background: '#f5f5f5' },
  card: { background: '#fff', padding: '2rem', borderRadius: '8px', boxShadow: '0 2px 8px rgba(0,0,0,0.1)', width: '100%', maxWidth: '400px' },
  title: { margin: '0 0 1.5rem', fontSize: '1.5rem', fontWeight: 700, textAlign: 'center' },
  form: { display: 'flex', flexDirection: 'column', gap: '1rem' },
  field: { display: 'flex', flexDirection: 'column', gap: '0.25rem' },
  label: { fontSize: '0.875rem', fontWeight: 500 },
  input: { padding: '0.625rem 0.75rem', border: '1px solid #d1d5db', borderRadius: '6px', fontSize: '1rem', outline: 'none' },
  error: { color: '#ef4444', fontSize: '0.75rem' },
  errorBox: { background: '#fef2f2', border: '1px solid #fecaca', color: '#b91c1c', padding: '0.75rem', borderRadius: '6px', fontSize: '0.875rem' },
  button: { padding: '0.75rem', background: '#2563eb', color: '#fff', border: 'none', borderRadius: '6px', fontSize: '1rem', fontWeight: 600, cursor: 'pointer', marginTop: '0.5rem' },
  footer: { textAlign: 'center', marginTop: '1.25rem', fontSize: '0.875rem', color: '#6b7280' },
  link: { color: '#2563eb', textDecoration: 'none', fontWeight: 500 },
}
