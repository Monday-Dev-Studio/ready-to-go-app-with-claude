import { useAuthStore } from '@/store/authStore'
import { useLogout } from '@/hooks/useAuth'

export default function HomePage() {
  const { user } = useAuthStore()
  const logout = useLogout()

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <h1 style={styles.greeting}>Hello, {user?.name} 👋</h1>
        <button onClick={() => logout.mutate()} style={styles.logoutBtn}>
          Sign Out
        </button>
      </div>

      <div style={styles.content}>
        <div style={styles.card}>
          <h2 style={styles.cardTitle}>Your app is ready!</h2>
          <p style={styles.cardText}>
            This is the home page of your app. Ask Claude to build features here.
          </p>
          <p style={styles.cardText}>
            Example: <em>"Add a task list where I can create, complete, and delete tasks"</em>
          </p>
        </div>
      </div>
    </div>
  )
}

const styles: Record<string, React.CSSProperties> = {
  container: { minHeight: '100vh', background: '#f9fafb' },
  header: { background: '#fff', borderBottom: '1px solid #e5e7eb', padding: '1rem 2rem', display: 'flex', alignItems: 'center', justifyContent: 'space-between' },
  greeting: { margin: 0, fontSize: '1.25rem', fontWeight: 600 },
  logoutBtn: { padding: '0.5rem 1rem', background: 'transparent', border: '1px solid #d1d5db', borderRadius: '6px', cursor: 'pointer', fontSize: '0.875rem' },
  content: { padding: '2rem', maxWidth: '800px', margin: '0 auto' },
  card: { background: '#fff', border: '1px solid #e5e7eb', borderRadius: '8px', padding: '2rem' },
  cardTitle: { margin: '0 0 1rem', fontSize: '1.25rem', fontWeight: 600 },
  cardText: { margin: '0 0 0.75rem', color: '#4b5563', lineHeight: 1.6 },
}
