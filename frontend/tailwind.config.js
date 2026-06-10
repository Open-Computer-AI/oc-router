/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // OC Design Language — semantic tokens via CSS custom properties (RGB channels)
        bg: {
          DEFAULT: 'rgb(var(--bg) / <alpha-value>)',
          elevated: 'rgb(var(--bg-elevated) / <alpha-value>)',
        },
        sidebar: 'rgb(var(--sidebar) / <alpha-value>)',
        text: {
          DEFAULT: 'rgb(var(--text) / <alpha-value>)',
          muted: 'rgb(var(--text-muted) / <alpha-value>)',
          dim: 'rgb(var(--text-dim) / <alpha-value>)',
        },
        border: {
          DEFAULT: 'rgb(var(--border) / <alpha-value>)',
        },
        muted: {
          DEFAULT: 'rgb(var(--muted) / <alpha-value>)',
        },
        accent: {
          DEFAULT: 'rgb(var(--accent) / <alpha-value>)',
          fg: 'rgb(var(--accent-fg) / <alpha-value>)',
          muted: 'rgb(var(--accent-muted) / 0.08)',
        },
        // Semantic status colors
        success: {
          DEFAULT: 'rgb(var(--success) / <alpha-value>)',
          muted: 'rgb(var(--success-muted) / 0.08)',
        },
        warning: {
          DEFAULT: 'rgb(var(--warning) / <alpha-value>)',
          muted: 'rgb(var(--warning-muted) / 0.08)',
        },
        danger: {
          DEFAULT: 'rgb(var(--danger) / <alpha-value>)',
          muted: 'rgb(var(--danger-muted) / 0.08)',
        },
        info: {
          DEFAULT: 'rgb(var(--info) / <alpha-value>)',
          muted: 'rgb(var(--info-muted) / 0.08)',
        },
        // Platform colors
        openai: '#10a37f',
        anthropic: '#d97757',
        google: '#4285f4',
        azure: '#0078d4',
        // Keep legacy Tailwind palette for gradual migration
        primary: {
          50: '#fce4ec',
          100: '#f8bbd0',
          200: '#f48fb1',
          300: '#f06292',
          400: '#ec407a',
          500: '#c2185b',
          600: '#ad1457',
          700: '#880e4f',
          800: '#6a0036',
          900: '#4a0025',
          950: '#2d0016',
        },
        dark: {
          50: '#f8fafc',
          100: '#f1f5f9',
          200: '#e2e8f0',
          300: '#cbd5e1',
          400: '#94a3b8',
          500: '#64748b',
          600: '#475569',
          700: '#334155',
          800: '#1e293b',
          900: '#0f172a',
          950: '#020617',
        },
      },
      fontFamily: {
        sans: [
          'Inter',
          'system-ui',
          '-apple-system',
          'BlinkMacSystemFont',
          'Segoe UI',
          'Roboto',
          'Helvetica Neue',
          'Arial',
          'PingFang SC',
          'Hiragino Sans GB',
          'Microsoft YaHei',
          'sans-serif',
        ],
        display: [
          'Newsreader',
          'Georgia',
          'Cambria',
          'Times New Roman',
          'Times',
          'serif',
        ],
        mono: [
          'JetBrains Mono',
          'ui-monospace',
          'SFMono-Regular',
          'Menlo',
          'Monaco',
          'Consolas',
          'monospace',
        ],
      },
      boxShadow: {
        card: '0 1px 3px rgba(0,0,0,0.04), 0 1px 2px rgba(0,0,0,0.06)',
        'card-hover': '0 10px 40px rgba(0,0,0,0.08)',
        float: '0 8px 30px rgba(0,0,0,0.08)',
        modal: '0 25px 50px -12px rgba(0,0,0,0.15)',
        glass: '0 8px 32px rgba(0,0,0,0.08)',
        'glass-sm': '0 4px 16px rgba(0,0,0,0.06)',
        glow: '0 0 20px rgba(194,24,91,0.25)',
        'glow-lg': '0 0 40px rgba(194,24,91,0.35)',
        'inner-glow': 'inset 0 1px 0 rgba(255,255,255,0.1)',
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-primary': 'linear-gradient(135deg, rgb(var(--accent)) 0%, #ad1457 100%)',
        'mesh-gradient':
          'radial-gradient(at 40% 20%, rgba(194,24,91,0.06) 0px, transparent 50%), radial-gradient(at 80% 0%, rgba(194,24,91,0.04) 0px, transparent 50%), radial-gradient(at 0% 50%, rgba(194,24,91,0.04) 0px, transparent 50%)',
      },
      animation: {
        'fade-in': 'fadeIn 0.3s ease-out',
        'slide-up': 'slideUp 0.3s ease-out',
        'slide-down': 'slideDown 0.3s ease-out',
        'slide-in-right': 'slideInRight 0.3s ease-out',
        'scale-in': 'scaleIn 0.2s ease-out',
        'pulse-slow': 'pulse 3s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        shimmer: 'shimmer 2s linear infinite',
        glow: 'glow 2s ease-in-out infinite alternate',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        slideUp: {
          '0%': { opacity: '0', transform: 'translateY(10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        slideDown: {
          '0%': { opacity: '0', transform: 'translateY(-10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        slideInRight: {
          '0%': { opacity: '0', transform: 'translateX(20px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' },
        },
        scaleIn: {
          '0%': { opacity: '0', transform: 'scale(0.95)' },
          '100%': { opacity: '1', transform: 'scale(1)' },
        },
        shimmer: {
          '0%': { backgroundPosition: '-200% 0' },
          '100%': { backgroundPosition: '200% 0' },
        },
        glow: {
          '0%': { boxShadow: '0 0 20px rgba(194,24,91,0.25)' },
          '100%': { boxShadow: '0 0 30px rgba(194,24,91,0.4)' },
        },
      },
      backdropBlur: {
        xs: '2px',
      },
      borderRadius: {
        '4xl': '2rem',
      },
    },
  },
  plugins: [],
}
